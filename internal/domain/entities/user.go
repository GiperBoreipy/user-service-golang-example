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

func NewUser(name string, email string, birthday time.Time) (*User, error) {
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
		Name:     name,
		Email:    email,
		Birthday: birthday,
	}, nil
}
