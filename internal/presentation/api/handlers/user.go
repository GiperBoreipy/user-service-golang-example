package handlers

import (
	"encoding/json"
	"net/http"
	"user_service/internal/application/data_objects"
	"user_service/internal/application/interactors"
	"user_service/internal/presentation/api/schemas"
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
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()

	var requestData schemas.RegisterUserInSchema
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	accessTokens, err := h.registerUser.Execute(requestData.Name, requestData.Email, requestData.Birthday, requestData.FirstPassword, requestData.SecondPassword)
	if err != nil {
		sendError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(accessTokens)
}

func (h *handlerUser) LoginUserHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *handlerUser) MeHandler(w http.ResponseWriter, r *http.Request) {
	authToken, ok := r.Context().Value(data_objects.UserAuthTokenContextKey).(data_objects.UserAuthToken)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := h.getUser.Execute(authToken)
	if err != nil {
		sendError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(user)
}
