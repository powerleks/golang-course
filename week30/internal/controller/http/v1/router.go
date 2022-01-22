package v1

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"restapi/internal/usecase"
)

func NewRouter(r *chi.Mux, t usecase.User) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	{
		newUserRoutes(r, t)
	}
}
