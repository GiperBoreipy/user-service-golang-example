package schemas

import "time"

type RegisterUserInSchema struct {
	Name           string    `json:"name"`
	Birthday       time.Time `json:"birthday"`
	Email          string    `json:"email"`
	FirstPassword  string    `json:"first_password"`
	SecondPassword string    `json:"second_password"`
}
