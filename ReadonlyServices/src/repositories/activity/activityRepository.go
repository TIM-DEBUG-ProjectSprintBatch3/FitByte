package activityRepository

import (
	"context"

	Entity "github.com/TimDebug/FitByte/src/model/entities/activity"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/samber/do/v2"
)

type ActivityRepository struct{}

func NewActivityRepository() ActivityRepositoryInterface {
	return &ActivityRepository{}
}

func NewActivityRepositoryInject(i do.Injector) (ActivityRepositoryInterface, error) {
	return NewActivityRepository(), nil
}

func (ar *ActivityRepository) GetAll(ctx context.Context, pool *pgxpool.Pool, queryArgs []interface{}) ([]Entity.Activity, error) {
	query := `
		SELECT id, activity_type, done_at, duration_in_minutes, calories_burned, created_at
		FROM activities
		WHERE
			user_id = $1
			AND ($2::TEXT IS NULL OR activity_type = $2)
			AND ($3::TIMESTAMP IS NULL OR done_at >= $3)
			AND ($4::TIMESTAMP IS NULL OR done_at <= $4)
			AND ($5::NUMERIC IS NULL OR calories_burned >= $5)
			AND ($6::NUMERIC IS NULL OR calories_burned <= $6)
		LIMIT $7 OFFSET $8;
	`
	rows, err := pool.Query(ctx, query, queryArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var activities []Entity.Activity
	for rows.Next() {
		var activity Entity.Activity
		if err := rows.Scan(
			&activity.ActivityId,
			&activity.ActivityType,
			&activity.DoneAt,
			&activity.DurationInMinutes,
			&activity.CaloriesBurned,
			&activity.CreatedAt,
		); err != nil {
			return activities, err
		}
		activities = append(activities, activity)
	}
	return activities, nil
}
