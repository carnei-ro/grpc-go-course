package main

import (
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

	unaryCall(c)
}

func unaryCall(c calculatorpb.CalcServiceClient) {
	req := &calculatorpb.CalcRequest {
		Calc: &calculatorpb.Calc{
			First: 10,
			Second: 3,
		},
	}

	res, err := c.Calc(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Calc RPC: %v", err)
	}
	log.Printf("Response from Calc: %v", res.Result)
}