package application

import (
	data_objects "user_service/internal/application/data_objects"

	"github.com/google/uuid"
)

type AccessTokenService interface {
	CreateAccessToken(userId uuid.UUID) (data_objects.UserAuthToken, error)
	CreateRefreshToken(userId uuid.UUID) (data_objects.UserAuthToken, error)
	GetUserId(authToken data_objects.UserAuthToken) (uuid.UUID, error)
}

type PasswordHasher interface {
	HashPassword(password string) (string, error)
	VerifyPassword(password string, hashedPassword string) (bool, error)
}

type Repository[ET any, FT any] interface {
	Add(entity ET) error
	Delete(entity ET) error
	Update(entity ET) error
	Get(filter FT) ([]ET, error)
	GetOne(filter FT) (ET, error)
}
