package userRepository

import (
	Entity "github.com/TimDebug/FitByte/src/model/entities/user"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepositoryInterface interface {
	Login(ctx *fiber.Ctx, pool *pgxpool.Pool, body *Entity.User) ([]Entity.User, error)
	CreateUser(ctx *fiber.Ctx, pool *pgxpool.Pool, user Entity.User) (userId string, err error)
	FindById(ctx *fiber.Ctx, id string) (*Entity.User, error)
	Update(ctx *fiber.Ctx, user Entity.User) (string, error)
}
