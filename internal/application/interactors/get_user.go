package interactors

import (
	"user_service/internal/application"
	"user_service/internal/application/data_objects"
	"user_service/internal/domain/entities"
)

type GetUser struct {
	userRepository     application.Repository[*entities.User, application.UserFilter]
	accessTokenService application.AccessTokenService
}

func NewGetUser(userRepository application.Repository[*entities.User, application.UserFilter], accessTokenService application.AccessTokenService) *GetUser {
	return &GetUser{
		userRepository:     userRepository,
		accessTokenService: accessTokenService,
	}
}

func (i *GetUser) Execute(authToken data_objects.UserAuthToken) (entities.User, error) {
	userId, err := i.accessTokenService.GetUserId(authToken)
	if err != nil {
		return entities.User{}, err
	}

	user, err := i.userRepository.GetOne(application.UserFilter{Id: userId})
	if err != nil {
		return entities.User{}, err
	}

	if user == nil {
		return entities.User{}, application.UserNotFoundError
	}

	return *user, nil
}
