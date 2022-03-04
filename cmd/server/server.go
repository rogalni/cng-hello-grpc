package main

import (
	"context"
	"fmt"
	"io"
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
func (us *userService) GetUsers(srv user.UserService_GetUsersServer) error {

	ctx := srv.Context()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		req, err := srv.Recv()
		if err == io.EOF {
			log.Println("exit")
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Received GetUsers stream with id %d", req.Id)
		resp := user.GetUsersResponse{
			Id:        req.Id,
			Username:  fmt.Sprintf("%djohn.doe", req.Id),
			Firstname: fmt.Sprintf("%djohn", req.Id),
			Lastname:  fmt.Sprintf("%ddoe", req.Id),
		}
		if err := srv.Send(&resp); err != nil {
			log.Printf("send error %v", err)
		}
	}
}
