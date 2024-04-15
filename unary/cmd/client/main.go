package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/gustavodalves/go-gRPC/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println(err)
	}

	client := pb.NewProductClient(conn)

	request := &pb.ProductRequest{
		Id: uuid.NewString(),
	}

	response, err := client.Get(context.Background(), request)

	if err != nil {
		panic(err)
	}

	log.Println(response)

}
