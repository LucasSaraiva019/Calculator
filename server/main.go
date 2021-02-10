package main

import (
	"context"
	"errors"
	"log"
	"net"

	pb "example.com/calculadora/calculator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

//Math ...
type Math struct {
	pb.UnsafeCalculatorServiceServer
}

//(c *Math) = Receiver
//Calculate2 ...
func (c *Math) Calculate(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	var res float32
	switch in.Operation {
	case pb.OperatorType_SUM:
		res = in.NumberOne + in.NumberTwo
	case pb.OperatorType_SUBTRACTION:
		res = in.NumberOne - in.NumberTwo
	case pb.OperatorType_MULTIPLICATION:
		res = in.NumberOne * in.NumberTwo
	case pb.OperatorType_DIVISION:
		res = in.NumberOne / in.NumberTwo
	default:
		return &pb.Response{Result: -1}, errors.New("Operador invalido")
	}
	return &pb.Response{Result: res}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Math{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}