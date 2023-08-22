package main

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/channelz/service"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/resolver/manual"

	pb "be/internal/proto"
)

var (
	addrs = []string{":5001", ":5002", ":5003"}
)

func main() {
	lis, err := net.Listen("tcp", ":4999")
	if err != nil {
		log.Fatalf("err=%v", err)
	}

	ser := grpc.NewServer()

	service.RegisterChannelzServiceToServer(ser)
	go ser.Serve(lis)
	if err != nil {
		log.Fatalf("err=%v")
	}

	r := manual.NewBuilderWithScheme("schema")
	conn, err := grpc.Dial(
		r.Scheme()+"://test.example",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithResolvers(r),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	)
	r.UpdateState(resolver.State{
		Addresses: []resolver.Address{
			{
				Addr: addrs[0],
			},
			{
				Addr: addrs[1],
			},
			{
				Addr: addrs[2],
			},
		},
	})
	if err != nil {
		log.Fatalf("err=%v")
	}

	client := pb.NewGreeterClient(conn)

	for i := 0; i < 100; i++ {
		ctx, _ := context.WithTimeout(context.Background(), 150*time.Millisecond)
		resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "test"})
		log.Printf("resp=%v, err=%v", resp, err)
	}
	select {}
}
