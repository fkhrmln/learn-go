package main

import (
	"context"

	pb "go-grpc/proto"
)

func (server *HelloServer) HelloUnary(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello " + request.Name,
	}, nil
}
