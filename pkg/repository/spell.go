package repository

import (
	"fmt"
	"spells.tvblackman1.ru/pkg/domain/dto"
)

type SpellsRepository struct {
}

func NewSpellsRepository() *SpellsRepository {
	return &SpellsRepository{}
}

func (rep *SpellsRepository) CreateSpell(spellDto dto.SpellToRepositoryDto) {
	fmt.Printf("created: %+v", spellDto)
}

func (rep *SpellsRepository) GetSpells(params dto.SearchSpellDto) []dto.SpellDto {
	return []dto.SpellDto{}
}

func (rep *SpellsRepository) GetById(id dto.SpellId) dto.SpellDto {
	return dto.SpellDto{}
}
