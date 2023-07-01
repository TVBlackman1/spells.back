package repository

import (
	"github.com/jmoiron/sqlx"
	"spells.tvblackman1.ru/pkg/domain/boundaries"
)

func NewRepository(db *sqlx.DB) (*boundaries.Repository, error) {
	repo := new(boundaries.Repository)
	repo.Users = NewUserRepository(db)
	repo.Sources = NewSourcesRepository(db)
	repo.Sets = NewSetsRepository(db)
	repo.Spells = NewSpellsRepository(db)
	repo.UrlSets = NewUrlSetsRepository(db)

	return repo, nil
}
