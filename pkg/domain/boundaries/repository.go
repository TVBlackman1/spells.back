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
	CreateUser(dto dto.UserToRepositoryDto) error
	GetById(id dto.UserId) (dto.UserDto, error)
	GetUsers(params dto.SearchUserDto) ([]dto.UserDto, error)
}

type SourcesRepository interface {
	CreateSource(sourceDto dto.SourceToRepositoryDto) error
	GetById(id dto.SourceId) (dto.SourceDto, error)
	GetSources(userId dto.UserId, params dto.SearchSourceDto) ([]dto.SourceDto, error)
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
	CreateSpell(spellDto dto.SpellToRepositoryDto) error
	GetById(id dto.SpellId) dto.SpellDto
	GetSpells(params dto.SearchSpellDto) []dto.SpellDto
	//UpdateSpell(userId dto.UserId, spellDto dto.CreateSpellDto)
}
