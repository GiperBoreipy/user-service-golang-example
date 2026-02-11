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

func (r *RegisterUser) execute(name string, email string, birthday time.Time, firstPassword string, secondPassword string) (app_entities.AccessToken, error) {
	entity, error := entities.NewUser(name, email, birthday)
	if error != nil {
		// TODO: ошибку кастомную возвращать
		return app_entities.AccessToken{}, error
	}

	go r.UserRepository.Add(entity)

	accessToken, error := r.AccessTokenService.CreateAccessToken(entity.Id)
	if error != nil {
		return app_entities.AccessToken{}, error
	}

	refreshToken, error := r.AccessTokenService.CreateRefreshToken(entity.Id)
	if error != nil {
		return app_entities.AccessToken{}, error
	}

	return app_entities.AccessToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
