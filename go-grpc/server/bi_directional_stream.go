package main

import (
	"fmt"
	"io"

	pb "go-grpc/proto"
)

func (server *HelloServer) HelloBiDirectionalStream(stream pb.HelloService_HelloBiDirectionalStreamServer) error {
	for {
		request, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		fmt.Println("Hello", request.Name)

		response := *&pb.HelloResponse{Message: "Hello " + request.Name}

		if err := stream.Send(&response); err != nil {
			return err
		}
	}

	return nil
}
