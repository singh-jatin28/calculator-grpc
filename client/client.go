package main

import (
	"log"

	methods "github.com/singh-jatin28/calculator-grpc/client/client_methods"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/singh-jatin28/calculator-grpc/proto"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Did not connect: %v\n", err)
	}

	defer conn.Close()
	c := pb.NewCalculatorServiceClient(conn)

	methods.CalculateSum(c)
	methods.CalculatePrime(c)
	methods.CalculateAvg(c)
	methods.CalculateMax(c)

}
