package interactors

import (
	"time"
	"user_service/internal/domain/entities"
	"user_service/internal/domain/interfaces"
)

type RegisterUser struct {
	UserRepository interfaces.Repository[*entities.User, entities.UserFilter]
}

func (r *RegisterUser) execute(name string, email string, birthday time.Time) error {
	entity, error := entities.NewUser(name, email, birthday)
	if error != nil {
		// TODO: ошибку кастомную возвращать
		return error
	}

	go r.UserRepository.Add(entity)

	return nil
}
