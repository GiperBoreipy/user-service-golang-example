package interactors

import (
	"user_service/internal/domain/entities"
	"user_service/internal/domain/interfaces"
)

type GetAllUsers struct {
	UserRepository interfaces.Repository[*entities.User, entities.UserFilter]
}

func (g *GetAllUsers) Execute() ([]*entities.User, error) {
	users, err := g.UserRepository.Get(entities.UserFilter{})
	if err != nil {
		return nil, err
	}

	return users, nil
}
