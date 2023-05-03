package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"spells.tvblackman1.ru/pkg/domain/dto"
	"strings"
)

type SpellsRepository struct {
	db *sqlx.DB
}

func NewSpellsRepository(db *sqlx.DB) *SpellsRepository {
	return &SpellsRepository{db}
}

func (rep *SpellsRepository) CreateSpell(spellDto dto.SpellToRepositoryDto) error {

	request := fmt.Sprintf("INSERT INTO %s(id, name, level, "+
		"description, classes, casting_time, duration, is_verbal, is_somatic, is_material, material_content, magical_school, "+
		"distance, is_ritual, source_id) VALUES ('%s', '%s', '%d', '%s', '%s', '%s', '%s', '%t', '%t', '%t', '%s', "+
		"'%s', '%s', '%t', '%s') RETURNING id;\n",
		SpellsDbName,
		uuid.UUID(spellDto.Id).String(),
		spellDto.Name,
		spellDto.Level,
		spellDto.Description,
		strings.Join(spellDto.Classes, ", "),
		spellDto.CastingTime,
		spellDto.Duration,
		spellDto.IsVerbal,
		spellDto.IsSomatic,
		spellDto.HasMaterialComponent,
		spellDto.MaterialComponent,
		spellDto.MagicalSchool,
		spellDto.Distance,
		spellDto.IsRitual,
		uuid.UUID(spellDto.SourceId).String(),
	)
	var uuidStr string
	err := rep.db.Get(&uuidStr, request)
	if err != nil {
		fmt.Println(request)
		return err
	}
	return nil
}

func (rep *SpellsRepository) GetSpells(params dto.SearchSpellDto) []dto.SpellDto {
	return []dto.SpellDto{}
}

func (rep *SpellsRepository) GetById(id dto.SpellId) dto.SpellDto {
	return dto.SpellDto{}
}
