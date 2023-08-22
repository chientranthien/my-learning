package main

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"be/data"
	"be/internal/greeter_server"
	pb "be/internal/proto"
)

func main() {
	addr := ":5001"

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}

	creds, err := credentials.NewServerTLSFromFile(data.Path("x509/server_cert.pem"), data.Path("x509/server_key.pem"))
	if err != nil {
		return
	}
	server := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterGreeterServer(server, greeter_server.GreeterServer{Name: addr})
	server.Serve(lis)
}
