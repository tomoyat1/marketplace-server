package main

import (
	"context"
	//"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/reflection"

	pb "github.com/tomoyat1/marketplace-server/marketplace"
)

type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) ListItems(ctx context.Context, in *pb.ListItemsRequest) (*pb.ListItemsReply, error) {
	items := []*pb.Item{
		&pb.Item{
			Id:    "a17e49a5-b915-445e-ae4d-fbc327c9d8d7",
			Name:  "Site Reliability Engineering",
			Price: 4000,
		},
		&pb.Item{
			Id:    "5718a9af-849c-45ab-a5b9-f78c185adbf3",
			Name:  "ErgoDox EZ",
			Price: 30000,
		},
	}
	if pr, ok := peer.FromContext(ctx); ok {
		log.Printf("{\"method\":\"ListItems\",\"client:\":\"%s\"}", pr.Addr.String())
	}
	return &pb.ListItemsReply{Items: items}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterItemsServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
