package request

type UserRegister struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UpdateProfile struct {
	Preference string  `json:"preference"`
	WeightUnit string  `json:"weightUnit"`
	HeightUnit string  `json:"heightUnit"`
	Weight     float32 `json:"weight"`
	Height     float32 `json:"height"`
	Name       *string `json:"name"`
	ImageUri   *string `json:"imageUri"`
}
