package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "be/internal/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to open connect err=%v", err)
	}
	client := pb.NewGreeterClient(conn)
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "test"})
	log.Printf("resp=%v, err=%v", resp, err)

}
