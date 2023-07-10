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
	usecases := new(UseCases)

	usecases.Set = NewSetUseCase(repository)
	usecases.Source = NewSourceUseCase(repository)
	usecases.User = NewUserUseCase(repository)
	usecases.Spell = NewSpellUseCase(repository)
	usecases.UrlSet = NewUrlSetUseCase(repository, usecases)

	return usecases
}
