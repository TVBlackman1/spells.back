package usecases

import (
	"errors"
	"github.com/google/uuid"
	"spells.tvblackman1.ru/pkg/domain/boundaries"
	"spells.tvblackman1.ru/pkg/domain/dto"
)

type SetUseCase struct {
	repository *boundaries.Repository
}

func NewSetUseCase(repository *boundaries.Repository) *SetUseCase {
	return &SetUseCase{repository}
}

func (usecase *SetUseCase) CreateSet(userId dto.UserId, setDto dto.CreateSetDto) error {
	if !usecase.isValidSetName(userId, setDto.Name) {
		return errors.New("not valid set name")
	}

	dataToWrite := dto.SetToRepositoryDto{
		Id:          dto.SetId(uuid.New()),
		Name:        setDto.Name,
		UserId:      userId,
		Description: setDto.Description,
		Sources:     setDto.Sources,
	}
	usecase.repository.Sets.CreateSet(dataToWrite)
	return nil
}

func (usecase *SetUseCase) GetSpells(id dto.SetId, params dto.SearchSpellDto) {

}

func (usecase *SetUseCase) EditSpellList(userId dto.UserId, setId dto.SetId, spells dto.UpdateSetSpellListDto) error {
	set := usecase.repository.Sets.GetById(setId)
	if set.UserId != userId {
		return errors.New("not valid user")
	}
	usecase.repository.Sets.UpdateSpellList(setId, spells)
	return nil
}

func (usecase *SetUseCase) EditSpellComments(userId dto.UserId, spellId dto.SetSpellId, refactor dto.EditSpellInSetDto) error {
	setSpell := usecase.repository.Sets.GetSpell(spellId)
	set := usecase.repository.Sets.GetById(setSpell.SetId)
	setOwner := set.UserId
	if setOwner != userId {
		errors.New("not valid user")
	}
	usecase.repository.Sets.EditSpellComments(spellId, refactor)
	return nil
}

func (usecase *SetUseCase) isValidSetName(userId dto.UserId, name string) bool {
	if len(name) == 0 {
		return false
	}
	setsWithSimilarNames := usecase.repository.Sets.GetSetsByName(name)
	for _, set := range setsWithSimilarNames {
		if set.Name == name && set.UserId == userId {
			return false
		}
	}
	return true
}
