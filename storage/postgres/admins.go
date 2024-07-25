package postgres

import (
	"context"
	"ecommerce/pkg/logger"
	"ecommerce/storage"

	pb "ecommerce/genproto/auth_service"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jackc/pgx/v5/pgxpool"
)

type adminRepo struct {
	db  *pgxpool.Pool
	log logger.ILogger
}

func NewAdminRepo(db *pgxpool.Pool, log logger.ILogger) storage.IAdminsStorage {
	return &adminRepo{
		db:  db,
		log: log,
	}
}

func (a *adminRepo) AddUser(context.Context, *pb.CreateUser) (*pb.User, error) {

}

func (a *adminRepo) GetUser(context.Context, *pb.PrimaryKey) (*pb.User, error)
func (a *adminRepo) UpdateUser(context.Context, *pb.User) (*pb.User, error)
func (a *adminRepo) DeleteUser(context.Context, *pb.PrimaryKey) (empty.Empty, error)
func (a *adminRepo) GetUserProducts(context.Context, *pb.PrimaryKey) (*pb.Products, error)
func (a *adminRepo) GetAllUsers(context.Context, *pb.UserFilter) (*pb.Users, error)
