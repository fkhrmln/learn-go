package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "go-grpc/proto"
)

func HelloServerStream(client pb.HelloServiceClient, names *pb.Names) {
	stream, err := client.HelloServerStream(context.Background(), names)

	if err != nil {
		log.Fatal(err)
	}

	for {
		response, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(response.Message)
	}
}
