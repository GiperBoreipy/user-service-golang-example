package handlers

import (
	"net/http"
	"user_service/internal/application/interactors"
)

type handlerUser struct {
	registerUser *interactors.RegisterUser
	loginUser    *interactors.LoginUser
	getUser      *interactors.GetUser
	getAllUsers  *interactors.GetAllUsers
}

func NewHandlerUser(registerUser *interactors.RegisterUser, loginUser *interactors.LoginUser, getUser *interactors.GetUser, getAllUsers *interactors.GetAllUsers) *handlerUser {
	return &handlerUser{
		registerUser: registerUser,
		loginUser:    loginUser,
		getUser:      getUser,
		getAllUsers:  getAllUsers,
	}
}

func (h *handlerUser) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *handlerUser) LoginUserHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *handlerUser) MeHandler(w http.ResponseWriter, r *http.Request) {
	token, ok L
}
