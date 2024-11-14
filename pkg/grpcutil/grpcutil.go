package grpcutil

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewServer(opts ...grpc.ServerOption) *grpc.Server {
	server := grpc.NewServer(opts...)
	reflection.Register(server)
	return server
}
