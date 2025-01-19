package userRepository

import (
	"context"

	Entity "github.com/TimDebug/FitByte/src/model/entities/user"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, pool *pgxpool.Pool, user Entity.User) (userId string, err error)
}
