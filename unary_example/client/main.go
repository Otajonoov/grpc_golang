package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc_golang/unary_example/server/protofiles/greetpb"
	"log"
)

func main() {
	cc, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	getGreeting("Quvonchbek", "uz", c)
	getGreeting("Quvonchbek", "en", c)
}

func getGreeting(name, country string, c greetpb.GreetServiceClient) {

	log.Println("creating greeting")

	res, err := c.Greet(context.Background(), &greetpb.GreetRequest{
		CountryCode: country,
		UserName:    name,
	})

	if err != nil {
		log.Println("error: ", err)
		panic(err)
	}

	log.Println(res.Result)
}
