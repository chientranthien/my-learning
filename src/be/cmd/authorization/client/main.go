package main

import (
	"context"
	"flag"
	"log"

	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"

	"be/cmd/authorization/token"
	"be/data"
	pb "be/internal/proto"
)

var (
	addr = flag.String("addr", "localhost:5001", "")
)

func main() {
	flag.Parse()
	creds, err := credentials.NewClientTLSFromFile(data.Path("x509/ca_cert.pem"), "*.test.example.com")
	if err != nil {
		log.Fatalf("failed to new client tls, err=%v", err)
	}

	tk := token.Token{
		Secret:   "my-secret",
		Username: "root",
	}
	encodeTk, err := tk.Encode()
	if err != nil {
		log.Printf("failed to encode tk, err=%v",err)
	}
	staticToken := &oauth2.Token{
		AccessToken: encodeTk,
	}
	tokenSource := oauth.TokenSource{TokenSource: oauth2.StaticTokenSource(staticToken)}
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(tokenSource),
	}

	conn, err := grpc.Dial(*addr, opts...)
	if err != nil {
		log.Fatalf("failed to dial, err=%v", err)
	}
	client := pb.NewGreeterClient(conn)
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "test"})
	log.Printf("resp=%v, err=%v", resp, err)
}
