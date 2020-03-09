package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/anaetrezve/golang-grpc-greeter/greet-pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error connecting to server %v", err)
	}
	
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	req := &pb.GreetRequest{
		Name: "Anaet",
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		fmt.Printf("something wrong getting greet %v", err)
	}
	fmt.Println(res.GetMessage())
}
