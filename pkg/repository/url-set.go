package repository

import (
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"os"
	"spells.tvblackman1.ru/lib/pagination"
	"spells.tvblackman1.ru/pkg/domain/dto"
)

type UrlSetsRepository struct {
	db *sqlx.DB
}

func NewUrlSetsRepository(db *sqlx.DB) *UrlSetsRepository {
	return &UrlSetsRepository{db}
}

func (rep *UrlSetsRepository) CreateUrlSet(repositoryDto dto.UrlSetToRepositoryDto) error {
	dialect := goqu.Dialect("postgres")
	request := dialect.Insert(UrlSetsDbName).Rows(
		goqu.Record{
			"id":   uuid.UUID(repositoryDto.Id).String(),
			"url":  repositoryDto.Uri,
			"name": "empty",
		}).Returning("id")

	sqlRequest, _, _ := request.ToSQL()
	var uuidStr string
	err := rep.db.Get(&uuidStr, sqlRequest)
	if err != nil {
		fmt.Println(os.Stderr, "Bad requests: %s while creating an url set. Request: ", err.Error(), sqlRequest)
		return err
	}
	fmt.Println("Created new url:", repositoryDto.Uri)
	return nil
}

func (rep *UrlSetsRepository) GetById(id dto.SpellId) (dto.UrlSetDto, error) {
	return dto.UrlSetDto{}, nil
}

func (rep *UrlSetsRepository) GetByLink(link string) (dto.UrlSetDto, error) {
	dialect := goqu.Dialect("postgres")
	request := dialect.
		Select("id", "url", "name").
		From(UrlSetsDbName).
		Where(goqu.C("url").Eq(link)).
		Limit(1)
	fromDb := UrlSetDb{}
	sqlRequest, _, _ := request.ToSQL()
	err := rep.db.Get(&fromDb, sqlRequest)
	if err != nil {
		fmt.Println(os.Stderr, "Bad requests: %s while getting an url set. Request: ", err.Error(), sqlRequest)
		return dto.UrlSetDto{}, err
	}
	return dto.UrlSetDto{
		Id:   dto.UrlSetId(fromDb.Id),
		Uri:  fromDb.Uri,
		Name: fromDb.Name,
	}, nil
}

func (rep *UrlSetsRepository) GetSpells(id dto.UrlSetId, params dto.SearchSpellDto, pagination pagination.Pagination) ([]dto.SpellDto, error) {
	return []dto.SpellDto{}, nil
}

func (rep *UrlSetsRepository) RenameUrlSet(id dto.UrlSetId, newName string) error {
	dialect := goqu.Dialect("postgres")
	urlId := uuid.UUID(id).String()
	request := dialect.
		Update(UrlSetsDbName).
		Set(goqu.Record{"name": newName}).
		Where(goqu.C("id").Eq(urlId)).
		Limit(1).Returning("id")
	sqlRequest, _, _ := request.ToSQL()
	var uuidStr string
	err := rep.db.Get(&uuidStr, sqlRequest)
	if err != nil {
		fmt.Println(os.Stderr, "Bad requests: %s while renaming an url set. Request: ", err.Error(), sqlRequest)
		return err
	}
	return nil
}

func (rep *UrlSetsRepository) AddSpell(id dto.UrlSetDto, spellId dto.SpellId) error {
	return nil
}

func (rep *UrlSetsRepository) RemoveSpell(id dto.UrlSetDto, spellId dto.SpellId) error {
	return nil
}

type UrlSetDb struct {
	Id   uuid.UUID `db:"id"`
	Name string    `db:"name"`
	Uri  string    `db:"url"`
}
