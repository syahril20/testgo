package models

import "time"

type User struct {
	ID        string       `bson:"_id,omitempty" json:"id,omitempty"`
	Email     UserEmail    `bson:"email" json:"email"`
	Level     int          `bson:"level" json:"level"`
	Password  UserPassword `bson:"password" json:"password"`
	Data      UserData     `bson:"data" json:"data"`
	Ring      UserRing     `bson:"ring" json:"ring"`
	Personal  UserPersonal `bson:"personal" json:"personal"`
	DeletedAt *time.Time   `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
	CreatedAt time.Time    `json:"created_at" bson:"created_at"`
	CreatedBy string       `json:"created_by" bson:"created_by"`
	UpdatedAt time.Time    `json:"updated_at" bson:"updated_at"`
	UpdatedBy string       `json:"updated_by" bson:"updated_by"`
}

type UserEmail struct {
	Value         string     `bson:"value" json:"value"`
	RequestValue  *string    `bson:"request_value,omitempty" json:"request_value,omitempty"`
	RequestForgot bool       `bson:"request_forgot" json:"request_forgot"`
	RequestChange bool       `bson:"request_change" json:"request_change"`
	RequestExpire *time.Time `bson:"request_expire,omitempty" json:"request_expire,omitempty"`
	Token         *string    `bson:"token,omitempty" json:"token,omitempty"`
	History       []string   `bson:"history" json:"history"`
}

type UserPassword struct {
	Value         string     `bson:"value" json:"value"`
	RequestValue  *string    `bson:"request_value,omitempty" json:"request_value,omitempty"`
	RequestForgot bool       `bson:"request_forgot" json:"request_forgot"`
	RequestChange bool       `bson:"request_change" json:"request_change"`
	RequestExpire *time.Time `bson:"request_expire,omitempty" json:"request_expire,omitempty"`
	Token         *string    `bson:"token,omitempty" json:"token,omitempty"`
	History       []string   `bson:"history" json:"history"`
}

type UserData struct {
	Name       string    `bson:"name" json:"name"`
	BirthDate  time.Time `bson:"birth_date" json:"birth_date"`
	Gender     string    `bson:"gender" json:"gender"`
	AckTOS     bool      `bson:"acknowledged_tos" json:"acknowledged_tos"`
	FirstLogin bool      `bson:"first_login" json:"first_login"`
}

type UserRing struct {
	MAC          *string    `bson:"mac,omitempty" json:"mac,omitempty"`
	PurchaseDate *time.Time `bson:"purchase_date,omitempty" json:"purchase_date,omitempty"`
	Size         int        `bson:"size" json:"size"`
	Color        string     `bson:"color" json:"color"`
	Connection   bool       `bson:"connection" json:"connection"`
}

type UserPersonal struct {
	Health    UserHealth   `bson:"health" json:"health"`
	Physical  UserPhysical `bson:"physical" json:"physical"`
	Habit     UserHabit    `bson:"habit" json:"habit"`
	Goals     UserGoals    `bson:"goals" json:"goals"`
	UpdatedAt time.Time    `bson:"updated_at" json:"updated_at"`
}

type UserHealth struct {
	Allergies     []string `bson:"allergies" json:"allergies"`
	Diseases      []string `bson:"diseases" json:"diseases"`
	Goals         []string `bson:"goals" json:"goals"`
	BloodType     string   `bson:"bloodType" json:"bloodType"`
	Issues        []string `bson:"issues" json:"issues"`
	SpecificGoals []string `bson:"specific_goals" json:"specific_goals"`
}

type UserPhysical struct {
	Height    float64 `bson:"height" json:"height"`
	Weight    float64 `bson:"weight" json:"weight"`
	Abdominal float64 `bson:"abdominal" json:"abdominal"`
	Waist     float64 `bson:"waist" json:"waist"`
	Unit      string  `bson:"unit" json:"unit"`
}

type UserHabit struct {
	Smoke   bool `bson:"smoke" json:"smoke"`
	Alcohol bool `bson:"alcohol" json:"alcohol"`
}

type UserGoals struct {
	Health  []string `bson:"health" json:"health"`
	Sports  []string `bson:"sports" json:"sports"`
	Medical []string `bson:"medical" json:"medical"`
}

type UserResponse struct {
	ID       string       `bson:"_id" json:"id,omitempty"`
	Email    UserEmail    `bson:"email" json:"email"`
	Level    int          `bson:"level" json:"level"`
	Data     UserData     `bson:"data" json:"data"`
	Ring     UserRing     `bson:"ring" json:"ring"`
	Personal UserPersonal `bson:"personal" json:"personal"`
}
