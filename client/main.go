package main

import (
	"log"

	pb "github.com/jopaleti/grpc-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList {
		Names: []string{"Opaleti", "Oluwatobi", "Johnson"},
	}

	// Unary RPC
	// callSayHello(client)

	// Server Streaming RPC
	// callSayHelloServerStream(client, names)

	// Client Streaming gRPC
	callSayHelloClientStream(client, names)

	// Bidirectional Streaming gRPC
	// callHelloBidirectionalStream(client, names)
}