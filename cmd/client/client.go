package main

import (
	"context"
	"log"
	"os"

	"github.com/rogalni/cng-hello-grpc/api/gen/chat"
	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	h := os.Getenv("GRPC_SERVER_HOST")
	p := os.Getenv("GRPC_SERVER_PORT")
	conn, err := grpc.Dial(h+":"+p, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	response, err := c.SayHello(context.Background(), &chat.Message{Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

}
