package interactors

import (
	data_objects "user_service/internal/application/data_objects"
	app_interfaces "user_service/internal/application/interfaces"
	"user_service/internal/domain/entities"
	"user_service/internal/domain/errors"
	"user_service/internal/domain/interfaces"
)

type GetUser struct {
	UserRepository     interfaces.Repository[*entities.User, entities.UserFilter]
	AccessTokenService app_interfaces.AccessTokenService
}

func (g *GetUser) Execute(authToken data_objects.UserAuthToken) (*entities.User, error) {
	userId, err := g.AccessTokenService.GetUserId(authToken)
	if err != nil {
		return nil, err
	}

	users, err := g.UserRepository.Get(entities.UserFilter{Id: &userId})
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.UserNotFoundError
	}

	return users[0], nil
}
