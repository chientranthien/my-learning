package greeter_server

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "be/internal/proto"
)

type GreeterServer struct {
	pb.UnimplementedGreeterServer
	Name string
}

func (g GreeterServer) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("server received %v", request.Name)
	return &pb.HelloReply{Message: fmt.Sprintf("hello back from %v", g.Name)}, nil
}

func (GreeterServer) SayHellos(stream pb.Greeter_SayHellosServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("sayHellos err=%v", err)
			return err
		}

		err = stream.Send(&pb.HelloReply{Message: fmt.Sprintf("hello %v", req.Name)})
		if err != nil {
			log.Printf("reply err=%v", err)
			return err
		}
	}

	return nil
}
