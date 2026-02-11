package interactors

import (
	app_entities "user_service/internal/application/entities"
	app_interfaces "user_service/internal/application/interfaces"
	"user_service/internal/domain/entities"
	"user_service/internal/domain/errors"
	"user_service/internal/domain/interfaces"
)

type LoginUser struct {
	UserRepository     interfaces.Repository[*entities.User, entities.UserFilter]
	AccessTokenService app_interfaces.AccessTokenService
	PasswordHasher     app_interfaces.PasswordHasher
}

func (l *LoginUser) Execute(email string, password string) (*app_entities.AccessToken, error) {
	users, err := l.UserRepository.Get(entities.UserFilter{Email: &email})
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.UserNotFoundError
	}

	user := users[0]

	status, err := l.PasswordHasher.VerifyPassword(password, user.HashedPassword)
	if err != nil {
		return nil, err
	}

	if !status {
		return nil, errors.UserPasswordNotMatchError
	}

	accessToken, err := l.AccessTokenService.CreateAccessToken(user.Id)
	if err != nil {
		return nil, err
	}

	return &app_entities.AccessToken{
		AccessToken: accessToken,
	}, nil
}
