package application

import (
	"time"

	"github.com/google/uuid"
)

type UserFilter struct {
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Id             uuid.UUID
	Name           string
	Email          string
	Birthday       time.Time
	HashedPassword string
}
