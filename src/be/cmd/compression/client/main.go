package main

import (
	"context"
	"flag"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/encoding/gzip"

	pb "be/internal/proto"
)

var (
	addr = flag.String("addr", "localhost:5001", "")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to start conn, err=%v", err)
	}

	client := pb.NewGreeterClient(conn)
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "test"}, grpc.UseCompressor(gzip.Name))
	log.Printf("resp=%v, err=%v", resp, err)
}
