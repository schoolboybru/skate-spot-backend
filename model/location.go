package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/schoolboybru/location-service/db"
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

func GetLocationById(id string) (*Location, error) {
	var location Location
	database, err := db.New()

	if err != nil {
		return nil, err
	}

	defer database.Close()

	err = database.Get(&location, "SELECT * FROM Location WHERE id=$1", id)

	if err != nil {
		println(err.Error())
	}

	return &location, nil

}

func AddLocation(newLocation Location) {

	database, err := db.New()

	if err != nil {
		println(err)
	}

	defer database.Close()

	_, err = database.NamedExec(`INSERT INTO LOCATION (id, name, longitude, latitude, date, city, country) VALUES (:id, :name, :longitude, :latitude, :date, :city, :country)`, &newLocation)

	if err != nil {
		println(err.Error())
	}
}
