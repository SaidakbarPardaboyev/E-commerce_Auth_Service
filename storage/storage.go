package storage

import (
	"context"
	pb "ecommerce/genproto/auth_service"

	"github.com/golang/protobuf/ptypes/empty"
)

type IStorage interface {
	Close()
	Admin() IAdminsStorage
}

type IAdminsStorage interface {
	AddUser(context.Context, *pb.CreateUser) (*pb.User, error)
	GetUser(context.Context, *pb.PrimaryKey) (*pb.User, error)
	UpdateUser(context.Context, *pb.User) (*pb.User, error)
	DeleteUser(context.Context, *pb.PrimaryKey) (empty.Empty, error)
	GetUserProducts(context.Context, *pb.PrimaryKey) (*pb.Products, error)
	GetAllUsers(context.Context, *pb.UserFilter) (*pb.Users, error)
}
