package interactors

import (
	"time"
	app_entities "user_service/internal/application/entities"
	app_interfaces "user_service/internal/application/interfaces"
	"user_service/internal/domain/entities"
	"user_service/internal/domain/interfaces"
)

type RegisterUser struct {
	UserRepository     interfaces.Repository[*entities.User, entities.UserFilter]
	AccessTokenService app_interfaces.AccessTokenService
}

func (r *RegisterUser) Execute(name string, email string, birthday time.Time, firstPassword string, secondPassword string) (*app_entities.AccessToken, error) {
	user, err := entities.NewUser(name, email, birthday)
	if err != nil {
		return nil, err
	}

	go r.UserRepository.Add(user)

	accessToken, err := r.AccessTokenService.CreateAccessToken(user.Id)
	if err != nil {
		return nil, err
	}

	refreshToken, err := r.AccessTokenService.CreateRefreshToken(user.Id)
	if err != nil {
		return nil, err
	}

	return &app_entities.AccessToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
