package core

import (
	"net/http"

	_ "user_service/docs"

	httpSwagger "github.com/swaggo/http-swagger"

	"user_service/internal/application/interactors"
	"user_service/internal/infra/impl"
	"user_service/internal/presentation/handlers"
)

func GetServer() *http.Server {
	config := LoadConfig()

	userRepository := impl.NewMemoryUserRepository()
	accessTokenService := impl.NewJwtAccessTokenServiceImpl(
		config.AccessTokenSecret,
		config.RefreshTokenSecret,
		config.AccessTokenTTL,
		config.RefreshTokenTTL,
	)
	passwordHasher := impl.NewBcryptPasswordHasher(0)

	registerUser := interactors.NewRegisterUser(&userRepository, accessTokenService, passwordHasher)
	loginUser := interactors.NewLoginUser(&userRepository, accessTokenService, passwordHasher)
	getUser := interactors.NewGetUser(&userRepository, accessTokenService)
	getAllUsers := interactors.NewGetAllUsers(&userRepository)

	mux := http.NewServeMux()
	handlers.InitUserHandlers(mux, registerUser, loginUser, getUser, getAllUsers)
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	return &http.Server{Handler: mux, Addr: ":8080"}
}
