package main

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/alts"

	"be/internal/greeter_server"
	pb "be/internal/proto"
)

func main() {
	addr := ":5001"

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}

	creds := alts.NewServerCreds(alts.DefaultServerOptions())
	server := grpc.NewServer(grpc.Creds(creds))

	pb.RegisterGreeterServer(server, greeter_server.GreeterServer{Name: addr})
	server.Serve(lis)
}
