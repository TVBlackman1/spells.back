package usecases

import "spells.tvblackman1.ru/pkg/domain/boundaries"

type UseCases struct {
	Set    *SetUseCase
	Source *SourceUseCase
	User   *UserUseCase
	Spell  *SpellUseCase
	UrlSet *UrlSetUseCase
}

func NewUseCases(repository *boundaries.Repository) *UseCases {
	return &UseCases{
		Set:    NewSetUseCase(repository),
		Source: NewSourceUseCase(repository),
		User:   NewUserUseCase(repository),
		Spell:  NewSpellUseCase(repository),
		UrlSet: NewUrlSetUseCase(repository),
	}
}
