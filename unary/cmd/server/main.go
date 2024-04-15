package main

import (
	"context"
	"fmt"
	"net"

	"github.com/gustavodalves/go-gRPC/pb"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedProductServer
}

func (s *Server) Get(ctx context.Context, in *pb.ProductRequest) (*pb.ProductResponse, error) {
	return &pb.ProductResponse{
		Name:  "GUSTAVO",
		Value: 90.0,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()

	pb.RegisterProductServer(server, &Server{})

	if err = server.Serve(listener); err != nil {
		fmt.Println("failed to server", err)
	}
}
