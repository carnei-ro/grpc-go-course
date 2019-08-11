package main

import (
	"io"
	"fmt"
	"context"
	"github.com/carnei-ro/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
)

func main(){
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalcServiceClient(cc)

	// unaryCall(c)
	serverStreamingCall(c)
}

func unaryCall(c calculatorpb.CalcServiceClient) {
	req := &calculatorpb.SumRequest {
		First: 10,
		Second: 3,
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Calc RPC: %v", err)
	}
	log.Printf("Response from Calc: %v", res.Result)
}

func serverStreamingCall(c calculatorpb.CalcServiceClient) {
	req := &calculatorpb.PrimeNumberDecompositionRequest {
		Number: int64(69873100),
	}
	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling PrimeNumberDecomposition RPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened: %v\n", err)
		}
		fmt.Printf("%v\n", res.GetPrimeFactor())
	}
	
}