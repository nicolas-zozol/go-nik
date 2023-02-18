package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	api "github.com/nicolas-zozol/go-nik/proto/nik"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

func main() {
	var grpcServerEndpoint = "localhost:8080"
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := api.RegisterGreeterHandlerFromEndpoint(context.Background(),
		mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Println("Listening on port 8081")
	port := ":8081"
	http.ListenAndServe(port, mux)
}
