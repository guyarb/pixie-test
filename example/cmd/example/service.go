package main

import (
	"context"
	"fmt"
	"net"

	"github.com/guyarb/pixie-test/internal/proto"
	"google.golang.org/grpc"
)

type Service struct {
	Server *grpc.Server

	proto.UnsafeEchoServer
}

func (s Service) Sample(_ context.Context, r *proto.SampleRequest) (*proto.SampleResponse, error) {
	return &proto.SampleResponse{
		Msg: "echo " + r.GetMsg(),
	}, nil
}

func main() {
	grpcServer := Service{}

	srv := grpc.NewServer()

	proto.RegisterEchoServer(srv, grpcServer)

	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	if err := srv.Serve(listener); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
