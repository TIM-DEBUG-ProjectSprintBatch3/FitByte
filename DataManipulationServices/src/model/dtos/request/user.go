package request

type UserRegister struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type UpdateProfile struct {
	Preference string  `json:"preference"`
	WeightUnit string  `json:"weightUnit"`
	HeightUnit string  `json:"heightUnit"`
	Weight     float32 `json:"weight"`
	Height     float32 `json:"height"`
	Name       *string `json:"name"`
	ImageUri   *string `json:"imageUri" validate:"uri"`
}

type ImageUri struct {
	ImageUri string `json:"imageUri" validate:"uri"`
}

type UpdateProfileCustom struct {
	Preference string       `json:"preference"`
	WeightUnit string       `json:"weightUnit"`
	HeightUnit string       `json:"heightUnit"`
	Weight     float32      `json:"weight"`
	Height     float32      `json:"height"`
	Name       CustomString `json:"name"`
	ImageUri   CustomString `json:"imageUri" validate:"uri"`
}
