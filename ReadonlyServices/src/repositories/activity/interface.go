package activityRepository

import (
	"context"

	Entity "github.com/TimDebug/FitByte/src/model/entities/activity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ActivityRepositoryInterface interface {
	GetAll(ctx context.Context, pool *pgxpool.Pool, queryArgs []interface{}) ([]Entity.Activity, error)
}
