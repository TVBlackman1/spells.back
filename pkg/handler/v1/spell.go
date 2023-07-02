package v1

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"net/http"
	"spells.tvblackman1.ru/lib/pagination"
	"spells.tvblackman1.ru/pkg/domain/dto"
	//"spells.tvblackman1.ru/pkg/domain/dto"
)

// /v1/spells
func (handler *V1Handler) spellsRoute(r chi.Router) {
	handler.getSpells(r)
}

// ShowAccount godoc
// @Summary      Get spell list
// @Description  get spells by filters
// @Tags         spells
// @Accept       json
// @Produce      json
// @Success      200  {object}  []prettySpell
// @Router       /v1/spells/ [get]
func (handler *V1Handler) getSpells(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		userId, _ := uuid.FromString("957a9c1a-a725-49fc-903d-f727e58146b5")
		// TODO should to use userId, _ := uuid.FromString("957a9c1a-a725-49fc-903d-f727e58146b5"); check why it works now with 0000000000000000
		spells, err := handler.usecases.Spell.GetSpellList(dto.UserId(userId),
			dto.SearchSpellDto{}, pagination.Pagination{
				Limit:      10,
				PageNumber: 1,
			})
		if err != nil {
			fmt.Println(err.Error())
		}
		prettySpells := make([]prettySpell, len(spells), len(spells))
		for index, spell := range spells {
			prettySpells[index] = spellDtoToPretty(spell)
		}
		w.Header().Set("Content-Type", "application/json")
		bytes, _ := json.Marshal(prettySpells)
		_, err = w.Write(bytes)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	})
}

func spellDtoToPretty(dto dto.SpellDto) prettySpell {
	return prettySpell{
		Id:                   uuid.UUID(dto.Id).String(),
		Name:                 dto.Name,
		Level:                dto.Level,
		Classes:              nil,
		Description:          dto.Description,
		CastingTime:          dto.CastingTime,
		Duration:             dto.Duration,
		IsVerbal:             dto.IsVerbal,
		IsSomatic:            dto.IsSomatic,
		HasMaterialComponent: dto.HasMaterialComponent,
		MaterialComponent:    dto.MaterialComponent,
		MagicalSchool:        dto.MagicalSchool,
		Distance:             dto.Distance,
		IsRitual:             dto.IsRitual,
		SourceId:             uuid.UUID(dto.SourceId).String(),
		SourceName:           dto.SourceName,
	}
}

func spellMarkedDtoToPretty(dto dto.SpellMarkedDto) prettySpellMarked {
	spellDto := dto.SpellDto
	prettyPart := spellDtoToPretty(spellDto)
	return prettySpellMarked{
		prettySpell: prettyPart,
		InSet:       dto.InSet,
	}
}

type prettySpell struct {
	Id                   string   `json:"id"`
	Name                 string   `json:"name"`
	Level                int      `json:"level"`
	Classes              []string `json:"classes"`
	Description          string   `json:"description"`
	CastingTime          string   `json:"casting_time"`
	Duration             string   `json:"duration"`
	IsVerbal             bool     `json:"is_verbal"`
	IsSomatic            bool     `json:"is_somatic"`
	HasMaterialComponent bool     `json:"has_material_component"`
	MaterialComponent    string   `json:"material_component"`
	MagicalSchool        string   `json:"magical_school"`
	Distance             string   `json:"distance"`
	IsRitual             bool     `json:"is_ritual"`
	SourceId             string   `json:"source_id"`
	SourceName           string   `json:"source_name"`
}

type prettySpellMarked struct {
	prettySpell
	InSet bool `json:"in_set"`
}
