package main

import (
	"context"
	"log"
	"net"

	"github.com/orinayo/grpc-go/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	firstNumber := req.GetFirstNumber()
	secondNumber := req.GetSecondNumber()
	result := firstNumber + secondNumber
	res := &calculatorpb.SumResponse{
		SumResult: result,
	}
	return res, nil
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	newServer := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(newServer, &server{})

	if err := newServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
