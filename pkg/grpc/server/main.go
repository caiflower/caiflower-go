package main

import (
	"caiflower.com/caiflower-go/pkg/grpc/server/api/pb"
	"caiflower.com/caiflower-go/pkg/grpc/server/service"
	"google.golang.org/grpc"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	// 注册helloImpl
	pb.RegisterIServiceServer(grpcServer, &service.HelloImpl{})

	err = grpcServer.Serve(listen)
	if err != nil {
		panic(err)
	}
}
