package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"grpc_golang/unary/server/protofiles/greetpb"
	"log"
)

func main() {

	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := greetpb.NewGreetServiceClient(conn)

	getGreeting("Quvonchbek", "uz", client)
	getGreeting("Quvonchbek", "en", client)
}

func getGreeting(name, country string, c greetpb.GreetServiceClient) {

	log.Println("creating greeting")

	res, err := c.Greet(context.Background(), &greetpb.GreetRequest{
		CountryCode: country,
		UserName:    name,
	})

	if err != nil {
		log.Println("error: ", err)
		grpclog.Fatalf("failed to coll to Greet: %v", err)
	}

	log.Println(res.Result)
}
