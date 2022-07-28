//go:build !test
// +build !test

package main

import (
	"log"
	"net"

	pb "github.com/singh-jatin28/calculator-grpc/proto"

	"google.golang.org/grpc"
)

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	log.Printf("Listening at %s\n", ":50051")

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error while listening: %v\n", err)
	}
}
