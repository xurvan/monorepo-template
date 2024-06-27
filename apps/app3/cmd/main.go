package main

import (
	"context"
	"log"
	"net"

	pb "github.com/xurvan/monorepo-template/apps/app3/gen/pb"
	"google.golang.org/grpc"
)

type HelloService struct {
	pb.UnimplementedHelloServiceServer
}

func (s *HelloService) SayHello(ctx context.Context, req *pb.SayHelloRequest) (res *pb.SayHelloResponse, err error) {
	res = &pb.SayHelloResponse{
		Message: "Hello, " + req.Name,
	}

	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &HelloService{})

	log.Println("Serving gRPC on 0.0.0.0:9090")
	log.Fatal(s.Serve(lis))
}
