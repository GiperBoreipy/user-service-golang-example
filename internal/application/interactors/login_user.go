package interactors

import (
	"user_service/internal/application/entities"
	app_interfaces "user_service/internal/application/interfaces"
	"user_service/internal/domain/entities"
	"user_service/internal/domain/interfaces"
)

type LoginUser struct {
	UserRepository     interfaces.Repository[*entities.User, entities.UserFilter]
	AccessTokenService app_interfaces.AccessTokenService
}

func (l *LoginUser) execute(email string, password string) (app_entities.AccessToken, error) {

}
