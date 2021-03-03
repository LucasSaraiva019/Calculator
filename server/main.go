package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/lucas.saraiva019/calculadora/proto/calculator"
	"google.golang.org/grpc"
)

//Math ...
type Server struct {
	pb.UnsafeCalculatorServiceServer
}

func newServer() *Server {
	return &Server{}
}

//(s *server) = Receiver
//Calculate2 ...
func (s *Server) Calculate(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	var res float32
	var err error

	switch in.Operation {
	case pb.OperatorType_SUM:
		res = in.NumberOne + in.NumberTwo
	case pb.OperatorType_SUBTRACTION:
		res = in.NumberOne - in.NumberTwo
	case pb.OperatorType_MULTIPLICATION:
		res = in.NumberOne * in.NumberTwo
	case pb.OperatorType_DIVISION:
		if in.NumberTwo == 0 {
			err = errors.New("Não é possivel dividir por Zero")
		} else {
			res = in.NumberOne / in.NumberTwo
		}
	default:
		return &pb.Response{Result: -1}, errors.New("Operador invalido")
	}
	if err != nil {
		return nil, err
	}
	return &pb.Response{Result: res}, nil
}

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	pb.RegisterCalculatorServiceServer(s, &Server{})

	//Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		log.Fatal(s.Serve(lis))
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()

	//Register Calculator
	err = pb.RegisterCalculatorServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}
	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}
	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())

}
