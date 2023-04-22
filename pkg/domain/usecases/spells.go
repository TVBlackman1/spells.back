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
		if !usecase.isUserHavingSpellSource(userId, spellDto.SourceId) {
			return errors.New("not valid user")
		}
		if !usecase.checkMaterialComponent(spellDto) {
			return errors.New("not valid material component")
		}
	}
	println("!!")

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

func (usecase *SpellUseCase) GetSpellList(userId dto.UserId, spellDto dto.SearchSpellDto) ([]dto.SpellDto, error) {
	if len(spellDto.Sources) == 0 {
		//TODO default sources + user custom sources (extern libs or written by this user)
	}
	return usecase.repository.Spells.GetSpells(spellDto), nil
}

func (usecase *SpellUseCase) checkMaterialComponent(spellDto dto.CreateSpellDto) bool {
	hasMaterialComponentText := len(spellDto.MaterialComponent) > 0
	return hasMaterialComponentText == spellDto.HasMaterialComponent
}

func (usecase *SpellUseCase) isUserHavingSpellSource(userId dto.UserId, sourceId dto.SourceId) bool {
	source := usecase.repository.Sources.GetById(sourceId)
	return source.UploadedBy == userId
}
