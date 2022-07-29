package server_methods

import (
	"fmt"

	pb "github.com/singh-jatin28/calculator-grpc/proto"
)

func (*Server) Prime(in *pb.PrimeRequest, stream pb.CalculatorService_PrimeServer) error {
	fmt.Println("Calling prime function")

	num := int(in.Num)
	for i := 2; i < num; i++ {
		if isPrime(i) {
			stream.Send(&pb.PrimeResponse{
				Res: int64(i),
			})
		}
	}

	return nil
}

func isPrime(n int) bool {

	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}

	// This is checked so that we can skip
	// middle five numbers in below loop
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i = i + 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true

}
