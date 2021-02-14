package main

import (
	"context"
	"fmt"
	"io"
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
	doServerStreaming(client)
	doClientStreaming(client)
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

func doServerStreaming(client calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Number: 12,
	}
	stream, err := client.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling PrimeDecomposition RPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}

		fmt.Println(res.GetPrimeFactor())
	}
}

func doClientStreaming(client calculatorpb.CalculatorServiceClient) {
	stream, err := client.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Error while opening stream: %v", err)
	}
	numbers := []int32{3, 5, 9, 23, 54}

	for _, num := range numbers {
		stream.Send(&calculatorpb.ComputeAverageRequest{
			Number: num,
		})
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response: %v", err)
	}

	fmt.Printf("The average is: %v", res.GetAverage())
}
