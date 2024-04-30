package main

import (
	"context"
	"fmt"
	proto "github.com/provider-go/pkg/grpc/pb"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedHelloServer
}

func (s *server) Say(ctx context.Context, req *proto.SayRequest) (*proto.SayResponse, error) {
	fmt.Println("接到请求:", req.Name)
	return &proto.SayResponse{Message: "Hello " + req.Name}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8001")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	proto.RegisterHelloServer(s, &server{})

	defer func() {
		s.Stop()
		err := listen.Close()
		if err != nil {
			return
		}
	}()

	fmt.Println("Serving 8001...")
	err = s.Serve(listen)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
