package core

import (
	"net/http"
	"user_service/internal/application/interactors"
	"user_service/internal/infra/impl"
	"user_service/internal/presentation/handlers"
)

func GetServer() *http.Server {
	userRepository := impl.NewMemoryUserRepository()

	registerUser := interactors.NewRegisterUser()
	loginUser := interactors.NewLoginUser()
	getUser := interactors.NewGetUser()
	getAllUsers := interactors.NewGetAllUsers()

	mux := http.NewServeMux()
	handlers.InitUserHandlers(mux, registerUser, loginUser, getUser, getAllUsers)
}
