package main

import (
	"context"
	"log"

	pb "github.com/bigh0z/learning-grpc/greet/proto"
)

func doGreet(client pb.GreetSserviceClient) {
	log.Println("doGreet was invoked!")
	res, err := client.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Bilal",
	})

	if err != nil {
		log.Fatalf("Error because: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}
