package repository

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"os"
	"spells.tvblackman1.ru/lib/pagination"
	"spells.tvblackman1.ru/lib/requests"
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

	request := fmt.Sprintf("INSERT INTO %s(id, name, level, "+
		"description, classes, casting_time, duration, is_verbal, is_somatic, is_material, material_content, magical_school, "+
		"distance, is_ritual, source_id) VALUES ('%s', '%s', '%d', '%s', '%s', '%s', '%s', '%t', '%t', '%t', '%s', "+
		"'%s', '%s', '%t', '%s') RETURNING id;\n",
		SpellsDbName,
		uuid.UUID(spellDto.Id).String(),
		spellDto.Name,
		spellDto.Level,
		spellDto.Description,
		strings.Join(spellDto.Classes, ", "),
		spellDto.CastingTime,
		spellDto.Duration,
		spellDto.IsVerbal,
		spellDto.IsSomatic,
		spellDto.HasMaterialComponent,
		spellDto.MaterialComponent,
		spellDto.MagicalSchool,
		spellDto.Distance,
		spellDto.IsRitual,
		uuid.UUID(spellDto.SourceId).String(),
	)
	var uuidStr string
	err := rep.db.Get(&uuidStr, request)
	if err != nil {
		fmt.Println(request)
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

	request := requests.NewRequest(SpellsDbName)
	request.Select("spells.id, spells.name as \"spells.name\", spells.level, spells.description, spells.casting_time")
	request.Select("spells.duration, spells.is_verbal, spells.is_somatic, spells.is_material")
	request.Select("spells.material_content, spells.magical_school, spells.distance, spells.is_ritual")
	request.Select("spells.source_id")

	request.LeftJoin(SourcesDbName, "spells.source_id=sources.id")
	request.Select("sources.name as \"sources.name\"")
	if len(params.Sources) > 0 {
		sourcesToRequest := getSourcesEnumeration(params.Sources)
		request.Where(fmt.Sprintf("sources.id in (%s)", sourcesToRequest))
	}
	if len(params.EqualsName) > 0 {
		request.Where(fmt.Sprintf("spells.name='%s'", params.EqualsName))
	} else if len(params.LikeName) > 0 {
		request.Where(fmt.Sprintf("spells.name ilike '%%%s%%'", params.LikeName))
	}
	if params.IsVerbal != tribool.Unset {
		request.Where(fmt.Sprintf("spells.is_verbal=%t", params.IsVerbal == tribool.True))
	}
	if params.IsSomatic != tribool.Unset {
		request.Where(fmt.Sprintf("spells.is_somatic=%t", params.IsSomatic == tribool.True))
	}
	if params.HasMaterialComponent != tribool.Unset {
		request.Where(fmt.Sprintf("spells.is_material=%t", params.HasMaterialComponent == tribool.True))
	}
	if params.IsRitual != tribool.Unset {
		request.Where(fmt.Sprintf("spells.is_ritual=%t", params.IsRitual == tribool.True))
	}
	if len(params.Levels) > 0 {
		levelsToRequest := getLevelsEnumeration(params.Levels)
		request.Where(fmt.Sprintf("spells.level in (%s)", levelsToRequest))
	}
	if len(params.MagicalSchools) > 0 {
		schoolsToRequest := getSchoolsEnumeration(params.MagicalSchools)
		request.Where(fmt.Sprintf("spells.magical_school in (%s)", schoolsToRequest))
	}
	request.OrderBy("spells.name")
	request.Limit(limit).Offset(offset)
	var spells []SpellDb
	err := rep.db.Select(&spells, request.String())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Bad request: %s\n", err.Error())
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
	Name            string         `db:"spells.name"`
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
	SourceName      string         `db:"sources.name"`
}
