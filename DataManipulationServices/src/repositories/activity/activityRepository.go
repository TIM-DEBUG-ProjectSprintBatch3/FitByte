package activityRepository

import (
	"context"
	"strconv"

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

func (ar *ActivityRepository) GetValidCaloriesFactors(ctx context.Context, pool *pgxpool.Pool, activityId, userId string) (*Entity.CaloriesFactor, error) {
	query := `SELECT activity_type, duration_in_minutes FROM activities WHERE id = $1 AND user_id = $2`

	rows, err := pool.Query(ctx, query, activityId, userId)
	if err != nil {
		return &Entity.CaloriesFactor{}, err
	}

	var factor Entity.CaloriesFactor
	for rows.Next() {
		err = rows.Scan(&factor.ActivityType, &factor.DurationInMinutes)
		if err != nil {
			return nil, err
		}
	}

	return &factor, nil
}

func (ar *ActivityRepository) Update(ctx context.Context, pool *pgxpool.Pool, activity Entity.Activity) error {
	// ` // Query:
	// UPDATE activities
	// SET
	// 	done_at = $1,
	// 	duration_in_minutes = $2,
	// 	calories_burned = $3,
	// 	activity_type = $4,
	// 	updated_at = $5
	// WHERE id = $6;`
	query := "UPDATE activities SET "
	args := []interface{}{}
	argId := 1

	if !activity.DoneAt.IsZero() {
		query += "done_at = $" + strconv.Itoa(argId) + ", "
		args = append(args, activity.DoneAt)
		argId++
	}
	if activity.DurationInMinutes != 0 {
		query += "duration_in_minutes = $" + strconv.Itoa(argId) + ", "
		args = append(args, activity.DurationInMinutes)
		argId++
	}
	if activity.CaloriesBurned != 0 {
		query += "calories_burned = $" + strconv.Itoa(argId) + ", "
		args = append(args, activity.CaloriesBurned)
		argId++
	}
	if activity.ActivityType != "" {
		query += "activity_type = $" + strconv.Itoa(argId) + ", "
		args = append(args, activity.ActivityType)
		argId++
	}

	query += "updated_at = $" + strconv.Itoa(argId) + " WHERE id = $" + strconv.Itoa(argId+1) + ";"
	args = append(args, activity.UpdatedAt, activity.ActivityId)

	_, err := pool.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
