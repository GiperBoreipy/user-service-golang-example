package entities

import (
	"errors"
	"regexp"
	"time"

	"github.com/google/uuid"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

var UserEmailNotValidError = errors.New("user email not valid")

type User struct {
	Meta
	Name           string
	Email          string
	Birthday       time.Time
	HashedPassword string
}

func NewUser(name string, email string, birthday time.Time, hashedPassword string) (*User, error) {
	uuid, error := uuid.NewV7()
	if error != nil {
		return nil, error
	}

	if !emailRegex.MatchString(email) {
		return nil, UserEmailNotValidError
	}

	return &User{
		Meta: Meta{
			Id:        uuid,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:           name,
		Email:          email,
		Birthday:       birthday,
		HashedPassword: hashedPassword,
	}, nil
}
