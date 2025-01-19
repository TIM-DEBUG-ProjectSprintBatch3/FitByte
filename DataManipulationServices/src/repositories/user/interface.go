package userRepository

import (
	"context"

	Entity "github.com/TimDebug/FitByte/src/model/entities/user"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepositoryInterface interface {
	Login(ctx context.Context, pool *pgxpool.Pool, body *Entity.User) ([]Entity.User, error)
	CreateUser(ctx context.Context, pool *pgxpool.Pool, user Entity.User) (userId string, err error)
	FindById(ctx context.Context, id string) (*Entity.User, error)
	Update(ctx context.Context, user Entity.User) (string, error)
}
