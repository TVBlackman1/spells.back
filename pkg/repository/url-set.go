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

func (rep *UrlSetsRepository) GetSpells(id dto.UrlSetId, params dto.SearchSpellDto, pagination pagination.Pagination) ([]dto.SpellDto, error) {
	dialect := goqu.Dialect("postgres")
	spellsWithSourceRequest := dialect.
		Select("spells.id", goqu.T("spells").Col("name").As("spells_name"), "spells.level", "spells.description", "spells.casting_time",
			"spells.duration", "spells.is_verbal", "spells.is_somatic", "spells.is_material",
			"spells.material_content", "spells.magical_school", "spells.distance", "spells.is_ritual", "spells.source_id",
			goqu.T("sources").Col("name").As("sources_name")).
		From(SpellsDbName).
		LeftJoin(goqu.T(SourcesDbName), goqu.On(goqu.I("spells.source_id").Eq(goqu.I("sources.id"))))

	spellsInUrlSet := dialect.From(UrlSetsToSpellsDbName).Select("spells_with_source.id", "spells_with_source.spells_name", "spells_with_source.level", "spells_with_source.description", "spells_with_source.casting_time",
		"spells_with_source.duration", "spells_with_source.is_verbal", "spells_with_source.is_somatic", "spells_with_source.is_material",
		"spells_with_source.material_content", "spells_with_source.magical_school", "spells_with_source.distance", "spells_with_source.is_ritual", "spells_with_source.source_id",
		"spells_with_source.sources_name").
		LeftJoin(spellsWithSourceRequest.As("spells_with_source"),
			goqu.On(goqu.I("spells_with_source.id").Eq(goqu.I(fmt.Sprintf("%s.spell_id", UrlSetsToSpellsDbName)))))
	var spells []dbdto.SpellDb
	sqlRequest, _, _ := spellsInUrlSet.ToSQL()
	err := rep.db.Select(&spells, sqlRequest)
	if err != nil {
		fmt.Printf("Bad request: %s\n", err.Error())
		fmt.Println(sqlRequest)
		return []dto.SpellDto{}, err
	}
	ret := make([]dto.SpellDto, len(spells))
	for i := range ret {
		ret[i] = dbdto.DbSpellToSpellDto(spells[i])
	}
	return ret, nil
}

func (rep *UrlSetsRepository) GetAllSpells(id dto.UrlSetId, params dto.SearchSpellDto, pagination pagination.Pagination) ([]dto.SpellMarkedDto, error) {

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
	allSpellsWithMark = allSpellsWithMark.Order(fields.Spell().Name().Asc()).Limit(10)
	sqlRequest, _, _ := allSpellsWithMark.ToSQL()
	var spells []dbdto.SpellMarkedDb
	err := rep.db.Select(&spells, sqlRequest)
	if err != nil {
		fmt.Printf("Bad request: %s\n", err.Error())
		fmt.Println(sqlRequest)
		return []dto.SpellMarkedDto{}, err
	}
	ret := make([]dto.SpellMarkedDto, len(spells), len(spells))
	for i := range ret {
		ret[i] = dbdto.DbSpellMarkedToSpellMarkedDto(spells[i])
	}
	return ret, nil
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
		Insert(UrlSetsToSpellsDbName).Rows(
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

func (rep *UrlSetsRepository) RemoveSpell(id dto.UrlSetId, spellId dto.SpellId) error {
	return nil
}
