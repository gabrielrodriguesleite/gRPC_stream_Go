package main

import (
	"context"
	"flag"
	"fmt"
	pb "grpc_stream_go/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "Server port")
)

type server struct {
	pb.UnimplementedServiceUpvoteServer
}

func (s *server) Upvote(ctx context.Context, in *pb.RequestChose) (*pb.ResponseOption, error) {

	log.Printf("Received :%s, %s, %s, %v", in.UserName, in.Email, in.Title, in.OptionChosed)

	return &pb.ResponseOption{Id: 0, Title: "Create first option"}, nil
}

func main() {

	// parse flags from standard input
	flag.Parse()

	// open tcp port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// instantiate a grpc server
	s := grpc.NewServer()

	// set a protobuff service on grpc server
	pb.RegisterServiceUpvoteServer(s, &server{})

	log.Printf("Server listening at %v", lis.Addr())

	// open server
	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
