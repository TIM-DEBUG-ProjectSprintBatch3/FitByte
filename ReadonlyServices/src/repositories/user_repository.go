package repositories

import (
	"context"
	Entity "github.com/rafitanujaya/go-fiber-template/src/model/entities"
)

type UserRepository interface {
	FindAll(ctx context.Context, inIds []string) ([]Entity.User, error)
	FindById(ctx context.Context, id string) (*Entity.User, error)
}
