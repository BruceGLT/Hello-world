// client.go
package main

import (
	"Hello_world_thrift/gen-go/helloworld"
	"context"
	"log"

	"github.com/apache/thrift/lib/go/thrift"
)

func main() {
	tsocket, err := thrift.NewTSocket("localhost:50051")
	if err != nil {
		log.Fatalf("Failed to create transport: %v", err)
	}

	transportFactory := thrift.NewTTransportFactory()
	transport, err := transportFactory.GetTransport(tsocket)

	socketTransport, ok := transport.(*thrift.TSocket)
	if !ok {
		log.Fatalf("Failed to convert transport to *thrift.TSocket")
	}

	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	client := helloworld.NewHelloWorldClientFactory(socketTransport, protocolFactory)

	if err := socketTransport.Open(); err != nil {
		log.Fatalf("Failed to open transport: %v", err)
	}
	defer socketTransport.Close()

	ctx := context.Background()
	resp, err := client.SayHello(ctx, "World")
	if err != nil {
		log.Fatalf("Failed to get response: %v", err)
	}

	log.Printf("Response from server: %s", resp)
}
