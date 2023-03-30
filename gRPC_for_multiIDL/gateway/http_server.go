package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	goodbyeworld "gRPC_for_multiIDL/idl/goodbyeworld"
	helloworld "gRPC_for_multiIDL/idl/helloworld"

	"google.golang.org/grpc"
)

func main() {
	http.HandleFunc("/", handleRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		http.Error(w, "Failed to dial gRPC server", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	idl := r.URL.Query().Get("idl")
	name := r.URL.Query().Get("name")

	var result string

	switch idl {
	case "hello":
		client := helloworld.NewHelloWorldClient(conn)
		response, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: name})
		if err != nil {
			http.Error(w, "Failed to call HelloWorld service", http.StatusInternalServerError)
			return
		}
		result = response.Message
	case "goodbye":
		client := goodbyeworld.NewGoodbyeWorldClient(conn)
		response, err := client.SayGoodbye(context.Background(), &goodbyeworld.GoodbyeRequest{Name: name})
		if err != nil {
			http.Error(w, "Failed to call GoodbyeWorld service", http.StatusInternalServerError)
			return
		}
		result = response.Message
	default:
		http.Error(w, "Unknown IDL", http.StatusBadRequest)
		return
	}

	response := map[string]string{"result": result}
	json.NewEncoder(w).Encode(response)
}
