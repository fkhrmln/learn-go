package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "go-grpc/proto"
)

const (
	HOST = "127.0.0.1"
	PORT = "8080"
)

type HelloServer struct {
	pb.HelloServiceServer
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", HOST, PORT))

	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	pb.RegisterHelloServiceServer(server, &HelloServer{})

	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
