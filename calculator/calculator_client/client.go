package main

import (
	"context"
	"log"

	"github.com/orinayo/grpc-go/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	clientConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer clientConn.Close()

	client := calculatorpb.NewCalculatorServiceClient(clientConn)
	doUnary(client)
}

func doUnary(client calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.SumRequest{
		FirstNumber:  10,
		SecondNumber: 15,
	}
	res, err := client.Sum(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling Sum RPC: %v", err)
	}

	log.Printf("Response from Sum: %v", res.SumResult)
}
