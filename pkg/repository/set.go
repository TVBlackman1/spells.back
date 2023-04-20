package repository

import (
	"fmt"
	"github.com/google/uuid"
	"spells.tvblackman1.ru/pkg/domain/dto"
)

type SetsRepository struct {
}

func NewSetsRepository() *SetsRepository {
	return &SetsRepository{}
}

func (rep *SetsRepository) CreateSet(setDto dto.SetToRepositoryDto) {
	fmt.Printf("created: %+v", setDto)
}

func (rep *SetsRepository) GetById(id dto.SetId) dto.SetDto {
	return dto.SetDto{}
}

func (rep *SetsRepository) UpdateSpellList(id dto.SetId, dto dto.UpdateSetSpellListDto) {
	fmt.Printf("updated list of spells in set: %+v", id)
}

func (rep *SetsRepository) GetSpells(id dto.SetId, params dto.SearchSpellDto) []dto.SetSpellDto {
	return []dto.SetSpellDto{
		{
			Id:                  dto.SetSpellId(uuid.New()),
			Original:            dto.SpellDto{},
			MasterComment:       "master-comment",
			VisualCustomization: "visual-comment",
		},
	}
}

func (rep *SetsRepository) GetSetsByName(name string) []dto.SetDto {
	return []dto.SetDto{}
}

func (rep *SetsRepository) GetSpell(params dto.SetSpellId) dto.SetSpellDto {
	return dto.SetSpellDto{}
}

func (rep *SetsRepository) EditSpellComments(id dto.SetSpellId, setDto dto.EditSpellInSetDto) {
	fmt.Printf("edited spell comments: %+v", id)
}
