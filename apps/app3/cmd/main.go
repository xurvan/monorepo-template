package main

import (
	"context"
	"log"
	"net/http"

	pb "github.com/xurvan/monorepo-template/apps/app3/gen/pb"
)

type HelloService struct{}

func (s *HelloService) SayHello(ctx context.Context, req *pb.SayHelloRequest) (res *pb.SayHelloResponse, err error) {
	res = &pb.SayHelloResponse{
		Message: "Hello, " + req.Name,
	}

	return res, nil
}

func main() {
	server := &HelloService{}
	twirpHandler := pb.NewHelloServiceServer(server)

	log.Println("Serving HTTP on 0.0.0.0:8080")
	err := http.ListenAndServe(":8080", twirpHandler)
	if err != nil {
		panic(err)
	}
}
