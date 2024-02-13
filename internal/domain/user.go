package domain

import "github.com/google/uuid"

type User struct {
	Id        uuid.UUID `json:"id" binding:"required" db:"id"`
	FirstName string    `json:"first_name" binding:"required" db:"first_name"`
	LastName  string    `json:"last_name" binding:"required" db:"last_name"`
	Password  string    `json:"password" binding:"required" db:"password"`
	Email     string    `json:"email" binding:"required" db:"email"`
}
