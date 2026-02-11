package interactors

import (
	"user_service/internal/domain/entities"
	"user_service/internal/domain/errors"
	"user_service/internal/domain/interfaces"

	"github.com/google/uuid"
)

type GetUser struct {
	UserRepository interfaces.Repository[*entities.User, entities.UserFilter]
}

func (g *GetUser) Execute(userId uuid.UUID) (*entities.User, error) {
	users, err := g.UserRepository.Get(entities.UserFilter{Id: &userId})
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.UserNotFoundError
	}

	return users[0], nil
}
