package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/utrack/clay/v3/log"
	"github.com/utrack/clay/v3/transport"
	"github.com/utrack/clay/v3/transport/middlewares/mwgrpc"
	"github.com/utrack/clay/v3/transport/server"
	"golang.org/x/net/context"

	pb "github.com/xurvan/monorepo-template/apps/clay/gen/pb"
)

type HelloService struct {
	pb.UnimplementedHelloServiceServer
}

func (s *HelloService) SayHello(ctx context.Context, req *pb.SayHelloRequest) (res *pb.SayHelloResponse, err error) {
	res = &pb.SayHelloResponse{
		Message: "Hello " + req.Name,
	}

	return res, nil
}

func (s *HelloService) GetDescription() transport.ServiceDesc {
	return pb.NewHelloServiceServiceDesc(s)
}

func main() {
	router := chi.NewRouter()
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("pong"))
		if err != nil {
			panic(err)
		}
	})

	svc := &HelloService{}
	srv := server.NewServer(
		8080,
		server.WithHTTPMux(router),
		server.WithGRPCUnaryMiddlewares(
			mwgrpc.UnaryPanicHandler(log.Default),
		),
	)
	err := srv.Run(svc)
	if err != nil {
		panic(err)
	}
}
