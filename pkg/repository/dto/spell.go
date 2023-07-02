package dbdto

import (
	"database/sql"
	"github.com/google/uuid"
)

type SpellDb struct {
	Id              uuid.UUID      `db:"id"`
	Name            string         `db:"spells_name"`
	Level           int            `db:"level"`
	Description     string         `db:"description"`
	CastingTime     string         `db:"casting_time"`
	MagicalSchool   string         `db:"magical_school"`
	Duration        string         `db:"duration"`
	IsVerbal        bool           `db:"is_verbal"`
	IsSomatic       bool           `db:"is_somatic"`
	HasMaterial     bool           `db:"is_material"`
	MaterialContent sql.NullString `db:"material_content"`
	IsRitual        bool           `db:"is_ritual"`
	Distance        string         `db:"distance"`
	SourceId        uuid.UUID      `db:"source_id"`
	SourceName      string         `db:"sources_name"`
}

type SpellMarkedDb struct {
	SpellDb
	InSet bool `db:"in_set"`
}
