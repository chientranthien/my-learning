package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	pb "be/internal/proto"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedRouteGuideServer
}

func (server) GetFeature(context.Context, *pb.Point) (*pb.Feature, error) {
	return &pb.Feature{
		Name: "tmp",
		Location: &pb.Point{
			Long: 1,
			Lat:  2,
		},
	}, nil
}
func (server) ListFeatures(_ *pb.Rectangle, stream pb.RouteGuide_ListFeaturesServer) error {
	for i := 0; i < 10; i++ {
		stream.Send(&pb.Feature{
			Name: fmt.Sprintf("name=%d", i),
			Location: &pb.Point{
				Long: int32(i),
				Lat:  int32(i),
			},
		})
	}

	return nil
}
func (server) RecordRoute(stream pb.RouteGuide_RecordRouteServer) error {

	startTime := time.Now()
	var pointCount, featureCount, distance int32
	for {
		point, err := stream.Recv()
		if err == io.EOF {
			stream.SendAndClose(&pb.RouteSummary{
				PointCount:   pointCount,
				FeatureCount: featureCount,
				Distance:     distance,
				ElapsedTime:  int32(time.Since(startTime).Seconds()),
			})
		}

		if err != nil {
			return err
		}

		pointCount++
		featureCount += 2
		distance += 3

		log.Printf("point =%v", point)
	}

	return nil
}
func (server) RouteChat(stream pb.RouteGuide_RouteChatServer) error {
	count := 0
	for {
		note, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}
		stream.Send(&pb.RouteNote{
			Location: &pb.Point{
				Long: note.Location.Long,
				Lat:  note.Location.Lat,
			},
			Message: fmt.Sprintf("%s:%d", note.Message, count),
		})
		count++
	}
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterRouteGuideServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
