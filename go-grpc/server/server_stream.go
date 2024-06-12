package main

import (
	"time"

	pb "go-grpc/proto"
)

func (server *HelloServer) HelloServerStream(request *pb.Names, stream pb.HelloService_HelloServerStreamServer) error {
	for _, name := range request.Names {
		response := &pb.HelloResponse{
			Message: "Hello " + name,
		}

		if err := stream.Send(response); err != nil {
			return err
		}

		time.Sleep(3 * time.Second)
	}

	return nil
}
