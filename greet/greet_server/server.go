package main

import (
	"context"
	"log"
	"net"

	"github.com/orinayo/grpc-go/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	newServer := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(newServer, &server{})

	if err := newServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
