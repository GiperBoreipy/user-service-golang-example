package interactors

import (
	data_objects "user_service/internal/application/data_objects"

	"github.com/google/uuid"
)

type AccessTokenService interface {
	CreateAccessToken(UserId uuid.UUID) (data_objects.UserAuthToken, error)
	CreateRefreshToken(UserId uuid.UUID) (data_objects.UserAuthToken, error)
	GetUserId(authToken data_objects.UserAuthToken) (uuid.UUID, error)
}
