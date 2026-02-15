package presentation

import (
	"encoding/json"
	"errors"
	"net/http"

	"user_service/internal/application"
	"user_service/internal/domain/entities"
)

func SendError(w http.ResponseWriter, err error) {
	statusCode := http.StatusInternalServerError
	message := "Internal server error"

	switch {
	case errors.Is(err, entities.ErrUserEmailNotValid):
		statusCode = http.StatusBadRequest
		message = "User email not valid"
	case errors.Is(err, application.ErrUserNotFound):
		statusCode = http.StatusNotFound
		message = "User not found"
	case errors.Is(err, application.ErrUserPasswordNotMatch):
		statusCode = http.StatusUnauthorized
		message = "User password not match"
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(map[string]string{"message": message})
}
