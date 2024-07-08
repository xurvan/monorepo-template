package main

import (
	"context"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"

	"github.com/xurvan/monorepo-template/apps/perfect/config"
	"github.com/xurvan/monorepo-template/apps/perfect/gen/db"
	pb "github.com/xurvan/monorepo-template/apps/perfect/gen/pb"
	"github.com/xurvan/monorepo-template/apps/perfect/internal/repository"
	"github.com/xurvan/monorepo-template/apps/perfect/internal/services"
)

func main() {
	cfg := config.LoadConfig()

	conn, err := pgx.Connect(context.Background(), cfg.DatabaseUrl)
	if err != nil {
		panic(err)
	}
	queries := db.New(conn)

	repo := repository.NewUserRepository(queries)
	srv := &services.UserService{
		Repository: repo,
	}
	twirpHandler := pb.NewUserServiceServer(srv)

	log.Println("Serving HTTP on 0.0.0.0:8080")
	err = http.ListenAndServe(":8080", twirpHandler)
	if err != nil {
		panic(err)
	}
}
