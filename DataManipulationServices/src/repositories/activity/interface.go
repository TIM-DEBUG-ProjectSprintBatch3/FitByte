package activityRepository

import (
	"context"

	Entity "github.com/TimDebug/FitByte/src/model/entities/activity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ActivityRepositoryInterface interface {
	Create(ctx context.Context, pool *pgxpool.Pool, activity Entity.Activity) (activityId string, err error)
}
