package main

import (
	"context"
	"flag"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "be/internal/proto"
)

var (
	addr = flag.String("addr", "localhost:5001", "")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}

	client := pb.NewGreeterClient(conn)

	ctx, cancelFunc := context.WithCancel(context.Background())
	stream, err := client.SayHellos(ctx)
	if err != nil {
		log.Fatalf("failed to sayHellos, err=%v", err)
		return
	}

	err = stream.Send(&pb.HelloRequest{Name: "test"})
	log.Printf(" err=%v", err)

	resp, err := stream.Recv()
	log.Printf("resp=%v, err=%v", resp, err)
	cancelFunc()

	err = stream.Send(&pb.HelloRequest{Name: "test"})
	log.Printf(" err=%v", err)

	resp, err = stream.Recv()
	log.Printf("resp=%v, err=%v", resp, err)


	log.Printf("---------------------")

	ctx, cancelFunc = context.WithCancel(context.Background())
	stream, err = client.SayHellos(ctx)
	if err != nil {
		log.Fatalf("failed to sayHellos, err=%v", err)
		return
	}

	err = stream.Send(&pb.HelloRequest{Name: "test"})
	log.Printf(" err=%v", err)

	resp, err = stream.Recv()
	log.Printf("resp=%v, err=%v", resp, err)
	cancelFunc()

	resp, err = stream.Recv()
	log.Printf("resp=%v, err=%v", resp, err)
}
