package repository

import (
	"fmt"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"spells.tvblackman1.ru/lib/pagination"
	"spells.tvblackman1.ru/pkg/domain/dto"
	dbdto "spells.tvblackman1.ru/pkg/repository/dto"
	"spells.tvblackman1.ru/pkg/repository/requests"
	"strings"
)

type SpellsRepository struct {
	db *sqlx.DB
}

func NewSpellsRepository(db *sqlx.DB) *SpellsRepository {
	return &SpellsRepository{db}
}

func (rep *SpellsRepository) CreateSpell(spellDto dto.SpellToRepositoryDto) error {
	dialect := goqu.Dialect("postgres")
	request := dialect.
		Insert(SpellsDbName).
		Rows(
			goqu.Record{
				"id":               uuid.UUID(spellDto.Id).String(),
				"name":             spellDto.Name,
				"level":            spellDto.Level,
				"description":      spellDto.Description,
				"classes":          strings.Join(spellDto.Classes, ", "),
				"casting_time":     spellDto.CastingTime,
				"duration":         spellDto.Duration,
				"is_verbal":        spellDto.IsVerbal,
				"is_somatic":       spellDto.IsSomatic,
				"is_material":      spellDto.HasMaterialComponent,
				"material_content": spellDto.MaterialComponent,
				"magical_school":   spellDto.MagicalSchool,
				"distance":         spellDto.Distance,
				"is_ritual":        spellDto.IsRitual,
				"source_id":        uuid.UUID(spellDto.SourceId).String(),
			}).
		Returning("id")
	sqlRequest, _, _ := request.ToSQL()
	var uuidStr string
	err := rep.db.Get(&uuidStr, sqlRequest)
	if err != nil {
		fmt.Println(sqlRequest)
		return err
	}
	return nil
}

func (rep *SpellsRepository) GetSpells(params dto.SearchSpellDto, pagination pagination.Pagination) ([]dto.SpellDto, error) {
	if pagination.PageNumber < 1 {
		pagination.PageNumber = 1
	}
	limit := pagination.Limit
	offset := pagination.Limit * (pagination.PageNumber - 1)

	request := requests.SelectSpellsWithSourceName(params)
	request = request.Order(goqu.C("spells_name").Asc())
	request = request.Limit(uint(limit)).Offset(uint(offset))
	var spells []dbdto.SpellDb
	sqlRequest, _, _ := request.ToSQL()
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

func (rep *SpellsRepository) GetById(_ dto.SpellId) dto.SpellDto {
	return dto.SpellDto{}
}
