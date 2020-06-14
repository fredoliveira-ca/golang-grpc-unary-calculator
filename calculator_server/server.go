package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/fredoliveira-ca/golang-grpc/calculator/calculatorpb"
)

type server struct{}

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

func (*server) Calculate(ctx context.Context, req *calculatorpb.Request) (*calculatorpb.Response, error) {
	fmt.Printf("Received RPC call: %v \n", req)

	firstNumber := req.FisrtNumber
	secondNumber := req.SecondNumber
	operation := req.Operation
	result := service.Calculate(firstNumber, secondNumber, operation)

	return &calculatorpb.Response{
		Result: result,
	}, nil
}
