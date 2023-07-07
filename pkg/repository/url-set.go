package repository

import (
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"spells.tvblackman1.ru/lib/pagination"
	"spells.tvblackman1.ru/pkg/domain/dto"
	dbdto "spells.tvblackman1.ru/pkg/repository/dto"
	"spells.tvblackman1.ru/pkg/repository/requests"
	"spells.tvblackman1.ru/pkg/repository/requests/fields"
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
		fmt.Printf("Bad request: %s while creating an url set. Request: %s\n", err.Error(), sqlRequest)
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
	fromDb := dbdto.UrlSetDb{}
	sqlRequest, _, _ := request.ToSQL()
	err := rep.db.Get(&fromDb, sqlRequest)
	if err != nil {
		fmt.Printf("Bad request: %s while getting an url set. Request: %s\n", err.Error(), sqlRequest)
		return dto.UrlSetDto{}, err
	}
	return dto.UrlSetDto{
		Id:   dto.UrlSetId(fromDb.Id),
		Uri:  fromDb.Uri,
		Name: fromDb.Name,
	}, nil
}

func (rep *UrlSetsRepository) GetSpells(id dto.UrlSetId, params dto.SearchSpellDto, pag pagination.Pagination) ([]dto.SpellDto, pagination.Meta, error) {
	// TODO add lib for page numbers
	if pag.PageNumber < 1 {
		pag.PageNumber = 1
	}
	limit := pag.Limit
	offset := pag.Limit * (pag.PageNumber - 1)

	urlSetId := uuid.UUID(id).String()
	allSpellsRequest := requests.SelectSpellsWithSourceName(params)
	usedSpellsRequest := goqu.Dialect("postgres").
		From(fields.UrlSetToSpell().T()).
		Where(fields.UrlSetToSpell().UrlSetId().Eq(urlSetId)).
		InnerJoin(allSpellsRequest.As("all"),
			goqu.On(fields.UrlSetToSpell().SpellId().Eq(fields.Spell().Aliased("all").Id())))
	usedSpellsRequest = usedSpellsRequest.Select(goqu.I("all.*"))

	requestCount := requests.CountRows(usedSpellsRequest)

	usedSpellsRequest = usedSpellsRequest.Order(goqu.C("level").Asc())
	// TODO usedSpellsRequest = usedSpellsRequest.Order(fields.Spell().Level().Asc()) like this
	usedSpellsRequest = usedSpellsRequest.OrderAppend(goqu.C("spells_name").Asc())

	usedSpellsRequest = usedSpellsRequest.Limit(uint(limit)).Offset(uint(offset))

	var spells []dbdto.SpellDb
	var count int
	sqlCountRequest, _, _ := requestCount.ToSQL()
	sqlRequest, _, _ := usedSpellsRequest.ToSQL()

	_ = rep.db.Get(&count, sqlCountRequest)
	err := rep.db.Select(&spells, sqlRequest)
	if err != nil {
		fmt.Printf("Bad request: %s\n", err.Error())
		fmt.Println(sqlRequest)
		return []dto.SpellDto{}, pagination.Meta{}, err
	}
	ret := make([]dto.SpellDto, len(spells))
	for i := range ret {
		ret[i] = dbdto.DbSpellToSpellDto(spells[i])
	}
	meta := pagination.GetMeta(limit, count, pag.PageNumber)
	return ret, meta, nil
}

func (rep *UrlSetsRepository) GetAllSpells(id dto.UrlSetId, params dto.SearchSpellDto, pag pagination.Pagination) ([]dto.SpellMarkedDto, pagination.Meta, error) {
	if pag.PageNumber < 1 {
		pag.PageNumber = 1
	}
	limit := pag.Limit
	offset := pag.Limit * (pag.PageNumber - 1)

	urlSetId := uuid.UUID(id).String()
	allSpellsRequest := requests.SelectSpellsWithSourceName(dto.SearchSpellDto{})

	usedSpellsRequest := goqu.Dialect("postgres").
		From(fields.UrlSetToSpell().T()).
		Where(fields.UrlSetToSpell().UrlSetId().Eq(urlSetId)).
		LeftJoin(fields.Spell().T(), goqu.On(fields.UrlSetToSpell().SpellId().Eq(fields.Spell().Id())))
	allSpellsWithMark := allSpellsRequest.LeftJoin(usedSpellsRequest.As("used"), goqu.On(
		fields.UrlSetToSpell().Aliased("used").SpellId().
			Eq(
				fields.Spell().Id())))
	allSpellsWithMark = allSpellsWithMark.SelectAppend(
		fields.UrlSetToSpell().Aliased("used").UrlSetId().IsNotNull().As("in_set"))

	requestCount := requests.CountRows(allSpellsWithMark)

	allSpellsWithMark = allSpellsRequest.Order(fields.Spell().Level().Asc())
	allSpellsWithMark = allSpellsWithMark.Order(fields.Spell().Name().Asc())
	allSpellsWithMark = allSpellsWithMark.Limit(uint(limit)).Offset(uint(offset))
	sqlRequest, _, _ := allSpellsWithMark.ToSQL()

	fmt.Println(sqlRequest)

	var spells []dbdto.SpellMarkedDb
	var count int
	sqlCountRequest, _, _ := requestCount.ToSQL()
	_ = rep.db.Get(&count, sqlCountRequest)
	err := rep.db.Select(&spells, sqlRequest)
	if err != nil {
		fmt.Printf("Bad request: %s\n", err.Error())
		fmt.Println(sqlRequest)
		return []dto.SpellMarkedDto{}, pagination.Meta{}, err
	}
	ret := make([]dto.SpellMarkedDto, len(spells), len(spells))
	for i := range ret {
		ret[i] = dbdto.DbSpellMarkedToSpellMarkedDto(spells[i])
	}
	meta := pagination.GetMeta(limit, count, pag.PageNumber)
	return ret, meta, nil
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
		fmt.Printf("Bad request: %s while renaming an url set. Request: %s\n", err.Error(), sqlRequest)
		return err
	}
	return nil
}

func (rep *UrlSetsRepository) AddSpell(id dto.UrlSetId, spellId dto.SpellId) error {
	dialect := goqu.Dialect("postgres")
	request := dialect.
		Insert(fields.UrlSetToSpell().T()).Rows(
		goqu.Record{
			"id":         uuid.New().String(),
			"url_set_id": uuid.UUID(id).String(),
			"spell_id":   uuid.UUID(spellId).String(),
		}).
		Returning("id")
	sqlRequest, _, _ := request.ToSQL()
	var uuidStr string
	err := rep.db.Get(&uuidStr, sqlRequest)
	if err != nil {
		fmt.Printf("Bad request: %s while adding spell to url set. Request: %s\n", err.Error(), sqlRequest)
		return err
	}
	return nil
}

func (rep *UrlSetsRepository) RemoveSpell(id dto.UrlSetId, _spellId dto.SpellId) error {
	urlSetId := uuid.UUID(id).String()
	spellId := uuid.UUID(_spellId).String()
	dialect := goqu.Dialect("postgres")
	request := dialect.From(fields.UrlSetToSpell().T()).
		Where(fields.UrlSetToSpell().UrlSetId().Eq(urlSetId)).
		Where(fields.UrlSetToSpell().SpellId().Eq(spellId)).Delete()
	sqlRequest, _, _ := request.ToSQL()
	var uuidStr string
	err := rep.db.Get(&uuidStr, sqlRequest)
	if err != nil {
		fmt.Printf("Bad request: %s while deleting spell to url set. Request: %s\n", err.Error(), sqlRequest)
		return err
	}
	return nil
}
