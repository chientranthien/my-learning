package main

import (
	"context"
	"flag"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/authz"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"be/cmd/authorization/token"
	"be/data"
	"be/internal/greeter_server"
	pb "be/internal/proto"
)

var (
	addr            = flag.String("addr", ":5001", "")
	errUnauthorized = status.Errorf(codes.Unauthenticated, "unauthorized")

	errInvalidRequest = status.Errorf(codes.InvalidArgument, "invalid request")
)

const (
	sayHelloRole = "SAY_HELLO:RW"
	authzPolicy  = `
{
  "name": "my-policy",
  "allow_rules": [
    {
      "name": "say-hello-rule",
      "request": {
        "paths": ["/be.Greeter/SayHello"],
        "headers": [
          {
            "key": "SAY_HELLO:RW",
            "values": [ "true" ]
          }
        ]
      }
    }
  ]
}
`
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to lis, err=%v", err)
	}

	creds, err := credentials.NewServerTLSFromFile(data.Path("x509/server_cert.pem"), data.Path("x509/server_key.pem"))
	if err != nil {
		log.Fatalf("failed to create creds, err=%v", err)
	}

	authzIntecepter, err := authz.NewStatic(authzPolicy)
	if err != nil {
		log.Fatalf("failed to new authz Policy, err=%v", err)
	}

	canaryInterceptor := grpc.ChainUnaryInterceptor(
		authorizeUser,
		authzIntecepter.UnaryInterceptor,
	)
	server := grpc.NewServer(grpc.Creds(creds), canaryInterceptor)
	pb.RegisterGreeterServer(server, &greeter_server.GreeterServer{
		Name: *addr,
	})

	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve, err=%v", err)
	}

}

func authorizeUser(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errInvalidRequest
	}

	username, ok := authorizeUserFromToken(md["authorization"])
	if !ok {
		return nil, errUnauthorized
	}

	newCtx := contextWithRole(ctx, username)
	newMd, ok := metadata.FromIncomingContext(newCtx)
	log.Println(newMd)
	return handler(
		newCtx,
		req,
	)
}

func authorizeUserFromToken(header []string) (string, bool) {
	if len(header) < 1 {
		return "", false
	}
	tokenStr := strings.TrimPrefix(header[0], "Bearer ")
	var token token.Token
	err := token.Decode(tokenStr)
	if err != nil {
		log.Printf("failed to decode token, err=%v", err)
		return "", false
	}
	if token.Secret != "my-secret" {
		return "", false
	}

	return token.Username, true
}

func contextWithRole(ctx context.Context, username string) context.Context {
	md := metadata.MD{}
	if username == "root" {
		md.Set(sayHelloRole, "false")
	}

	return metadata.NewIncomingContext(ctx, md)
}
