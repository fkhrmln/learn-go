package main

import (
	"context"
	"fmt"
	"log"

	pb "go-grpc/proto"
)

func HelloUnary(client pb.HelloServiceClient, name string) {
	ctx := context.Background()

	request := &pb.HelloRequest{Name: name}

	response, err := client.HelloUnary(ctx, request)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response.Message)
}
