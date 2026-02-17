package handlers

import (
	"encoding/json"
	"net/http"
	"user_service/internal/application/data_objects"
	"user_service/internal/application/interactors"
	"user_service/internal/presentation"
	"user_service/internal/presentation/schemas"
)

type handlerUser struct {
	registerUser interactors.RegisterUser
	loginUser    interactors.LoginUser
	getUser      interactors.GetUser
	getAllUsers  interactors.GetAllUsers
}

func NewHandlerUser(registerUser interactors.RegisterUser, loginUser interactors.LoginUser, getUser interactors.GetUser, getAllUsers interactors.GetAllUsers) handlerUser {
	return handlerUser{
		registerUser: registerUser,
		loginUser:    loginUser,
		getUser:      getUser,
		getAllUsers:  getAllUsers,
	}
}

// RegisterUserHandler godoc
// @Summary Register user
// @Tags user
// @Accept json
// @Produce json
// @Param request body schemas.RegisterUserInSchema true "register request"
// @Success 200 {object} schemas.AccessTokenOutSchema
// @Failure 400 {object} schemas.ErrorOutSchema
// @Failure 405 {object} schemas.ErrorOutSchema
// @Router /user/register [post]
func (h handlerUser) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
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
		presentation.SendError(w, err)
		return
	}

	response := schemas.AccessTokenOutSchema{
		AccessToken:  string(accessTokens.AccessToken),
		RefreshToken: string(accessTokens.RefreshToken),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}

// LoginUserHandler godoc
// @Summary Login user
// @Tags user
// @Accept json
// @Produce json
// @Param request body schemas.LoginUserInSchema true "login request"
// @Success 200 {object} schemas.LoginUserOutSchema
// @Failure 400 {object} schemas.ErrorOutSchema
// @Failure 401 {object} schemas.ErrorOutSchema
// @Failure 405 {object} schemas.ErrorOutSchema
// @Router /user/login [post]
func (h handlerUser) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()

	var requestData schemas.LoginUserInSchema
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	accessToken, err := h.loginUser.Execute(requestData.Email, requestData.Password)
	if err != nil {
		presentation.SendError(w, err)
		return
	}

	response := schemas.LoginUserOutSchema{AccessToken: string(accessToken.AccessToken)}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}

// MeHandler godoc
// @Summary Current user
// @Tags user
// @Produce json
// @Success 200 {object} schemas.UserOutSchema
// @Failure 401 {object} schemas.ErrorOutSchema
// @Failure 405 {object} schemas.ErrorOutSchema
// @Router /user/me [get]
func (h handlerUser) MeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	authToken, ok := r.Context().Value(data_objects.UserAuthTokenContextKey).(data_objects.UserAuthToken)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := h.getUser.Execute(authToken)
	if err != nil {
		presentation.SendError(w, err)
		return
	}

	response := schemas.UserOutSchema{
		Id:        user.Id.String(),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		Email:     user.Email,
		Birthday:  user.Birthday,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}
