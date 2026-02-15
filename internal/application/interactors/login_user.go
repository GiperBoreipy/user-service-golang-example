package interactors

import (
	"user_service/internal/application"
	"user_service/internal/application/data_objects"
	"user_service/internal/domain/entities"
)

type LoginUser struct {
	userRepository     application.Repository[entities.User, application.UserFilter]
	accessTokenService application.AccessTokenService
	passwordHasher     application.PasswordHasher
}

func NewLoginUser(userRepository application.Repository[entities.User, application.UserFilter], accessTokenService application.AccessTokenService, passwordHasher application.PasswordHasher) LoginUser {
	return LoginUser{
		userRepository:     userRepository,
		accessTokenService: accessTokenService,
		passwordHasher:     passwordHasher,
	}
}

func (i LoginUser) Execute(email string, password string) (data_objects.AccessToken, error) {
	user, err := i.userRepository.GetOne(application.UserFilter{Email: &email})
	if err != nil {
		return data_objects.AccessToken{}, err
	}

	status, err := i.passwordHasher.VerifyPassword(password, user.HashedPassword)
	if err != nil {
		return data_objects.AccessToken{}, err
	} else if !status {
		return data_objects.AccessToken{}, application.ErrUserPasswordNotMatch
	}

	accessToken, err := i.accessTokenService.CreateAccessToken(user.Id)
	if err != nil {
		return data_objects.AccessToken{}, err
	}

	return data_objects.AccessToken{
		AccessToken: accessToken,
	}, nil
}
