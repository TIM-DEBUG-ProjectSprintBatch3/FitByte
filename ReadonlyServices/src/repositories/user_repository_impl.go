package repositories

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	Entity "github.com/rafitanujaya/go-fiber-template/src/model/entities"
	"github.com/samber/do/v2"
)

type UserRepositoryImpl struct {
	db *pgxpool.Pool
}

func (u UserRepositoryImpl) FindAll(
	ctx context.Context,
	inIds []string,
) ([]Entity.User, error) {
	query := `
		SELECT id, email, preference, weight_unit, height_unit, weight, height, name, image_uri, created_at, updated_at
		FROM Users
		WHERE id = ANY($1::text[]);
	`
	rows, err := u.db.Query(ctx, query, inIds)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err) // Return wrapped errors
	}
	defer rows.Close()

	var users []Entity.User
	for rows.Next() {
		var user Entity.User
		err := rows.Scan(&user.Id, &user.Email, &user.Preference, &user.WeightUnit, &user.HeightUnit, &user.Weight, &user.Height, &user.Name, &user.ImageUri, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u UserRepositoryImpl) FindById(
	ctx context.Context,
	id string,
) (*Entity.User, error) {
	row := u.db.QueryRow(
		ctx,
		`
			SELECT id, email, preference, weight_unit, height_unit, weight, height, name, image_uri, created_at, updated_at 
			FROM Users 
			WHERE id = $1`,
		id,
	)

	var user Entity.User
	err := row.Scan(&user.Id, &user.Email, &user.Preference, &user.WeightUnit, &user.HeightUnit, &user.Weight, &user.Height, &user.Name, &user.ImageUri, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func NewUserRepositoryImpl(db *pgxpool.Pool) UserRepository {
	return UserRepositoryImpl{
		db: db,
	}
}

func NewUserRepositoryImplInject(i do.Injector) (UserRepository, error) {
	return NewUserRepositoryImpl(
		do.MustInvoke[*pgxpool.Pool](i),
	), nil
}
