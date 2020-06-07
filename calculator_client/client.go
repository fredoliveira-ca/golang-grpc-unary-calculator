package main

import (
	"github.com/fredoliveira-ca/golang-grpc/calculator/calculatorpb"
	"fmt"
	"context"
	"log"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Calculator Client has been started!")
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer connection.Close()

	calc := calculatorpb.NewCalculatorServiceClient(connection)

	doUnaryCall(calc)
}

func doUnaryCall(calc calculatorpb.CalculatorServiceClient) {
	fmt.Println("Startin an unary RPC call...!")
	request := &calculatorpb.SumRequest{
		FisrtNumber: 2,
		SecondNumber: 2,
	}

	response, err := calc.Sum(context.Background(), request)
	if err != nil {
		log.Fatalf("A failure occurred while calling server: %v", err)
	}
	log.Printf("Response: %v", response.SumResult)
}