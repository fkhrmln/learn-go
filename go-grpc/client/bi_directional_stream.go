package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "go-grpc/proto"
)

func HelloBiDirectionalStream(client pb.HelloServiceClient, names *pb.Names) {
	stream, err := client.HelloBiDirectionalStream(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	channel := make(chan struct{})

	go func() {
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
	}()

	for _, name := range names.Names {
		request := *&pb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(&request); err != nil {
			log.Fatal(err)
		}

		time.Sleep(3 * time.Second)
	}

	stream.CloseSend()

	<-channel
}
