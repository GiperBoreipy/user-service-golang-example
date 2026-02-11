package entities

import (
	"regexp"
	"time"

	"user_service/internal/domain/errors"

	"github.com/google/uuid"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

type UserFilter struct {
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
	Id             *uuid.UUID
	Name           *string
	Email          *string
	Birthday       *time.Time
	HashedPassword *string
}

type User struct {
	Meta
	Name           string
	Email          string
	Birthday       time.Time
	HashedPassword string
}

func NewUser(name string, email string, birthday time.Time) (*User, error) {
	uuid, error := uuid.NewV7()
	if error != nil {
		return nil, error
	}

	if !emailRegex.MatchString(email) {
		return nil, errors.UserEmailNotValidError
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
