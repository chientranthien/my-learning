package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"be/data"
	pb "be/internal/proto"
)

func main() {
	pair, err := credentials.NewClientTLSFromFile(data.Path("x509/ca_cert.pem"), "*.test.example.com")
	if err != nil {
		return
	}
	conn, err := grpc.Dial(":5001", grpc.WithTransportCredentials(pair))
	if err != nil {
		return
	}

	client := pb.NewGreeterClient(conn)
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{
		Name: "aaa",
	})
	if err != nil {
		log.Fatalf("err=%v",err)
		return
	}

	log.Printf("resp=%v,err=%v", resp, err)
}
