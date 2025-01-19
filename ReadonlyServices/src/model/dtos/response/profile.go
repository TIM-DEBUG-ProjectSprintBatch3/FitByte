package response

import Entity "github.com/TimDebug/FitByte/src/model/entities"

type ProfileResponse struct {
	Preference *Entity.Preference `json:"preference"`
	WeightUnit *Entity.WeightUnit `json:"weightUnit"`
	HeightUnit *Entity.HeightUnit `json:"heightUnit"`
	Weight     *float32           `json:"weight"`
	Height     *float32           `json:"height"`
	Email      string             `json:"email"`
	Name       *string            `json:"name"`
	ImageUri   *string            `json:"imageUri"`
}
