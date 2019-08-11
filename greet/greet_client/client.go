package main

import (
	"io"
	"context"
	"github.com/carnei-ro/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello I'm a client")
	
	cc,err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close() // defer faz executar o codigo logo antes de a execucao do bloco terminar

	c := greetpb.NewGreetServiceClient(cc)
	
	//doUnary(c)

	doServerStreaming(c)
}

func doUnary (c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Unary RPC")

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Leandro",
			LastName: "Carneiro",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Leandro",
			LastName: "Carneiro",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes RPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}
		log.Printf("Response from GreetManyTimes: %v", msg.GetResult() )
	}
}