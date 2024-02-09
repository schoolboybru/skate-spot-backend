package domain

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	  Id             uuid.UUID `json:"id" binding:"required" db:"id"`
    PostId         uuid.UUID `json:"post_id" binding:"required" db:"post_id"`
    UserId         uuid.UUID `json:"user_id" binding:"required" db:"user_id"`
    Value          string    `json:"value" binding:"required" db:"value"`
    CreatedAt      time.Time `json:"created_at" binding:"required" db:"created_at"`  
}
