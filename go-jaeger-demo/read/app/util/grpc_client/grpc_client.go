package grpc_client

import (
	"context"
	"fmt"
	"demo/go-jaeger-demo/read/app/util/jaeger_service"
	"google.golang.org/grpc"
)

func CreateServiceListenConn(ctx context.Context) *grpc.ClientConn {
	return createGrpcClient("10.192.8.173:31113", ctx)
}

func createGrpcClient(serviceAddress string, ctx context.Context) *grpc.ClientConn {
	conn, err := grpc.Dial(serviceAddress, grpc.WithInsecure(), grpc.WithUnaryInterceptor(jaeger_service.ClientInterceptor(jaeger_service.Tracer, ctx)))
	if err != nil {
		fmt.Println(serviceAddress, "grpc conn err:", err)
	}
	return conn
}
