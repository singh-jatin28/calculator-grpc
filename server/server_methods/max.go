package server_methods

import (
	"fmt"
	"io"
	"log"

	pb "github.com/singh-jatin28/calculator-grpc/proto"
)

func (*Server) Max(stream pb.CalculatorService_MaxServer) error {
	fmt.Println("Calling max function")
	var maxnum int64 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client : %v\n", err)
		}

		if req.Num > maxnum {
			maxnum = req.Num
			err = stream.Send(&pb.MaxResponse{
				Res: maxnum,
			})

			if err != nil {
				log.Fatalf("Error while sending data : %v\n", err)
			}
		}
	}
}
