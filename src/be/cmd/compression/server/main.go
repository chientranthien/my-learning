package main

import (
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"

	_ "google.golang.org/grpc/encoding/gzip"

	"be/internal/greeter_server"
	pb "be/internal/proto"
)

var (
	addr = flag.String("addr", ":5001", "")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to lis, err=%v", err)
	}

	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &greeter_server.GreeterServer{Name: *addr})
	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve, err=%v", err)
	}
}
