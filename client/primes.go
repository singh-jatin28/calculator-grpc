package main

import (
	"context"
	"io"
	"log"

	pb "github.com/singh-jatin28/calculator-grpc/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Println("Generating prime numbers")
	req := &pb.PrimeRequest{
		Num: 44,
	}
	stream, err := c.Prime(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling Prime func : %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something happened: %v\n", err)
		}

		log.Println(res.Res)
	}
}
