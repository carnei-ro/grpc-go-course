package main

import (
	"google.golang.org/grpc"
	"net"
	"github.com/carnei-ro/grpc-go-course/calculator/calculatorpb"
	"context"
	"log"
)

type server struct{}

func (*server) Calc(ctx context.Context, req *calculatorpb.CalcRequest) (*calculatorpb.CalcResponse, error) {
	first := req.GetCalc().GetFirst()
	second := req.GetCalc().GetSecond()
	result := first + second

	res := &calculatorpb.CalcResponse {
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