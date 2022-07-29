package server_methods

import (
	"context"
	"fmt"

	pb "github.com/singh-jatin28/calculator-grpc/proto"
)

func (*Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	fmt.Println("Calling Sum function")
	return &pb.SumResponse{Res: in.Firstnum + in.Secondnum}, nil
}
