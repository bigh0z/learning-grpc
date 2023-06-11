package main

import (
	"log"
	"net"

	pb "github.com/bigh0z/learning-grpc/greet/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetSserviceServer
}

func main() {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v/n", err)
	}

	log.Printf("Listening on: %v/n", addr)

	server := grpc.NewServer()
	pb.RegisterGreetSserviceServer(server, &Server{})

	if err := server.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v/n", err)
	}
}
