package greeter_server

import (
	"context"
	"fmt"

	pb "be/internal/proto"
)

type GreeterServer struct {
	pb.UnimplementedGreeterServer
	Name string
}

func (g GreeterServer) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: fmt.Sprintf("hello back from %v", g.Name)}, nil
}
