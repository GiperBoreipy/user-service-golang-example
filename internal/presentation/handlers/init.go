package handlers

import (
	"net/http"
	"user_service/internal/application/interactors"
)

func InitUserHandlers(mux *http.ServeMux, registerUser interactors.RegisterUser, loginUser interactors.LoginUser, getUser interactors.GetUser, getAllUsers interactors.GetAllUsers) {
	handlerUser := NewHandlerUser(registerUser, loginUser, getUser, getAllUsers)

	mux.HandleFunc("/user/register", handlerUser.RegisterUserHandler)
	mux.HandleFunc("/user/login", handlerUser.LoginUserHandler)
	mux.HandleFunc("/user/me", handlerUser.MeHandler)
}
