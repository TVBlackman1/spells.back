package v1

import (
	"github.com/go-chi/chi/v5"
	"spells.tvblackman1.ru/pkg/domain/usecases"
)

type V1Handler struct {
	router   chi.Router
	usecases *usecases.UseCases
}

func NewHandler(usecases *usecases.UseCases, router chi.Router) *V1Handler {
	handler := new(V1Handler)
	handler.usecases = usecases
	handler.router = router
	return handler
}

func (handler *V1Handler) Route(r chi.Router) {
	r.Route("/spells", handler.spellsRoute)
}
