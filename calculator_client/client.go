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
	fmt.Println("Startin to call an unary RPC call...!")
	request := &calculatorpb.SumRequest{
		FisrtNumber: 1423,
		SecondNumber: 1548,
	}

	response, err := calc.Sum(context.Background(), request)
	if err != nil {
		log.Fatalf("Failed while calling RPC call: %v", err)
	}
	log.Printf("Response: %v", response.SumResult)
}