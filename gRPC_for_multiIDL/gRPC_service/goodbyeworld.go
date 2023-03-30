package main

import (
	"context"
	"log"
	"net"

	"gRPC_for_multiIDL/idl/goodbyeworld"

	"google.golang.org/grpc"
)

type server struct {
	goodbyeworld.UnimplementedGoodbyeWorldServer
}

func (s *server) SayGoodbye(ctx context.Context, in *goodbyeworld.GoodbyeRequest) (*goodbyeworld.GoodbyeResponse, error) {
	return &goodbyeworld.GoodbyeResponse{Message: "Goodbye, " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	goodbyeworld.RegisterGoodbyeWorldServer(grpcServer, &server{})

	log.Println("Starting server on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
