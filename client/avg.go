package main

import (
	"context"
	"log"

	pb "github.com/singh-jatin28/calculator-grpc/proto"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("Doing average")
	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("Error while opening stream: %v\n", err)
	}

	numbers := []int64{33, 65, 97, 52, 235}

	for _, number := range numbers {
		stream.Send(&pb.AvgRequest{
			Num: number,
		})
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from server: %v\n", err)
	}

	log.Printf("Average: %f\n", res.Res)
}
