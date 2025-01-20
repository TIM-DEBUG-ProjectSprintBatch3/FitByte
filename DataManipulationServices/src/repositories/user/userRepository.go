package userRepository

import (
	"fmt"

	Entity "github.com/TimDebug/FitByte/src/model/entities/user"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/samber/do/v2"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepositoryInterface {
	return &UserRepository{
		db: db,
	}
}

func NewUserRepositoryInject(i do.Injector) (UserRepositoryInterface, error) {
	return NewUserRepository(
		do.MustInvoke[*pgxpool.Pool](i),
	), nil
}

func (ur *UserRepository) Login(ctx *fiber.Ctx, pool *pgxpool.Pool, body *Entity.User) ([]Entity.User, error) {
	query := `
		SELECT id, email, password_hash
		FROM Users
		WHERE email = $1
	`
	rows, err := pool.Query(ctx.Context(), query, body.Email)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var users []Entity.User
	for rows.Next() {
		var user Entity.User
		if err := rows.Scan(&user.Id, &user.Email, &user.PasswordHash); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *UserRepository) CreateUser(ctx *fiber.Ctx, pool *pgxpool.Pool, user Entity.User) (userId string, err error) {
	query := `INSERT INTO users(email, password_hash) VALUES($1, $2) RETURNING id`
	fmt.Printf("Email: %s, Password %s", user.Email, user.PasswordHash)

	row := pool.QueryRow(ctx.Context(), query, user.Email, user.PasswordHash)
	err = row.Scan(&userId)
	if err != nil {
		return "", err
	}
	return userId, nil

}

func (ur *UserRepository) FindById(ctx *fiber.Ctx, id string) (*Entity.User, error) {
	row := ur.db.QueryRow(
		ctx.Context(),
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

func (ur *UserRepository) Update(ctx *fiber.Ctx, user Entity.User) (string, error) {
	_, err := ur.db.Exec(
		ctx.Context(),
		`UPDATE users SET 
			preference= $1, 
			weight_unit=$2, 
			height_unit=$3, 
			weight=$4, 
			height=$5,
			name=$6,
			image_uri=$7
		WHERE id = $8`,
		user.Preference,
		user.WeightUnit,
		user.HeightUnit,
		user.Weight,
		user.Height,
		user.Name,
		user.ImageUri,
		user.Id,
	)
	return user.Id, err
}
