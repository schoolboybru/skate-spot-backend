package rest

import (
	"github.com/schoolboybru/skate-spot/internal/services"
)

type Handlers interface {
	LocationHandler
	CommentHandler
	PostHandler
}

type handler struct {
	services services.Services
}

func NewHandler(s services.Services) Handlers {
	return &handler{services: s}
}
