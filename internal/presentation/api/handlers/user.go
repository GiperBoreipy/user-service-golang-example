package handlers

import (
	"net/http"
	"user_service/internal/application/interactors"
)

type userHandler struct {
	registerUser *interactors.RegisterUser
	loginUser    *interactors.LoginUser
	getUser      *interactors.GetUser
	getAllUsers  *interactors.GetAllUsers
}

func NewHandler(registerUser *interactors.RegisterUser, loginUser *interactors.LoginUser, getUser *interactors.GetUser, getAllUsers *interactors.GetAllUsers) *userHandler {
	return &userHandler{
		registerUser: registerUser,
		loginUser:    loginUser,
		getUser:      getUser,
		getAllUsers:  getAllUsers,
	}
}

func (u *userHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {

}
