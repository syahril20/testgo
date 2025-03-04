package dto

type UserResponse struct {
	Email     string `json:"email"`
	Level     int    `json:"level"`
	Name      string `json:"name"`
	BirthDate string `json:"birth_date"`
	Gender    string `json:"gender"`
}
