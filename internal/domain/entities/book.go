package entities

import (
	"time"

	"github.com/google/uuid"
)

type BookFilter struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Id          uuid.UUID
	Title       string
	Author      string
	PublishedAt time.Time
}

type Book struct {
	Meta
	Title       string
	Author      string
	PublishedAt time.Time
}
