package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	domain_errors "user_service/internal/domain/errors"
)

func sendError(w http.ResponseWriter, err error) {
	statusCode := http.StatusInternalServerError
	message := "Internal server error"

	switch {
	case errors.Is(err, domain_errors.UserEmailNotValidError):
		statusCode = http.StatusBadRequest
		message = "User email not valid"
	case errors.Is(err, domain_errors.UserNotFoundError):
		statusCode = http.StatusNotFound
		message = "User not found"
	case errors.Is(err, domain_errors.UserPasswordNotMatchError):
		statusCode = http.StatusUnauthorized
		message = "User password not match"
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(map[string]string{"message": message})
}
