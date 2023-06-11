package main

import (
	"fmt"
	"log"

	pb "github.com/bigh0z/learning-grpc/greet/proto"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetSservice_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes func was invoked %v\n", in)
	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", in.FirstName, i)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}
