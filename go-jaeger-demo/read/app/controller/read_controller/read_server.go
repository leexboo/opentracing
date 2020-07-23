package read_controller

import (
	"fmt"
	"golang.org/x/net/context"
	"demo/go-jaeger-demo/read/app/proto/read"
	"demo/go-jaeger-demo/read/app/util"
	"demo/go-jaeger-demo/read/app/proto/listen"
	"demo/go-jaeger-demo/read/app/util/grpc_client"
)

type ReadController struct{}

func (s *ReadController) ReadData(ctx context.Context, in *read.Request) (*read.Response, error) {

	// 调用 gRPC 服务
	grpcListenClient := listen.NewListenClient(grpc_client.CreateServiceListenConn(ctx))
	resListen, _ := grpcListenClient.ListenData(context.Background(), &listen.Request{Name: "listen"})

	// 调用 HTTP 服务
	resHttpGet := ""
	_, err := util.HttpGet("http://10.192.8.173:31117/sing", ctx)
	if err == nil {
		resHttpGet = "[HttpGetOk]"
	}

	msg := "[" + fmt.Sprintf("%s", in.Name) + "-" +
		   resListen.Message + "-" +
		   resHttpGet +
		   "]"
	return &read.Response{Message : msg}, nil
}
