package main

import (
	"context"
	"flag"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "be/internal/proto"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial err=%v", err)
	}

	defer conn.Close()
	client := pb.NewRouteGuideClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()
	feature, err := client.GetFeature(ctx, &pb.Point{
		Long: 1,
		Lat:  1,
	})
	log.Printf("feature=%v, err=%v", feature, err)

	features, err := client.ListFeatures(ctx, &pb.Rectangle{
		Lo: &pb.Point{
			Long: 1,
			Lat:  1,
		},
		Hi: &pb.Point{
			Long: 2,
			Lat:  2,
		},
	})
	for {
		feature, err := features.Recv()
		if err == io.EOF {
			log.Printf("eof")
			return
		}
		if err != nil {
			log.Fatalln("err=%v", err)
		}

		log.Printf("feature=%v, err=%v", feature, err)

	}
}
