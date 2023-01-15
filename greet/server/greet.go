package main

import (
	"context"
	"log"

	pb "github.com/lunastryke/grpc-intro/greet/proto"
)

// Implements Greet RPC endpoints

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function called with %v", in)
	return &pb.GreetResponse{Result: "Hello " + in.FirstName}, nil
}

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function called with %v", in)
	return &pb.SumResponse{Result: in.FirstNumber + in.SecondNumber}, nil
}
