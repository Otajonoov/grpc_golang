package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"grpc_golang/unary/server/protofiles/greetpb"
	"log"
	"net"
)

// this struct could be used to dependency injection
type servver struct{}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	fmt.Println("starting server")

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, servver{})

	if err := s.Serve(listener); err != nil {
		panic(err)
	}
}

// handler
func (s servver) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {

	log.Println("User name: ", req.UserName)
	log.Println("Country code: ", req.CountryCode)

	var greeting string

	switch req.CountryCode {
	case "uz":
		greeting = "Assalomu alaykum " + req.UserName
	case "en	":
		greeting = "Hello " + req.UserName
	default:
		greeting = "Assalomu alaykum/Hello " + req.UserName
	}

	return &greetpb.GreetResponse{
		Result: greeting,
	}, nil
}
