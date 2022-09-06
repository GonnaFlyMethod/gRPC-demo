package main

import (
	"github.com/GonnaFlyMethod/grpc-demo/chat"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to run server, err: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, &s)

	log.Println("running gRPC server on localhost:9000")
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to run gRPC server, err: %v", err)
	}
}
