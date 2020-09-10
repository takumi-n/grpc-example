package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	pb "github.com/takumi-n/grpc-example/calculator"
	"google.golang.org/grpc"
)

const address = "localhost:50000"

func main() {
	if len(os.Args) < 4 {
		log.Fatalln("Need 3 args: <operator> <operandX> <operandY>")
	}

	conn, _ := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	defer conn.Close()

	c := pb.NewCalculatorClient(conn)
	operator := os.Args[1]
	operandX, _ := strconv.ParseFloat(os.Args[2], 64)
	operandY, _ := strconv.ParseFloat(os.Args[3], 64)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var (
		req    = &pb.OpRequest{X: operandX, Y: operandY}
		result *pb.Result
		err    error
	)

	switch operator {
	case "add":
		result, err = c.Add(ctx, req)
	case "sub":
		result, err = c.Sub(ctx, req)
	case "mul":
		result, err = c.Mul(ctx, req)
	case "div":
		result, err = c.Div(ctx, req)
	}

	if err != nil {
		log.Fatalf("Faild to run rpc: %v\n", err)
	}

	log.Println(result)
}
