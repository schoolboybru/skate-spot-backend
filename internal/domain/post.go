package domain

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	Id        uuid.UUID `json:"id" binding:"required" db:"id"`
	User      User      `json:"user" binding:"required"`
	Comments  []Comment `json:"comments"`
	Location  Location  `json:"location" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"required" db:"created_at"`
	Likes     int       `json:"likes"  db:"likes"`
}
