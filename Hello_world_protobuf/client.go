// client.go
package main

import (
	"Hello_world_protobuf/helloworld"
	"context"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	client := helloworld.NewHelloWorldClient(conn)

	resp, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "World"})
	if err != nil {
		log.Fatalf("Failed to get response: %v", err)
	}

	log.Printf("Response from server: %s", resp.GetMessage())

}
