package client_methods

import (
	"context"
	"io"
	"log"

	pb "github.com/singh-jatin28/calculator-grpc/proto"
)

func CalculateMax(c pb.CalculatorServiceClient) {
	log.Println("Finding max on numbers")

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("Error while opening stream: %v\n", err)
	}

	wait := make(chan struct{})

	go func() {
		array := []int64{1, 3, 2, 7, 10, 12, 8, 9, 23, 13, 25}
		for _, number := range array {
			stream.Send(&pb.MaxRequest{
				Num: number,
			})

		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while reading server stream: %v\n", err)
				break
			}

			log.Printf("Max num till now : %v\n", res.Res)
		}
		close(wait)
	}()
	<-wait
}
