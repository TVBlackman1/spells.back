package dto

import "github.com/google/uuid"

type SpellId uuid.UUID
type SetSpellId uuid.UUID

type SpellDto struct {
	Id                   SpellId
	Name                 string
	Level                int
	Classes              []string
	Version              int
	Description          string
	Action               string
	Duration             string
	IsVerbal             bool
	IsSomatic            bool
	HasMaterialComponent bool
	MaterialComponent    string
	MagicalSchool        string
	Distance             string
	IsRitual             bool
	SourceIds            []SourceId
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
	SourceIds            []SourceId
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
	SourceIds            []SourceId
}

type SearchSpellDto struct {
	Name                 string
	IsRitual             bool
	IsVerbal             bool
	IsSomatic            bool
	HasMaterialComponent bool
	MagicalSchools       []string
	Levels               []int
	Classes              []string
	Sources              []SourceId
}
