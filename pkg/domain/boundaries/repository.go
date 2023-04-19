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
}

type SourcesRepository interface {
}

type SetsRepository interface {
}

type SpellsRepository interface {
}
