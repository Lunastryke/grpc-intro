package main

import (
	"log"

	pb "github.com/lunastryke/grpc-intro/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:8080"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// Init client from protobuf
	c := pb.NewGreetServiceClient(conn)
	// Do functions calls here
	doGreet(c)
	doSum(c)
}
