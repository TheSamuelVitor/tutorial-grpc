package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	v1 "tutorialgrpc/gen/product/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	port     = flag.String("port", "8081", "this is the port the server will use")
	endpoint = flag.String("endpoint", "localhost:8080", "endpoint is the gRPC server's adress")
)

func main() {
	flag.Parse()
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := v1.RegisterProductServiceHandlerFromEndpoint(ctx, mux, *endpoint, opts)
	if err != nil {
		return err
	}
	return http.ListenAndServe(":"+*port, mux)
}
