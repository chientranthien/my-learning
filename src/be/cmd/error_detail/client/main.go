package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	pb "be/internal/proto"
)

func main() {
	conn, err := grpc.Dial(":5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}

	client := pb.NewGreeterClient(conn)
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "aaa"})
	log.Printf("resp=%v,err=%v", resp, err)
	convert := status.Convert(err)
	d := convert.Details()
	log.Println(d)
}
