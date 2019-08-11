package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"github.com/carnei-ro/grpc-go-course/calculator/calculatorpb"
	"context"
	"log"
)

type server struct{}

func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalcService_PrimeNumberDecompositionServer) error {
	fmt.Printf("Received PrimeNumberDecomposition RPC: %v\n", req)
	number := req.GetNumber()
	divisor := int64(2)

	for number > 1 {
		if number % divisor == 0 {
			stream.Send(&calculatorpb.PrimeNumberDecompositionResponse {
				PrimeFactor: divisor,
			})
			number = number / divisor
		} else {
			divisor++
			fmt.Printf("Divisor has increased to: %v\n", divisor) //log make it slower - in the client side can be 
		}
	}
	return nil
}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	first := req.GetFirst()
	second := req.GetSecond()
	result := first + second

	res := &calculatorpb.SumResponse {
		Result: result,
	}

	return res, nil
}

func main(){
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalcServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}