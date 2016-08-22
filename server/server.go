package main

import (
	"log"
	"net"

	pb "github.com/hassanriaz/grpc/server/add"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50000"
)

type addServer struct{}

func (s *addServer) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	log.Printf("Adding: %v and %v", in.First, in.Second)
	sum := in.First + in.Second
	return &pb.AddResponse{Sum: sum}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAddServiceServer(s, &addServer{})
	s.Serve(lis)
}
