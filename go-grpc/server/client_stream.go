package main

import (
	"fmt"
	"io"

	pb "go-grpc/proto"
)

func (server *HelloServer) HelloClientStream(stream pb.HelloService_HelloClientStreamServer) error {
	messages := []string{}

	for {
		request, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.Messages{Messages: messages})
		}

		if err != nil {
			return err
		}

		fmt.Println("Hello", request.Name)
	}
}
