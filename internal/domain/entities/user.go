package entities

import (
	"time"

	"github.com/google/uuid"
)

type UserFilter struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	Id        uuid.UUID
	Name      string
	Email     string
	Birthday  time.Time
}

type User struct {
	Meta
	Name     string
	Email    string
	Birthday time.Time
}

func (u *User) New(Name string, Email string, Birthday time.Time) (*User, error) {
	uuid, error := uuid.NewV7()
	if error != nil {
		return nil, error
	}

	return &User{
		Meta: Meta{
			Id:        uuid,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:     Name,
		Email:    Email,
		Birthday: Birthday,
	}, nil
}
