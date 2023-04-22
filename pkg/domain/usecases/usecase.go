package usecases

import "spells.tvblackman1.ru/pkg/domain/boundaries"

type UseCases struct {
	Set    *SetUseCase
	Source *SourceUseCase
	User   *UserUseCase
	Spell  *SpellUseCase
}

func NewUseCases(repository *boundaries.Repository) *UseCases {
	return &UseCases{
		Set:    NewSetUseCase(repository),
		Source: NewSourceUseCase(repository),
		User:   NewUserUseCase(repository),
		Spell:  NewSpellUseCase(repository),
	}
}
