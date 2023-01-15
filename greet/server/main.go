package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/lunastryke/grpc-intro/greet/proto"
	"google.golang.org/grpc"
)

var addr string = "localhost"
var port string = "8080"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	// Create a new server
	lis, err := net.Listen("tcp", addr+":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("Listening on " + addr + ":" + port)
	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
