package main

import (
	"context"
	"hello-grpc/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedHelloServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello, " + in.GetName(),
	}, nil
}

func main() {
	println("running gRPC server")

	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	pb.RegisterHelloServer(server, &Server{})
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
