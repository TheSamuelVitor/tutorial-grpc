package main

import (
	"context"
	"hello-grpc/pb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	cliente := pb.NewHelloClient(conn)
	req := &pb.HelloRequest{
		Name: "Samuel",
	}

	res, err := cliente.SayHello(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)
}
