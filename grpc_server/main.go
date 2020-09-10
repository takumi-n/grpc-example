package main

import (
	"context"
	"log"
	"net"

	pb "github.com/takumi-n/grpc-example/calculator"
	"google.golang.org/grpc"
)

const port = ":50000"

type server struct{}

func (s *server) Add(c context.Context, in *pb.OpRequest) (*pb.Result, error) {
	return &pb.Result{Result: in.X + in.Y}, nil
}

func (s *server) Sub(c context.Context, in *pb.OpRequest) (*pb.Result, error) {
	return &pb.Result{Result: in.X - in.Y}, nil
}

func (s *server) Mul(c context.Context, in *pb.OpRequest) (*pb.Result, error) {
	return &pb.Result{Result: in.X * in.Y}, nil
}

func (s *server) Div(c context.Context, in *pb.OpRequest) (*pb.Result, error) {
	return &pb.Result{Result: in.X / in.Y}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
