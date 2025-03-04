package dto

type CreateUserRequest struct {
	Email     Email    `json:"email" binding:"required"`
	Password  Password `json:"password" binding:"required"`
	BirthDate string   `json:"birthDate" binding:"required"`
	Level     int      `json:"level" binding:"required"`
	Name      string   `json:"name" binding:"required"`
	Gender    string   `json:"gender" binding:"required"`
	Ring      Ring     `json:"ring"`
	Personal  Personal `json:"personal"`
}

type Ring struct {
	Size       int    `json:"size"`
	Color      string `json:"color"`
	Connection bool   `json:"connection"`
}

type Personal struct {
	Health   Health   `json:"health"`
	Physical Physical `json:"physical"`
	Habit    Habit    `json:"habit"`
}

type Health struct {
	Allergies     []string `json:"allergies"`
	Diseases      []string `json:"diseases"`
	Goals         []string `json:"goals"`
	BloodType     string   `json:"bloodType"`
	Issues        []string `json:"issues"`
	SpecificGoals []string `json:"specific_goals"`
}

type Physical struct {
	Height    float64 `json:"height"`
	Weight    float64 `json:"weight"`
	Abdominal float64 `json:"abdominal"`
	Waist     float64 `json:"waist"`
	Unit      string  `json:"unit"`
}

type Habit struct {
	Smoke   bool `json:"smoke"`
	Alcohol bool `json:"alcohol"`
}

type UpdateUserRequest struct {
	Email    Email    `json:"email"`
	Password Password `json:"password"`
	Data     UserData `json:"data"`
	Personal Personal `json:"personal"`
}

type Password struct {
	Value         string  `json:"value"`
	RequestForgot bool    `json:"request_forgot"`
	RequestChange bool    `json:"request_change"`
	History       *string `json:"history"`
}

type UserData struct {
	Name            string `json:"name"`
	BirthDate       string `json:"birth_date"`
	Gender          string `json:"gender"`
	AcknowledgedTOS bool   `json:"acknowledged_tos"`
	FirstLogin      bool   `json:"first_login"`
}

type Email struct {
	Value         string  `json:"value"`
	RequestForgot bool    `json:"request_forgot"`
	RequestChange bool    `json:"request_change"`
	History       *string `json:"history"`
}
