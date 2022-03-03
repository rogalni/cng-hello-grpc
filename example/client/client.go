package main

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/rogalni/cng-hello-grpc/gen/user/v1"
	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	h, e := os.LookupEnv("GRPC_SERVER_HOST")
	if !e {
		h = "localhost"
	}
	p, e := os.LookupEnv("GRPC_SERVER_PORT")
	if !e {
		p = "9000"
	}
	conn, err := grpc.Dial(h+":"+p, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	uc := user.NewUserServiceClient(conn)
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			runGetUser(uc, i)
		}(i)
	}
	wg.Wait()
}

func runGetUser(uc user.UserServiceClient, i int) {
	response, err := uc.GetUser(context.Background(), &user.GetUserRequest{
		Id: int64(i),
	})
	if err != nil {
		log.Fatalf("Error when calling GetUser: %s", err)
	}
	log.Printf("Response from server %v", response)

}
