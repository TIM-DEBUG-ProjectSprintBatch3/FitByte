package activityRepository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	Entity "github.com/rafitanujaya/go-fiber-template/src/model/entities/activity"
)

type ActivityRepositoryInterface interface {
	GetAll(ctx context.Context, pool *pgxpool.Pool, queryArgs []interface{}) ([]Entity.Activity, error)
}
