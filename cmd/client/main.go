package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	"my-project/api"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := api.NewApiClient(conn)
	response, err := c.Hello(context.TODO(), &api.HelloRequest{})
	if err != nil {
		fmt.Println("err", err)
		return
	}

	fmt.Println(response.Code)
	fmt.Println(response.Msg)
}
