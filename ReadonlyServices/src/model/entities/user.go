package Entity

import "time"

const (
	PreferenceCardio Preference = "CARDIO"
	PreferenceWeight Preference = "WEIGHT"

	WeightUnitKg  WeightUnit = "KG"
	WeightUnitLbs WeightUnit = "LBS"

	HeightUnitCm   HeightUnit = "CM"
	HeightUnitInch HeightUnit = "INCH"
)

type User struct {
	Id           string
	Email        string
	PasswordHash string
	Preference   *Preference
	WeightUnit   *WeightUnit
	HeightUnit   *HeightUnit
	Weight       *float32
	Height       *float32
	Name         *string
	ImageUri     *string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Preference string

type WeightUnit string

type HeightUnit string
