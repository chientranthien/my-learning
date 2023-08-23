package main

import (
	"context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "be/internal/proto"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s server) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	err := status.New(codes.InvalidArgument, "invalid argument")
	details, err2 := err.WithDetails(&pb.HelloRequest{Name: "aaaaaaaaaaaaaaaaaaa"})
	if err2 != nil {
		return nil, err2
	}

	return nil, details.Err()
}

func (s server) SayHellos(hellosServer pb.Greeter_SayHellosServer) error {
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		return
	}

	ser := grpc.NewServer()
	pb.RegisterGreeterServer(ser, server{})
	ser.Serve(lis)
}
