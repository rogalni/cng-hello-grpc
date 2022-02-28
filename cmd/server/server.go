package main

import (
	"context"
	"log"
	"net"

	"github.com/rogalni/cng-hello-grpc/api/gen/chat"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := Server{}
	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *chat.Message) (*chat.Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &chat.Message{Body: "Hello From the Server!"}, nil
}
