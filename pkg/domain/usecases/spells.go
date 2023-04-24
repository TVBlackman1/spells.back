package usecases

import (
	"errors"
	"github.com/google/uuid"
	"os"
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
		if !usecase.isUserHavingSpellSource(userId, spellDto.SourceIds) {
			return errors.New("not valid user")
		}
		if !usecase.checkMaterialComponent(spellDto) {
			return errors.New("not valid material component")
		}
		if !usecase.isNewNameInSpellSource(spellDto.Name, spellDto.SourceIds) {
			return errors.New("not unique name of spell in source")
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
		SourceIds:            spellDto.SourceIds,
	}
	err := usecase.repository.Spells.CreateSpell(dataToWrite)
	return err
}

func (usecase *SpellUseCase) GetSpellList(userId dto.UserId, searchDto dto.SearchSpellDto) ([]dto.SpellDto, error) {
	if len(searchDto.Sources) == 0 {
		// TODO default sources + user custom sources (extern libs or written by this user)
	}
	return usecase.repository.Spells.GetSpells(searchDto), nil
}

func (usecase *SpellUseCase) checkMaterialComponent(spellDto dto.CreateSpellDto) bool {
	hasMaterialComponentText := len(spellDto.MaterialComponent) > 0
	return hasMaterialComponentText == spellDto.HasMaterialComponent
}

func (usecase *SpellUseCase) isUserHavingSpellSource(userId dto.UserId, sourceIds []dto.SourceId) bool {
	for _, sourceId := range sourceIds {
		source := usecase.repository.Sources.GetById(sourceId)
		if source.UploadedBy != userId {
			return false
		}
	}
	return true
}

func (usecase *SpellUseCase) isNewNameInSpellSource(spellName string, sourceIds []dto.SourceId) bool {
	// TODO examination: if already created - expand list of sources in spell dto
	// now just check first source id
	spells := usecase.repository.Spells.GetSpells(dto.SearchSpellDto{
		Name:    spellName,
		Sources: sourceIds[:1],
	})
	for _, spell := range spells {
		if spell.Name == spellName {
			return false
		}
	}
	return true
}
