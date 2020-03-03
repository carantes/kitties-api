package kittiesbundle

import (
	"net/http"

	"github.com/carantes/kitties-api/app/core"
)

// KittiesBundle handle kitties resources
type KittiesBundle struct {
	routes []core.Route
}

// NewKittiesBundle instance
func NewKittiesBundle() core.Bundle {
	repo := NewKittiesRepository()
	kc := NewKittiesController(repo)

	r := []core.Route{
		core.Route{
			Method:  http.MethodGet,
			Path:    "/kitties",
			Handler: kc.Index,
		},
		core.Route{
			Method:  http.MethodPost,
			Path:    "/kitties",
			Handler: kc.Create,
		},
		core.Route{
			Method:  http.MethodDelete,
			Path:    "/kitties/{id}",
			Handler: kc.Delete,
		},
	}

	return &KittiesBundle{
		routes: r,
	}
}

// GetRoutes implement interface core.Bundle
func (b *KittiesBundle) GetRoutes() []core.Route {
	return b.routes
}
