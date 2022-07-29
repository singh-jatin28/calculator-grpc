package server_methods

import (
	pb "github.com/singh-jatin28/calculator-grpc/proto"
)

type Server struct {
	pb.CalculatorServiceServer
}
