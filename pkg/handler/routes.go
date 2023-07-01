package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"spells.tvblackman1.ru/pkg/handler/v1"
	"time"
)

func (handler *Handler) addRoutes() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	v1Handler := v1.NewHandler(handler.usecases, handler.router)
	r.Route("/v1", v1Handler.Route)
	handler.router = r
}
