package main

import (
	"context"
	"github.com/GonnaFlyMethod/grpc-demo/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("can't connect to the server, err: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		if err := conn.Close(); err != nil {
			log.Println("can't close client's connection")
		}
	}(conn)

	c := chat.NewChatServiceClient(conn)

	message := &chat.Message{
		Body: "Hello from the client!",
	}

	response, err := c.SayHello(context.Background(), message)
	if err != nil {
		log.Fatalf("error when calling SayHello, err: %v", err)
	}

	log.Printf("response from server: %s ", response.Body)
}
