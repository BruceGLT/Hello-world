// server.go
package main

import (
	"context"
	"log"
	"net"

	"Hello_world_protobuf/helloworld"

	"google.golang.org/grpc"
)

type server struct {
	helloworld.UnimplementedHelloWorldServer
}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloResponse, error) {
	return &helloworld.HelloResponse{Message: "Hello, " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	helloworld.RegisterHelloWorldServer(grpcServer, &server{})

	log.Println("Starting server on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
