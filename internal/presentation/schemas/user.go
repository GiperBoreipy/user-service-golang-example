package schemas

import "time"

type RegisterUserInSchema struct {
	Name           string    `json:"name"`
	Birthday       time.Time `json:"birthday"`
	Email          string    `json:"email"`
	FirstPassword  string    `json:"first_password"`
	SecondPassword string    `json:"second_password"`
}

type LoginUserInSchema struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AccessTokenOutSchema struct {
	AccessToken  string `json:"AccessToken"`
	RefreshToken string `json:"RefreshToken"`
}

type LoginUserOutSchema struct {
	AccessToken string `json:"AccessToken"`
}

type ErrorOutSchema struct {
	Message string `json:"message"`
}

type UserOutSchema struct {
	Id        string    `json:"Id"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
	Name      string    `json:"Name"`
	Email     string    `json:"Email"`
	Birthday  time.Time `json:"Birthday"`
}
