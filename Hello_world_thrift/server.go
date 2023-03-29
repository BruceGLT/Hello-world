// server.go
package main

import (
	"Hello_world_thrift/gen-go/helloworld"
	"context"
	"log"

	"github.com/apache/thrift/lib/go/thrift"
)

type helloWorldHandler struct{}

func (h *helloWorldHandler) SayHello(ctx context.Context, name string) (string, error) {
	return "Hello, " + name, nil
}

func main() {
	transport, err := thrift.NewTServerSocket(":50051")
	if err != nil {
		log.Fatalf("Failed to create transport: %v", err)
	}

	processor := helloworld.NewHelloWorldProcessor(&helloWorldHandler{})
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTTransportFactory()

	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	log.Println("Starting server on :50051")
	if err := server.Serve(); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
