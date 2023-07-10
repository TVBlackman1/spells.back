package usecases

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"os"
	"spells.tvblackman1.ru/lib/pagination"
	"spells.tvblackman1.ru/pkg/domain/boundaries"
	"spells.tvblackman1.ru/pkg/domain/dto"
)

type SpellUseCase struct {
	repository *boundaries.Repository
}

func NewSpellUseCase(repository *boundaries.Repository) *SpellUseCase {
	return &SpellUseCase{repository}
}

func (usecase *SpellUseCase) CreateSpell(userId dto.UserId, spellDto dto.CreateSpellDto) error {
	verificationOff := os.Getenv("UNSAVED_SPELL_CREATING") == "true"
	if !verificationOff {
		if !usecase.isUserHavingSpellSource(userId, spellDto.SourceId) {
			return errors.New("not valid user")
		}
		if !usecase.checkMaterialComponent(spellDto) {
			println(spellDto.Name)
			return errors.New("not valid material component")
		}
		unique, err := usecase.isNewNameInSpellSource(spellDto.Name, spellDto.SourceId)
		if err != nil {
			return errors.New(err.Error())
		}
		if !unique {
			errorText := fmt.Sprintf("not unique name of spell in source: [%s]", spellDto.Name)
			return errors.New(errorText)
		}
	}

	dataToWrite := dto.SpellToRepositoryDto{
		Id:                   dto.SpellId(uuid.New()),
		Name:                 spellDto.Name,
		Level:                spellDto.Level,
		Classes:              spellDto.Classes,
		Version:              1,
		Description:          spellDto.Description,
		CastingTime:          spellDto.CastingTime,
		Duration:             spellDto.Duration,
		IsVerbal:             spellDto.IsVerbal,
		IsSomatic:            spellDto.IsSomatic,
		HasMaterialComponent: spellDto.HasMaterialComponent,
		MaterialComponent:    spellDto.MaterialComponent,
		MagicalSchool:        spellDto.MagicalSchool,
		Distance:             spellDto.Distance,
		IsRitual:             spellDto.IsRitual,
		SourceId:             spellDto.SourceId,
	}
	err := usecase.repository.Spells.CreateSpell(dataToWrite)
	return err
}

func (usecase *SpellUseCase) GetSpellList(userId dto.UserId, searchDto dto.SearchSpellDto, pag pagination.Pagination) ([]dto.SpellDto, pagination.Meta, error) {
	if len(searchDto.Sources) == 0 {
		// TODO default sources + user custom sources (extern libs or written by this user)
	}
	return usecase.repository.Spells.GetSpells(searchDto, pag)
}

func (usecase *SpellUseCase) GetCommonSpellList(searchDto dto.SearchSpellDto, pag pagination.Pagination) ([]dto.SpellDto, pagination.Meta, error) {
	if len(searchDto.Sources) == 0 {
		// TODO default sources + user custom sources (extern libs or written by this user)
	}
	return usecase.repository.Spells.GetSpells(searchDto, pag)
}

func (usecase *SpellUseCase) checkMaterialComponent(spellDto dto.CreateSpellDto) bool {
	hasMaterialComponentText := len(spellDto.MaterialComponent) > 0
	return hasMaterialComponentText == spellDto.HasMaterialComponent
}

func (usecase *SpellUseCase) isUserHavingSpellSource(userId dto.UserId, sourceId dto.SourceId) bool {
	source, err := usecase.repository.Sources.GetById(sourceId)
	if err != nil {
		return false
	}
	return source.UploadedBy == userId
}

func (usecase *SpellUseCase) isNewNameInSpellSource(spellName string, sourceId dto.SourceId) (bool, error) {
	spells, _, err := usecase.repository.Spells.GetSpells(dto.SearchSpellDto{
		EqualsName: spellName,
		Sources:    []dto.SourceId{sourceId},
	}, pagination.Pagination{
		Limit: 1,
	})
	if err != nil {
		return false, err
	}
	for _, spell := range spells {
		if spell.Name == spellName {
			return false, nil
		}
	}
	return true, nil
}
