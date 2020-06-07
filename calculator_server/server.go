package main

import (
	"fmt"
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	"github.com/fredoliveira-ca/golang-grpc/calculator/calculatorpb"
)

type server struct {}

func main() {
	fmt.Println("Calculator Server is runnig and waiting to serve...")

	listen, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	newServer := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(newServer, &server{})

	if err := newServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Received RPC call: %v \n", req)
	
	firstNumber := req.FisrtNumber
	secondNumber := req.SecondNumber
	sum := firstNumber + secondNumber
	
	return &calculatorpb.SumResponse {
		SumResult: sum,
	}, nil
}