package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "go-grpc/proto"
)

func HelloClientStream(client pb.HelloServiceClient, names *pb.Names) {
	stream, err := client.HelloClientStream(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	for _, name := range names.Names {
		request := &pb.HelloRequest{Name: name}

		if err := stream.Send(request); err != nil {
			log.Fatal(err)
		}

		time.Sleep(3 * time.Second)
	}

	messages, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
