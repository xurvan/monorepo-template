package services

import (
	"context"

	google "google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/xurvan/monorepo-template/apps/perfect/gen/pb"
	"github.com/xurvan/monorepo-template/apps/perfect/internal/repository"
)

type UserService struct {
	Repository *repository.UserRepository
}

func (s *UserService) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (res *pb.ListUsersResponse, err error) {
	return s.Repository.List(ctx, req)
}

func (s *UserService) GetUser(ctx context.Context, req *pb.IdRequest) (res *pb.User, err error) {
	return nil, nil
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.User) (res *pb.User, err error) {
	return s.Repository.Create(ctx, req)
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.User) (res *pb.User, err error) {
	return nil, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.IdRequest) (res *google.Empty, err error) {
	res = &google.Empty{}

	return res, err
}
