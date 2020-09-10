package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	pb "github.com/takumi-n/grpc-example/calculator"
	"google.golang.org/grpc"
)

const num = 5000

func requestHttp() {
	url := "http://localhost:8888?op=add&x=100&y=200"
	resp, _ := http.Get(url)
	defer resp.Body.Close()
}

func requestGrpc() {
	conn, _ := grpc.Dial("localhost:50000", grpc.WithInsecure(), grpc.WithBlock())
	defer conn.Close()

	c := pb.NewCalculatorClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.OpRequest{X: 100, Y: 200}

	c.Add(ctx, req)
}

func requestGrpcWithSingleConnection() int64 {
	conn, _ := grpc.Dial("localhost:50000", grpc.WithInsecure(), grpc.WithBlock())
	defer conn.Close()

	grpcStart := time.Now()
	for i := 0; i < num; i++ {
		c := pb.NewCalculatorClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		req := &pb.OpRequest{X: 100, Y: 200}

		c.Add(ctx, req)
	}
	grpcEnd := time.Now()

	return grpcEnd.Sub(grpcStart).Milliseconds()
}

func main() {
	httpStart := time.Now()
	for i := 0; i < num; i++ {
		requestHttp()
	}
	httpEnd := time.Now()

	grpcStart := time.Now()
	for i := 0; i < num; i++ {
		requestGrpc()
	}
	grpcEnd := time.Now()

	singleConnectionTime := requestGrpcWithSingleConnection()

	fmt.Printf("http = %d [ms]\n", httpEnd.Sub(httpStart).Milliseconds())
	fmt.Printf("grpc = %d [ms]\n", grpcEnd.Sub(grpcStart).Milliseconds())
	fmt.Printf("grpc with single connection = %d [ms]\n", singleConnectionTime)
}
