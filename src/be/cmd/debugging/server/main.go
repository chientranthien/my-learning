package main

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/channelz/service"

	pb "be/internal/proto"
)

var (
	addrs = []string{":5001", ":5002", ":5003"}
)

type server struct {
	pb.UnimplementedGreeterServer
	name string
}

func (s server) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello back from" + s.name}, nil
}

func (s server) SayHellos(hellosServer pb.Greeter_SayHellosServer) error {
	return nil
}

type slowerServer struct {
	pb.UnimplementedGreeterServer
	name string
}

func (s slowerServer) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	time.Sleep(200 * time.Millisecond)
	return &pb.HelloReply{Message: "hello back from" + s.name}, nil
}

func (s slowerServer) SayHellos(hellosServer pb.Greeter_SayHellosServer) error {
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("err=%v", err)
	}
	ser := grpc.NewServer()
	service.RegisterChannelzServiceToServer(ser)
	go ser.Serve(lis)
	if err != nil {
		log.Fatalf("err=%v", err)
	}
	defer lis.Close()

	for i, addr := range addrs {
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatalf("err=%v", err)
		}

		ser := grpc.NewServer()
		if i != 2 {
			pb.RegisterGreeterServer(ser, &server{name: addr})
		} else {
			pb.RegisterGreeterServer(ser, &slowerServer{name: addr})
		}
		go ser.Serve(lis)
	}

	select {}
}
