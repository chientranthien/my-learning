package main

import (
	"context"
	"flag"
	"log"

	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"

	"be/data"
	pb "be/internal/proto"
)

var (
	addr = flag.String("addr", "localhost:5001", "")
)

func main() {
	flag.Parse()

	tokenSource := oauth.TokenSource{TokenSource: oauth2.StaticTokenSource(fetchToken())}
	cert, err := credentials.NewClientTLSFromFile(data.Path("x509/ca_cert.pem"), "*.test.example.com")
	if err != nil {
		log.Fatalf("failed to new client tls, err=%v", err)
	}

	opts := []grpc.DialOption{
		grpc.WithPerRPCCredentials(tokenSource),
		grpc.WithTransportCredentials(cert),
	}
	conn, err := grpc.Dial(*addr, opts...)
	if err != nil {
		log.Fatalf("failed to dial, err=%v", err)
	}

	client := pb.NewGreeterClient(conn)

	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "test"})
	log.Printf("res=%v, err=%v", resp, err)
}

func fetchToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken: "some-secret-token",
	}
}
