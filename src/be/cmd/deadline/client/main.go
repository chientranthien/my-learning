package main

import (
	"context"
	"flag"
	"log"
	"time"

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
		log.Fatalf("err=%v",err)
	}


	client := pb.NewGreeterClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "propagate"})
	log.Printf("resp=%v, err=%v", resp, err)

}
