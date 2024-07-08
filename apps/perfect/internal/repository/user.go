package repository

import (
	"context"

	"github.com/jinzhu/copier"

	"github.com/xurvan/monorepo-template/apps/perfect/gen/db"
	pb "github.com/xurvan/monorepo-template/apps/perfect/gen/pb"
)

type UserRepository struct {
	qr *db.Queries
}

func NewUserRepository(qr *db.Queries) (rep *UserRepository) {
	return &UserRepository{
		qr: qr,
	}
}

func (r *UserRepository) List(ctx context.Context, req *pb.ListUsersRequest) (res *pb.ListUsersResponse, err error) {
	params := db.ListUsersParams{}
	err = copier.Copy(&params, req)

	rows, err := r.qr.ListUsers(ctx, params)
	if err != nil {
		return nil, err
	}

	res = &pb.ListUsersResponse{}
	for _, row := range rows {
		var user pb.User

		err = copier.Copy(&user, row)
		if err != nil {
			return nil, err
		}

		res.Users = append(res.Users, &user)
	}

	return res, err
}

func (r *UserRepository) Create(ctx context.Context, req *pb.User) (user *pb.User, err error) {
	params := db.CreateUserParams{}
	err = copier.Copy(&params, req)
	if err != nil {
		return nil, err
	}

	row, err := r.qr.CreateUser(ctx, params)
	if err != nil {
		return nil, err
	}

	user = &pb.User{}
	err = copier.Copy(&user, row)
	return user, err
}
