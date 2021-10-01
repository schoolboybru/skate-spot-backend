package domain

import (
	"time"

	"github.com/google/uuid"
)

type Location struct {
	Id        uuid.UUID `json:"id" binding:"required" db:"id"`
	Name      string    `json:"name" binding:"required" db:"name"`
	Longitude float64   `json:"longitude" binding:"required" db:"longitude"`
	Latitude  float64   `json:"latitude" binding:"required" db:"latitude"`
	Date      time.Time `json:"date" binding:"required" db:"date"`
	City      string    `json:"city" binding:"required" db:"city"`
	Country   string    `json:"country" binding:"required" db:"country"`
}
