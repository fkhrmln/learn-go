package main

import (
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "go-grpc/proto"
)

const (
	HOST = "127.0.0.1"
	PORT = "8080"
)

func main() {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", HOST, PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)

	// HelloUnary(client, "Fakhri Maulana Ihsan")

	names := *&pb.Names{Names: []string{"Fakhri Maulana Ihsan", "Rifky Ferdiansyah", "Audry Elgalia"}}

	// HelloServerStream(client, &names)

	// HelloClientStream(client, &names)

	HelloBiDirectionalStream(client, &names)
}
