package v1

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"net/http"
	"spells.tvblackman1.ru/lib/pagination"
	"spells.tvblackman1.ru/pkg/domain/dto"
)

func (handler *V1Handler) spellsRoute(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		userId, _ := uuid.FromString("957a9c1a-a725-49fc-903d-f727e58146b5")
		spells, err := handler.usecases.Spell.GetSpellList(dto.UserId(userId),
			dto.SearchSpellDto{}, pagination.Pagination{
				Limit:      10,
				PageNumber: 1,
			})
		if err != nil {
			fmt.Println(err.Error())
		}
		w.Header().Set("Content-Type", "application/json")
		bytes, _ := json.Marshal(spells)
		w.Write(bytes)
	})
}
