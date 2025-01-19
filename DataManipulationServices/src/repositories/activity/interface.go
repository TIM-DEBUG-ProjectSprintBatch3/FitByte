package activityRepository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	Entity "github.com/rafitanujaya/go-fiber-template/src/model/entities/activity"
)

type ActivityRepositoryInterface interface {
	Create(ctx context.Context, pool *pgxpool.Pool, activity Entity.Activity) (activityId string, err error)
}
