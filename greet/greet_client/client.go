package main

import (
	"log"

	"github.com/orinayo/grpc-go/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	clientConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer clientConn.Close()

	client := greetpb.NewGreetServiceClient(clientConn)
}
