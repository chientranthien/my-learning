package main

import (
	"context"
	"crypto/tls"
	"flag"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"be/data"
	"be/internal/greeter_server"
	pb "be/internal/proto"
)

var (
	addr               = flag.String("addr", ":5001", "")
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errUnauthenticated = status.Errorf(codes.Unauthenticated, "unauthenticated")
)

func main() {
	flag.Parse()
	cert, err := tls.LoadX509KeyPair(data.Path("x509/server_cert.pem"), data.Path("x509/server_key.pem"))
	if err != nil {
		log.Fatalf("failed to load cert, err=%v", err)
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(ensureToken),
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listern err=%v", err)
	}
	server := grpc.NewServer(opts...)
	pb.RegisterGreeterServer(server, &greeter_server.GreeterServer{})
	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve err=%v", err)
	}
}

func valid(authorization []string) bool {
	log.Println(authorization)
	if len(authorization) < 1 {
		return false
	}

	token := strings.TrimPrefix(authorization[0], "Bearer ")
	// Perform the token validation here. For the sake of this example, the code
	// here forgoes any of the usual OAuth2 token validation and instead checks
	// for a token matching an arbitrary string.
	return token == "some-secret-token"
}
func ensureToken(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	metadata, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}
	if !valid(metadata["authorization"]) {
		return nil, errUnauthenticated
	}

	return handler(ctx, req)
}
