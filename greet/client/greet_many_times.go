package main

import (
	"context"
	"io"
	"log"

	pb "github.com/bigh0z/learning-grpc/greet/proto"
)

func doGreetManyTimes(client pb.GreetSserviceClient) {
	log.Println("doGreetManyTimes was invoked!")

	req := &pb.GreetRequest{
		FirstName: "Bilal",
	}
	stream, err := client.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling the stream: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream : %v\n", err)
		}

		log.Printf("doGreetManyTimes %s\n", msg.Result)
	}
}
