package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/rogalni/cng-hello-grpc/gen/user/v1"
	"google.golang.org/grpc"
)

var (
	iterationCount = 100
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
	s := time.Now()
	runAsyncGetUser(uc)
	fmt.Printf("Get single user async took %dms\n", time.Since(s).Milliseconds())
	s = time.Now()
	runGetUsersStream(uc)

	fmt.Printf("Get users in bidirectional stream took %dms\n", time.Since(s).Milliseconds())
}

func runAsyncGetUser(uc user.UserServiceClient) {
	var wg sync.WaitGroup
	for i := 0; i < iterationCount; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			runGetUser(uc, i)
		}(i)
	}
	wg.Wait()
	log.Printf("Done Get User Async")
}

func runGetUser(uc user.UserServiceClient, i int) {
	res, err := uc.GetUser(context.Background(), &user.GetUserRequest{
		Id: int64(i),
	})
	if err != nil {
		log.Fatalf("Error when calling GetUser: %s", err)
	}
	log.Printf("Response from server %v", res)

}

func runGetUsersStream(uc user.UserServiceClient) {

	done := make(chan bool)
	stm, err := uc.GetUsers(context.Background())
	ctx := stm.Context()
	if err != nil {
		log.Fatal("Failed to open GetUsers stream")
	}

	go func() {
		for i := 1; i <= iterationCount; i++ {
			req := &user.GetUsersRequest{Id: int64(i)}
			if err := stm.Send(req); err != nil {
				log.Fatalf("can not send %v", err)
			}
			log.Printf("UserId: %d sent", req.Id)
		}
		if err := stm.CloseSend(); err != nil {
			log.Println(err)
		}
	}()

	go func() {
		for {
			res, err := stm.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}
			log.Printf("User received %v", res)
		}
	}()

	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			log.Println(err)
		}
	}()

	<-done
	log.Println("finished with getting users in stream")
}
