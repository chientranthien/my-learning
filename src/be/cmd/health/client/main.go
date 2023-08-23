package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	_ "google.golang.org/grpc/health"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/resolver/manual"

	pb "be/internal/proto"
)

func main() {
	r := manual.NewBuilderWithScheme("aaaaaaaa")
	conn, err := grpc.Dial(
		r.Scheme()+":///unused",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithResolvers(r),
		grpc.WithDefaultServiceConfig(
			`
		{
			"loadBalancingPolicy":"round_robin",
			"healthCheckConfig":{"serviceName":""}
		}
		`,
		),
	)
	if err != nil {
		log.Fatalf("err=%v", err)
	}
	r.UpdateState(resolver.State{
		Addresses: []resolver.Address{
			{
				Addr: ":5001",
			},
			{
				Addr: ":5002",
			},
		},
	})

	client := pb.NewGreeterClient(conn)

	resp, err := client.SayHello(
		context.Background(),
		&pb.HelloRequest{Name: "abc"},
	)
	log.Printf("resp=%v,err=%v", resp, err)
	for i := 0; i < 10; i++ {
		resp, err := client.SayHello(
			context.Background(),
			&pb.HelloRequest{Name: "abc"},
		)
		log.Printf("resp=%v,err=%v", resp, err)
		time.Sleep(1 * time.Second)
	}
}
