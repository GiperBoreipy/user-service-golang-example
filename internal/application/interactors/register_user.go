package interactors

import (
	"time"

	"user_service/internal/application"
	"user_service/internal/application/data_objects"
	"user_service/internal/domain/entities"
)

type RegisterUser struct {
	userRepository     application.Repository[entities.User, application.UserFilter]
	accessTokenService application.AccessTokenService
	passwordHasher     application.PasswordHasher
}

func NewRegisterUser(userRepository application.Repository[entities.User, application.UserFilter], accessTokenService application.AccessTokenService, passwordHasher application.PasswordHasher) RegisterUser {
	return RegisterUser{
		userRepository:     userRepository,
		accessTokenService: accessTokenService,
		passwordHasher:     passwordHasher,
	}
}

func (i RegisterUser) Execute(name string, email string, birthday time.Time, firstPassword string, secondPassword string) (data_objects.AccessToken, error) {
	if firstPassword != secondPassword {
		return data_objects.AccessToken{}, application.UserPasswordNotMatchError
	}

	password, err := i.passwordHasher.HashPassword(firstPassword)
	if err != nil {
		return data_objects.AccessToken{}, err
	}

	user, err := entities.NewUser(name, email, birthday, password)
	if err != nil {
		return data_objects.AccessToken{}, err
	}

	if err := i.userRepository.Add(user); err != nil {
		return data_objects.AccessToken{}, err
	}

	accessToken, err := i.accessTokenService.CreateAccessToken(user.Id)
	if err != nil {
		return data_objects.AccessToken{}, err
	}

	refreshToken, err := i.accessTokenService.CreateRefreshToken(user.Id)
	if err != nil {
		return data_objects.AccessToken{}, err
	}

	return data_objects.AccessToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
