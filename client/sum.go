package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/singh-jatin28/calculator-grpc/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("Doing sum")
	r, err := c.Sum(context.Background(), &pb.SumRequest{Firstnum: 12, Secondnum: 1245})

	if err != nil {
		log.Fatalf("Could not calcualte sum: %v\n", err)
	}

	fmt.Println("Sum: ", r.Res)
}
