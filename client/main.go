package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	pb "example.com/calculadora/calculator"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func conversorFloat(numer1 string) float32 {
	numero1, err := strconv.ParseFloat(numer1, 32)
	if err != nil {
		log.Fatalf("Numero invalido: %v", err)
	}
	return float32(numero1)
}

func conversorInt(operator string) int32 {
	numero1, err := strconv.ParseInt(operator, 10, 32)
	if err != nil {
		log.Fatalf("Operador invalido: %v", err)
	}
	return int32(numero1)
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCalculatorServiceClient(conn)

	var num1 string
	var operator string
	var num2 string

	fmt.Print("Digite um numero:")
	fmt.Scanln(&num1)
	n1 := conversorFloat(num1)

	fmt.Print("Digite um Operador:")
	fmt.Scanln(&operator)
	op := conversorInt(operator)

	fmt.Print("Digite um numero:")
	fmt.Scanln(&num2)
	n2 := conversorFloat(num2)

	request := &pb.Request{
		NumberOne: float32(n1),
		NumberTwo: float32(n2),
		Operation: pb.OperatorType(op),
	}

	res, err := c.Calculate(context.Background(), request)
	if res != nil {
		log.Println(res.Result)
	}
	log.Println(err)

}
