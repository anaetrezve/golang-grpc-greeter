package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/anaetrezve/golang-grpc-greeter/greet-pb"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	return &pb.GreetResponse{
		Message: "Hello " + in.GetName(),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("can not listen to port %v", err)
	}

	srv := grpc.NewServer()

	pb.RegisterGreeterServer(srv, &server{})

	fmt.Println("Server is running on port: 50051")

	if err := srv.Serve(lis); err != nil {
		log.Fatal("can't start server")
	}
}
