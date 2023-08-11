package main

import (
	"flag"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"

	"be/internal/greeter_server"
	pb "be/internal/proto"
)

var (
	addr = flag.String("addr", ":50001", "")
)

func startServer(addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to start, addr=%v, err=%v", addr, err)
		return
	}

	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &greeter_server.GreeterServer{Name: addr})
	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("failed to start, addr=%v, err=%v", addr, err)
		return
	}
}
func main() {
	flag.Parse()
	wg := &sync.WaitGroup{}
	startServer(*addr)
	wg.Wait()
}
