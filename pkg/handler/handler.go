package handler

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"spells.tvblackman1.ru/pkg/domain/usecases"
)

type Handler struct {
	router   chi.Router
	usecases *usecases.UseCases
}

func NewHandler(usecases *usecases.UseCases) *Handler {
	handler := new(Handler)
	handler.usecases = usecases
	handler.addRoutes()
	return handler
}

func (handler *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.router.ServeHTTP(w, r)
}
