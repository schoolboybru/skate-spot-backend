package rest

import (
	"github.com/schoolboybru/location-service/internal/services"
)

type Handlers interface {
    LocationHandler
    CommentHandler
}

type handler struct {
	services services.Services
}

func NewHandler(s services.Services) Handlers {
	return &handler{services: s}
}

