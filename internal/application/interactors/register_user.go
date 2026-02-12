package interactors

import (
	"time"
	"user_service/internal/application/data_objects"
	app_interfaces "user_service/internal/application/interfaces"
	"user_service/internal/domain/entities"
	"user_service/internal/domain/interfaces"
)

type RegisterUser struct {
	UserRepository     interfaces.Repository[*entities.User, entities.UserFilter]
	AccessTokenService app_interfaces.AccessTokenService
}

func (r *RegisterUser) Execute(name string, email string, birthday time.Time, firstPassword string, secondPassword string) (*data_objects.AccessToken, error) {
	user, err := entities.NewUser(name, email, birthday)
	if err != nil {
		return nil, err
	}

	if err := r.UserRepository.Add(user); err != nil {
		return nil, err
	}

	accessToken, err := r.AccessTokenService.CreateAccessToken(user.Id)
	if err != nil {
		return nil, err
	}

	refreshToken, err := r.AccessTokenService.CreateRefreshToken(user.Id)
	if err != nil {
		return nil, err
	}

	return &data_objects.AccessToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
