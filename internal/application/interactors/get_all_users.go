package interactors

import (
	"user_service/internal/application"
	"user_service/internal/domain/entities"
)

type GetAllUsers struct {
	userRepository application.Repository[*entities.User, application.UserFilter]
}

func NewGetAllUsers(userRepository application.Repository[*entities.User, application.UserFilter]) *GetAllUsers {
	return &GetAllUsers{userRepository: userRepository}
}

func (i *GetAllUsers) Execute() ([]*entities.User, error) {
	users, err := i.userRepository.Get(application.UserFilter{})
	if err != nil {
		return nil, err
	}

	return users, nil
}
