package repository

import (
	"database/sql"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"os"
	"spells.tvblackman1.ru/lib/pagination"
	"spells.tvblackman1.ru/lib/tribool"
	"spells.tvblackman1.ru/pkg/domain/dto"
	"strconv"
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
	sql, _, _ := request.ToSQL()
	//fmt.Println(sql)
	var uuidStr string
	err := rep.db.Get(&uuidStr, sql)
	if err != nil {
		fmt.Println(sql)
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

	dialect := goqu.Dialect("postgres")
	request := dialect.
		Select("spells.id", "spells.name as \"spells_name\"", "spells.level", "spells.description", "spells.casting_time").
		Select("spells.duration", "spells.is_verbal", "spells.is_somatic", "spells.is_material").
		Select("spells.material_content", "spells.magical_school", "spells.distance", "spells.is_ritual").
		Select("spells.source_id").
		From(SpellsDbName).
		LeftJoin(goqu.T(SourcesDbName), goqu.On(goqu.I("spells.source_id").Eq(goqu.I("sources.id")))).
		Select(goqu.T("sources").Col("name").As("sources_name"))
	if len(params.Sources) > 0 {
		sourcesToRequest := getSourcesEnumeration(params.Sources)
		request.Where(goqu.C("sources.id").In(sourcesToRequest))
	}
	if len(params.EqualsName) > 0 {
		request.Where(goqu.C("spells.name").Eq(params.EqualsName))
	} else if len(params.LikeName) > 0 {
		request.Where(goqu.C("spells.name").ILike(params.LikeName))
	}
	if params.IsVerbal != tribool.Unset {
		request.Where(goqu.C("spells.is_verbal").Eq(params.IsVerbal == tribool.True))
	}
	if params.IsSomatic != tribool.Unset {
		request.Where(goqu.C("spells.is_somatic").Eq(params.IsSomatic == tribool.True))
	}
	if params.HasMaterialComponent != tribool.Unset {
		request.Where(goqu.C("spells.is_material").Eq(params.HasMaterialComponent == tribool.True))
	}
	if params.IsRitual != tribool.Unset {
		request.Where(goqu.C("spells.is_ritual").Eq(params.IsRitual == tribool.True))
	}
	if len(params.Levels) > 0 {
		levelsToRequest := getLevelsEnumeration(params.Levels)
		request.Where(goqu.C("spells.level").In(levelsToRequest))
	}
	if len(params.MagicalSchools) > 0 {
		schoolsToRequest := getSchoolsEnumeration(params.MagicalSchools)
		request.Where(goqu.C("spells.magical_school").In(schoolsToRequest))
	}
	request.Order(goqu.C("spells.name").Asc())
	request.Limit(uint(limit)).Offset(uint(offset))
	var spells []SpellDb
	sql, _, _ := request.ToSQL()
	err := rep.db.Select(&spells, sql)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Bad request: %s\n", err.Error())
		fmt.Println(sql)
		return []dto.SpellDto{}, err
	}
	ret := make([]dto.SpellDto, len(spells))
	for i := range ret {
		ret[i] = rep.dbSpellToSpellDto(spells[i])
	}
	return ret, nil
}

func (rep *SpellsRepository) GetById(id dto.SpellId) dto.SpellDto {
	return dto.SpellDto{}
}

func (rep *SpellsRepository) dbSpellToSpellDto(spellDb SpellDb) dto.SpellDto {
	res := dto.SpellDto{
		Id:                   dto.SpellId(spellDb.Id),
		Name:                 spellDb.Name,
		Level:                spellDb.Level,
		Description:          spellDb.Description,
		CastingTime:          spellDb.CastingTime,
		Duration:             spellDb.Duration,
		IsVerbal:             spellDb.IsVerbal,
		IsSomatic:            spellDb.IsSomatic,
		HasMaterialComponent: spellDb.HasMaterial,
		MagicalSchool:        spellDb.MagicalSchool,
		Distance:             spellDb.Distance,
		IsRitual:             spellDb.IsRitual,
		SourceId:             dto.SourceId{},
		SourceName:           spellDb.SourceName,
	}
	if spellDb.MaterialContent.Valid {
		res.MaterialComponent = spellDb.MaterialContent.String
	}
	return res
}

// TODO refactor next methods
func getSourcesEnumeration(ids []dto.SourceId) string {
	sources := strings.Builder{}
	for index, source := range ids {
		sources.WriteRune('\'')
		sources.WriteString(uuid.UUID(source).String())
		sources.WriteRune('\'')
		if index+1 < len(ids) {
			sources.WriteRune(',')
		}
	}
	return sources.String()
}

func getLevelsEnumeration(levels []int) string {
	levelsStr := strings.Builder{}
	for index, level := range levels {
		levelsStr.WriteRune('\'')
		levelsStr.WriteString(strconv.Itoa(level))
		levelsStr.WriteRune('\'')
		if index+1 < len(levels) {
			levelsStr.WriteRune(',')
		}
	}
	return levelsStr.String()
}

func getSchoolsEnumeration(schools []string) string {
	schoolsStr := strings.Builder{}
	for index, school := range schools {
		schoolsStr.WriteRune('\'')
		schoolsStr.WriteString(school)
		schoolsStr.WriteRune('\'')
		if index+1 < len(schools) {
			schoolsStr.WriteRune(',')
		}
	}
	return schoolsStr.String()
}

type SpellDb struct {
	Id              uuid.UUID      `db:"id"`
	Name            string         `db:"spells_name"`
	Level           int            `db:"level"`
	Description     string         `db:"description"`
	CastingTime     string         `db:"casting_time"`
	MagicalSchool   string         `db:"magical_school"`
	Duration        string         `db:"duration"`
	IsVerbal        bool           `db:"is_verbal"`
	IsSomatic       bool           `db:"is_somatic"`
	HasMaterial     bool           `db:"is_material"`
	MaterialContent sql.NullString `db:"material_content"`
	IsRitual        bool           `db:"is_ritual"`
	Distance        string         `db:"distance"`
	SourceId        uuid.UUID      `db:"source_id"`
	SourceName      string         `db:"sources_name"`
}
