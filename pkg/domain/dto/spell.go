package dto

import (
	"github.com/google/uuid"
	"spells.tvblackman1.ru/lib/tribool"
)

type SpellId uuid.UUID
type SetSpellId uuid.UUID

type SpellDto struct {
	Id                   SpellId
	Name                 string
	Level                int
	Classes              []string
	Version              int
	Description          string
	CastingTime          string
	Duration             string
	IsVerbal             bool
	IsSomatic            bool
	HasMaterialComponent bool
	MaterialComponent    string
	MagicalSchool        string
	Distance             string
	IsRitual             bool
	SourceId             SourceId
	SourceName           string
}

type SpellMarkedDto struct {
	SpellDto
	InSet bool
}

type SetSpellDto struct {
	Id                  SetSpellId
	SetId               SetId
	Original            SpellDto
	MasterComment       string
	VisualCustomization string
}

type UpdateSetSpellListDto struct {
	SetId        SetId
	ListToRemove []SpellId
	ListToAdd    []SpellId
}

type EditSpellInSetDto struct {
	MasterComment       string
	VisualCustomization string
}

type CreateSpellDto struct {
	Name                 string
	Level                int
	Classes              []string
	Description          string
	CastingTime          string
	Duration             string
	IsVerbal             bool
	IsSomatic            bool
	HasMaterialComponent bool
	MaterialComponent    string
	MagicalSchool        string
	Distance             string
	IsRitual             bool
	SourceId             SourceId
}

type SpellToRepositoryDto struct {
	Id                   SpellId
	Name                 string
	Level                int
	Classes              []string
	Version              int
	Description          string
	CastingTime          string
	Duration             string
	IsVerbal             bool
	IsSomatic            bool
	HasMaterialComponent bool
	MaterialComponent    string
	MagicalSchool        string
	Distance             string
	IsRitual             bool
	SourceId             SourceId
}

type SearchSpellDto struct {
	Id                   SpellId
	LikeName             string
	EqualsName           string
	IsRitual             tribool.Tribool
	IsVerbal             tribool.Tribool
	IsSomatic            tribool.Tribool
	HasMaterialComponent tribool.Tribool
	WasteMaterial        tribool.Tribool // TODO
	MagicalSchools       []string
	Levels               []int
	Classes              []string // TODO
	Sources              []SourceId
}
