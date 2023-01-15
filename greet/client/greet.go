package main

import (
	"context"
	"log"

	pb "github.com/lunastryke/grpc-intro/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("Calling Greet function")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{FirstName: "XH"})

	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	log.Printf("Greet Response: %s", res.Result)
}

func doSum(c pb.GreetServiceClient) {
	log.Println("Calling Sum function")
	res, err := c.Sum(context.Background(), &pb.SumRequest{FirstNumber: 10, SecondNumber: 3})
	if err != nil {
		log.Fatalf("Could not sum: %v", err)
	}
	log.Printf("Sum Response: %d", res.Result)
}
