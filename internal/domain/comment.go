package domain

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	Id        uuid.UUID `json:"id" binding:"required" db:"id"`
	UserId    uuid.UUID `json:"user_id" binding:"required" db:"fk_user_id"`
	Value     string    `json:"value" binding:"required" db:"value"`
	CreatedAt time.Time `json:"created_at" binding:"required" db:"created_at"`
}
