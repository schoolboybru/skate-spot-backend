package main

import (
	"github.com/schoolboybru/location-service/db"
	"github.com/schoolboybru/location-service/internal/adding"
	"github.com/schoolboybru/location-service/internal/http/rest"
)

func main() {

	repository, err := db.New()

	if err != nil {
		panic(err)
	}

	service := adding.New(repository)

	handler := rest.Handler(service)

	handler.Run(":8000")
}
