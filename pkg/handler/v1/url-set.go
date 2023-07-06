package v1

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"net/http"
	"spells.tvblackman1.ru/lib/pagination"
	"spells.tvblackman1.ru/pkg/domain/dto"
	"strconv"
)

// /v1/spells
func (handler *V1Handler) urlSetsRoute(r chi.Router) {
	handler.createUrlSet(r)
	handler.getUrlSet(r)
	handler.renameUrlSet(r)
	handler.getAllSpells(r)
	handler.addSpellToUrlSet(r)
	handler.getSpellsOfUrlSet(r)
	handler.removeSpellFromUrlSet(r)
}

// ShowAccount godoc
// @Summary      Create url set
// @Description  Create empty url set without name.
// @Tags         url-sets
// @Accept       json
// @Produce      json
// @Success      200 {string} string "link to url set"
// @Router       /v1/url-sets/ [Post]
func (handler *V1Handler) createUrlSet(r chi.Router) {
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		link, err := handler.usecases.UrlSet.CreateUrlSet()
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte(link))
		}
	})
}

// ShowAccount godoc
// @Summary      Get url set
// @Description  Get basic info about url set. Just name.
// @Tags         url-sets
// @Accept       json
// @Produce      json
// @Param        unique   path  string  true  "url set unique link part"
// @Success      200 {object} prettySetDto
// @Router       /v1/url-sets/{unique} [get]
func (handler *V1Handler) getUrlSet(r chi.Router) {
	r.Get("/{unique}", func(w http.ResponseWriter, r *http.Request) {
		uniqueLinkPart := chi.URLParam(r, "unique")
		urlSet, err := handler.usecases.UrlSet.GetUrlSet(uniqueLinkPart)
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		prettyUrlSet := urlSetDtoToPretty(urlSet)
		w.Header().Set("Content-Type", "application/json")
		bytes, _ := json.Marshal(prettyUrlSet)
		_, err = w.Write(bytes)

	})
}

// ShowAccount godoc
// @Summary      Rename url set
// @Tags         url-sets
// @Accept       json
// @Produce      json
// @Param        unique   path  string  true  "url set unique link part"
// @Param		 body	body		renameUrlSetDto	true	"Updating data"
// @Success      200
// @Router       /v1/url-sets/{unique} [put]
func (handler *V1Handler) renameUrlSet(r chi.Router) {
	r.Put("/{unique}", func(w http.ResponseWriter, r *http.Request) {
		var body renameUrlSetDto
		json.NewDecoder(r.Body).Decode(&body)
		uniqueLinkPart := chi.URLParam(r, "unique")
		handler.usecases.UrlSet.RenameUrlSet(uniqueLinkPart, body.Name)
	})
}

// ShowAccount godoc
// @Summary      Get list of all spells
// @Description  get all spells. Overview for adding/removing spells to/from set
// @Tags         url-sets
// @Param        unique   path  string  true  "url set unique link part"
// @Accept       json
// @Produce      json
// @Success      200  {object}  []prettySpell
// @Router       /v1/url-sets/{unique}/all-spells [get]
func (handler *V1Handler) getAllSpells(r chi.Router) {
	r.Get("/{unique}/all-spells", func(w http.ResponseWriter, r *http.Request) {
		pageQueryParam := r.URL.Query().Get("page")
		pageNumber, _ := strconv.Atoi(pageQueryParam)
		uniqueLinkPart := chi.URLParam(r, "unique")
		markedSpells, _, err := handler.usecases.UrlSet.GetAllSpells(uniqueLinkPart, dto.SearchSpellDto{}, pagination.Pagination{
			Limit:      15,
			PageNumber: pageNumber,
		})
		if err != nil {
			fmt.Println(err.Error())
		}
		prettySpells := make([]prettySpellMarked, len(markedSpells), len(markedSpells))
		for index, spell := range markedSpells {
			prettySpells[index] = spellMarkedDtoToPretty(spell)
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

// ShowAccount godoc
// @Summary      Add spell to url set
// @Description  Can add spell to url set
// @Tags         url-sets
// @Param        unique   path  string  true  "url set unique link part"
// @Param        spellId  path  string  true  "id of spell"
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /v1/url-sets/{unique}/add/{spellId} [post]
func (handler *V1Handler) addSpellToUrlSet(r chi.Router) {
	r.Post("/{unique}/add/{spellId}", func(w http.ResponseWriter, r *http.Request) {
		uniqueLinkPart := chi.URLParam(r, "unique")
		spellIdURLParam := chi.URLParam(r, "spellId")
		_spellId, err := uuid.FromString(spellIdURLParam)
		if err != nil {
			fmt.Println(err)
			return
		}
		spellId := dto.SpellId(_spellId)
		handler.usecases.UrlSet.AddSpell(uniqueLinkPart, spellId)
	})
}

// ShowAccount godoc
// @Summary      Remove spell from url set
// @Description  Can remove spell from url set
// @Tags         url-sets
// @Param        unique   path  string  true  "url set unique link part"
// @Param        spellId  path  string  true  "id of spell"
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /v1/url-sets/{unique}/remove/{spellId} [delete]
func (handler *V1Handler) removeSpellFromUrlSet(r chi.Router) {
	r.Delete("/{unique}/remove/{spellId}", func(w http.ResponseWriter, r *http.Request) {
		uniqueLinkPart := chi.URLParam(r, "unique")
		spellIdURLParam := chi.URLParam(r, "spellId")
		_spellId, err := uuid.FromString(spellIdURLParam)
		if err != nil {
			fmt.Println(err)
			return
		}
		spellId := dto.SpellId(_spellId)
		handler.usecases.UrlSet.RemoveSpell(uniqueLinkPart, spellId)
	})
}

// ShowAccount godoc
// @Summary      Get list of url set spells
// @Description  Get only spells in url set
// @Tags         url-sets
// @Param        unique   path  string  true  "url set unique link part"
// @Accept       json
// @Produce      json
// @Success      200  {object}  []prettySpell
// @Router       /v1/url-sets/{unique}/spells [get]
func (handler *V1Handler) getSpellsOfUrlSet(r chi.Router) {
	r.Get("/{unique}/spells", func(w http.ResponseWriter, r *http.Request) {
		pageQueryParam := r.URL.Query().Get("page")
		pageNumber, _ := strconv.Atoi(pageQueryParam)
		uniqueLinkPart := chi.URLParam(r, "unique")
		spells, _, err := handler.usecases.UrlSet.GetSpells(uniqueLinkPart, dto.SearchSpellDto{}, pagination.Pagination{
			Limit:      pageNumber,
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

func urlSetDtoToPretty(dto dto.UrlSetDto) prettySetDto {
	return prettySetDto{
		Id:   uuid.UUID(dto.Id).String(),
		Name: dto.Name,
		Uri:  dto.Uri,
	}
}

type prettySetDto struct {
	Id   string `json:"id,omitempty"`
	Uri  string `json:"uri,omitempty"`
	Name string `json:"name,omitempty"`
}

type renameUrlSetDto struct {
	Name string `json:"name,omitempty"`
}
