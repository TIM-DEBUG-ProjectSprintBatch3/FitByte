package Entity

import "time"

type User struct {
	Id           string
	Email        string
	PasswordHash string
	Preference   *string
	WeightUnit   *string
	HeightUnit   *string
	Weight       *float32
	Height       *float32
	Name         *string
	ImageUri     *string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
