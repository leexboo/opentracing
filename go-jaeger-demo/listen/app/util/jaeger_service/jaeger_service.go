package jaeger_service

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"io"
	"strings"
)

// ServerOption grpc server option
func ServerOption(tracer opentracing.Tracer) grpc.ServerOption {
	return grpc.UnaryInterceptor(serverInterceptor(tracer))
}

var Tracer opentracing.Tracer

func NewJaegerTracer(serviceName string) (opentracing.Tracer, io.Closer, error) {

	cfg := &jaegerConfig.Configuration {
		Sampler: &jaegerConfig.SamplerConfig{
			Type  : "const", //固定采样
			Param : 1,       //1=全采样、0=不采样
		},

		Reporter: &jaegerConfig.ReporterConfig{
			LogSpans           : true,
		},

		ServiceName: serviceName,
	}

	Tracer, closer, err := cfg.NewTracer(jaegerConfig.Logger(jaeger.StdLogger))

	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.SetGlobalTracer(Tracer)
	return Tracer, closer, err
}

type MDReaderWriter struct {
	metadata.MD
}

// ForeachKey implements ForeachKey of opentracing.TextMapReader
func (c MDReaderWriter) ForeachKey(handler func(key, val string) error) error {
	for k, vs := range c.MD {
		for _, v := range vs {
			if err := handler(k, v); err != nil {
				return err
			}
		}
	}
	return nil
}

// Set implements Set() of opentracing.TextMapWriter
func (c MDReaderWriter) Set(key, val string) {
	key = strings.ToLower(key)
	c.MD[key] = append(c.MD[key], val)
}

// ServerInterceptor grpc server
func serverInterceptor(tracer opentracing.Tracer) grpc.UnaryServerInterceptor {
	return func(ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {

		var parentContext context.Context

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		}

		spanContext, err := tracer.Extract(opentracing.TextMap, MDReaderWriter{md})
		if err != nil && err != opentracing.ErrSpanContextNotFound {
			grpclog.Errorf("extract from metadata err: %v", err)
		} else {
			span := tracer.StartSpan(
				info.FullMethod,
				ext.RPCServerOption(spanContext),
				opentracing.Tag{Key: string(ext.Component), Value: "gRPC"},
				ext.SpanKindRPCServer,
			)
			defer span.Finish()

			parentContext = opentracing.ContextWithSpan(ctx, span)
		}

		return handler(parentContext, req)
	}
}
