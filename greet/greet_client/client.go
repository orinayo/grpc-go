package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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
	doUnary(client)
	doServerStreaming(client)
	doClientStreaming(client)
}

func doUnary(client greetpb.GreetServiceClient) {
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Ayo",
			LastName:  "Oyelade",
		},
	}
	res, err := client.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}

	log.Printf("Response from Greet: %v", res.Result)
}

func doServerStreaming(client greetpb.GreetServiceClient) {
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Ayo",
			LastName:  "Oyelade",
		},
	}

	resStream, err := client.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes RPC: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}

		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
	}
}

func doClientStreaming(client greetpb.GreetServiceClient) {
	requests := []*greetpb.LongGreetRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Ayo",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Efe",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Victory",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Andy",
			},
		},
	}
	stream, err := client.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v", err)
	}

	for _, req := range requests {
		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v", err)
	}

	fmt.Printf("LongGreet Response: %v", res)
}
