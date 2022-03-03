package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/rogalni/cng-hello-grpc/gen/user/v1"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, &userService{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

type userService struct {
	user.UnimplementedUserServiceServer
}

func (us *userService) GetUser(ctx context.Context, r *user.GetUserRequest) (*user.GetUserResponse, error) {
	log.Printf("Received GetUser request %v", r)
	return &user.GetUserResponse{
		Id:        r.Id,
		Username:  fmt.Sprintf("%djohn.doe", r.Id),
		Firstname: fmt.Sprintf("%djohn", r.Id),
		Lastname:  fmt.Sprintf("%ddoe", r.Id),
	}, nil
}
