package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	//"write/app/proto/write"
	//"write/app/controller/write_controller"
	//"write/app/util/jaeger_service"
	"demo/go-jaeger-demo/write/app/proto/write"
	"demo/go-jaeger-demo/write/app/controller/write_controller"
	"demo/go-jaeger-demo/write/app/util/jaeger_service"
	"log"
	"net"
	"os"
)

const (
	ServiceName     = "gRPC-Service-Write"
	ServiceHostPort = "0.0.0.0:9904"

)

func main() {

	var serviceOpts []grpc.ServerOption

	tracer, _, err := jaeger_service.NewJaegerTracer(ServiceName)
	if err != nil {
		fmt.Printf("new tracer err: %+v\n", err)
		os.Exit(-1)
	}
	if tracer != nil {
		serviceOpts = append(serviceOpts, jaeger_service.ServerOption(tracer))
	}

	l, err := net.Listen("tcp", ServiceHostPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer(serviceOpts...)

	// 服务注册
	write.RegisterWriteServer(s, &write_controller.WriteController{})

	log.Println("Listen on " + ServiceHostPort)
	reflection.Register(s)
	if err := s.Serve(l); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
