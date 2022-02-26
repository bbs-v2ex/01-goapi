package main

import (
	"context"
	"net"

	"google.golang.org/grpc"

	"n01/006_grpc_go_2/proto"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Message: "hello " + request.Name,
	}, nil
}

func main() {
	//实例化 grpc
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &Server{})
	//启动服务
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
	err = g.Serve(listener)
	if err != nil {
		panic(err)
	}
}
