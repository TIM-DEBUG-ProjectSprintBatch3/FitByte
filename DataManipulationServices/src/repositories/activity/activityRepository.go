package activityRepository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	Entity "github.com/rafitanujaya/go-fiber-template/src/model/entities/activity"
	"github.com/samber/do/v2"
)

type ActivityRepository struct{}

func NewActivityRepository() ActivityRepositoryInterface {
	return &ActivityRepository{}
}

func NewActivityRepositoryInject(i do.Injector) (ActivityRepositoryInterface, error) {
	return NewActivityRepository(), nil
}

func (ar *ActivityRepository) Create(ctx context.Context, pool *pgxpool.Pool, activity Entity.Activity) (string, error) {
	query := `
		INSERT INTO activities (
		user_id, 
		done_at, 
		duration_in_minutes, 
		calories_burned, 
		activity_type, 
		created_at,
		updated_at
	) VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id;
	`
	var activityId string
	err := pool.QueryRow(
		ctx,
		query,
		activity.UserId,
		activity.DoneAt,
		activity.DurationInMinutes,
		activity.CaloriesBurned,
		activity.ActivityType,
		activity.CreatedAt,
		activity.UpdatedAt,
	).Scan(&activityId)
	if err != nil {
		return "", err
	}

	return activityId, nil
}
