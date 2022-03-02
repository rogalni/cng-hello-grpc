package main

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/rogalni/cng-hello-grpc/api/gen/chat"
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

	c := chat.NewChatServiceClient(conn)
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			runHello(c, i)
		}(i)
	}
	wg.Wait()
}

func runHello(c chat.ChatServiceClient, i int) {

	response, err := c.SayHello(context.Background(), &chat.Message{
		Body:  "Hello From Client!",
		Index: int64(i),
	})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server for iteration %d: %s", i, response.Body)

}
