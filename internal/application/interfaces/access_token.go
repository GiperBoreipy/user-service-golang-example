package interactors

import "github.com/google/uuid"

type AccessTokenService interface {
	CreateAccessToken(UserId uuid.UUID) (string, error)
	CreateRefreshToken(UserId uuid.UUID) (string, error)
}
