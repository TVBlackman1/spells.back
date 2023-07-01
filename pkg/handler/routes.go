package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/swaggo/http-swagger/v2"
	_ "spells.tvblackman1.ru/docs"
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
	r.Route("/api/v1", v1Handler.Route)
	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL(
		"http://localhost:8080/docs/doc.json")))
	handler.router = r
}
