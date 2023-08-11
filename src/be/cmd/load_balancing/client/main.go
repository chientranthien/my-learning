package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"

	pb "be/internal/proto"
)

const (
	schema      = "schema"
	serviceName = "service_name"
)

var addrs = []string{"localhost:50001", "localhost:50002"}

func main() {
	firstConn, err := grpc.Dial(fmt.Sprintf("%s:///%s", schema, serviceName), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect, err=%v", err)
		return
	}
	defer firstConn.Close()

	makeRpc(firstConn)

	log.Println("-------------------")

	secondConn, err := grpc.Dial(
		fmt.Sprintf("%s:///%s", schema, serviceName),
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`), // This sets the initial balancing policy.
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect, err=%v", err)
	}
	makeRpc(secondConn)
}

func sayHello(client pb.GreeterClient) {
	ctx := context.Background()
	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "test name"})
	if err != nil {
		log.Printf("failed, err=%v", err)
		return
	}

	log.Println(resp.Message)
}

func makeRpc(conn *grpc.ClientConn) {
	c := pb.NewGreeterClient(conn)
	for i := 0; i < 10; i++ {
		sayHello(c)
	}
}

type resolverBuilder struct{}

func (r *resolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	myResolver := &myResolver{
		target: target,
		cc:     cc,
		addressStored: map[string][]string{
			serviceName: addrs,
		},
	}

	myResolver.start()
	return myResolver, nil
}
func (r *resolverBuilder) Scheme() string {
	return schema
}

type myResolver struct {
	target        resolver.Target
	cc            resolver.ClientConn
	addressStored map[string][]string
}

func (r *myResolver) start() {
	addrStored := r.addressStored[r.target.Endpoint()]
	addrs := make([]resolver.Address, len(addrStored))
	for i, s := range addrStored {
		addrs[i] = resolver.Address{Addr: s}
	}

	r.cc.UpdateState(resolver.State{
		Addresses: addrs,
	})
}
func (r *myResolver) ResolveNow(options resolver.ResolveNowOptions) {
}

func (r *myResolver) Close() {
}

func init() {
	resolver.Register(&resolverBuilder{})
}
