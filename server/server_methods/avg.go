package server_methods

import (
	"fmt"
	"io"
	"log"

	pb "github.com/singh-jatin28/calculator-grpc/proto"
)

func (*Server) Avg(stream pb.CalculatorService_AvgServer) error {
	fmt.Println("Calling avg function")

	var sum int64 = 0
	count := 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Res: float64(sum) / float64(count),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client : %v\n", err)
		}

		sum = sum + req.Num
		count++
	}
}
