package boundaries

import "spells.tvblackman1.ru/pkg/domain/dto"

type Repository struct {
	Tags    TagsRepository
	Users   UsersRepository
	Sources SourcesRepository
	Sets    SetsRepository
	Spells  SpellsRepository
}

type TagsRepository interface {
}

type UsersRepository interface {
	CreateUser(dto dto.UserToRepositoryDto)
	GetById(id dto.UserId) dto.UserDto
}

type SourcesRepository interface {
	CreateSource(sourceDto dto.SourceToRepositoryDto)
	GetById(id dto.SourceId) dto.SourceDto
	//AddCustomSourceToUser(userId dto.UserId, dto.SourceId)
	//CreateCopyWithNextVersion(userId dto.UserId, id dto.SourceId)
	//CloneSource(userId dto.UserId, id dto.SourceId)
}

type SetsRepository interface {
	CreateSet(setDto dto.SetToRepositoryDto)
	GetById(id dto.SetId) dto.SetDto
	UpdateSpellList(id dto.SetId, dto dto.UpdateSetSpellListDto)
	GetSpells(id dto.SetId, params dto.SearchSpellDto) []dto.SetSpellDto
	GetSetsByName(name string) []dto.SetDto
	GetSpell(params dto.SetSpellId) dto.SetSpellDto
	EditSpellComments(id dto.SetSpellId, setDto dto.EditSpellInSetDto)
}

type SpellsRepository interface {
	CreateSpell(spellDto dto.SpellToRepositoryDto)
	GetById(id dto.SpellId) dto.SpellDto
	GetSpells(params dto.SearchSpellDto) []dto.SpellDto
	//UpdateSpell(userId dto.UserId, spellDto dto.CreateSpellDto)
}