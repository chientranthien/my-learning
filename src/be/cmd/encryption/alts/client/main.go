package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/alts"

	pb "be/internal/proto"
)

func main() {
	creds := alts.NewClientCreds(alts.DefaultClientOptions())
	conn, err := grpc.Dial("localhost:5001",grpc.WithTransportCredentials(creds))
	if err != nil {
		return
	}

	client := pb.NewGreeterClient(conn)

	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "test"})
	log.Printf("resp=%v,err=%v", resp, err)
}
