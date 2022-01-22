package app

import (
	"net/http"

	"github.com/go-chi/chi"

	v1 "restapi/internal/controller/http/v1"
	"restapi/internal/usecase"
	"restapi/internal/usecase/repo"
)

func Run() {
	userUseCase := usecase.New(
		repo.New(),
	)

	handler := chi.NewRouter()
	v1.NewRouter(handler, userUseCase)
	http.ListenAndServe(":3333", handler)
}
