package main

import (
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	"be/internal/greeter_server"
	pb "be/internal/proto"
)

func main() {
	addrs := []string{":5001", ":5002"}
	for _, addr := range addrs {
		t := addr
		healthServer := health.NewServer()
		server := grpc.NewServer()
		healthpb.RegisterHealthServer(server, healthServer)
		go func() {
			for {

				if t == ":5001" {
					healthServer.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
				} else {
					healthServer.SetServingStatus("", healthpb.HealthCheckResponse_NOT_SERVING)
				}
				time.Sleep(5 * time.Second)
			}
		}()

		lis, err := net.Listen("tcp", t)
		if err != nil {
			log.Fatalf("err=%v", err)
		}
		pb.RegisterGreeterServer(server, greeter_server.GreeterServer{Name: t})
		go server.Serve(lis)
	}

	select {}
}
