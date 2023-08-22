package main

import (
	"context"
	"flag"
	"log"
	"net"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "be/internal/proto"
)

var (
	addr = flag.String("addr", "localhost:5001", "")
)

type server struct {
	pb.UnimplementedGreeterServer
	c pb.GreeterClient
}

func NewServer() *server {
	conn, _ := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	c := pb.NewGreeterClient(conn)
	return &server{
		c: c,
	}
}

func (s server) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Println("request")
	if strings.HasPrefix(request.Name, "propagate") {
		time.Sleep(200 * time.Millisecond)
		 s.c.SayHello(ctx, request)
	}
	time.Sleep(500 * time.Millisecond)

	return &pb.HelloReply{Message: "hello"}, nil
}

func (s server) SayHellos(hellosServer pb.Greeter_SayHellosServer) error {
	return nil
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to lis, err=%v", err)
	}

	ser := grpc.NewServer()
	pb.RegisterGreeterServer(ser, NewServer())
	ser.Serve(lis)
}
