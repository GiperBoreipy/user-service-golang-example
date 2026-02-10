package entities

import (
	"time"

	"github.com/google/uuid"
)

type Meta struct {
	Id        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}
