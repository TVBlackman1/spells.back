package repository

import (
	"database/sql"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"os"
	"spells.tvblackman1.ru/pkg/domain/dto"
)

type SourcesRepository struct {
	db *sqlx.DB
}

func NewSourcesRepository(db *sqlx.DB) *SourcesRepository {
	return &SourcesRepository{db}
}

func (rep *SourcesRepository) CreateSource(sourceDto dto.SourceToRepositoryDto) error {
	dialect := goqu.Dialect("postgres")
	request := dialect.Insert(SourcesDbName).Rows(
		goqu.Record{
			"id":          uuid.UUID(sourceDto.Id).String(),
			"name":        sourceDto.Name,
			"description": sourceDto.Description,
			"is_official": sourceDto.IsOfficial,
			"author":      sourceDto.Author,
			"uploaded_by": uuid.UUID(sourceDto.UploadedBy).String(),
		}).
		Returning("id")
	sql, _, _ := request.ToSQL()
	var uuidStr string
	err := rep.db.Get(&uuidStr, sql)
	if err != nil {
		fmt.Println(sql)
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (rep *SourcesRepository) GetById(id dto.SourceId) (dto.SourceDto, error) {
	idStringifier := uuid.UUID(id).String()
	dialect := goqu.Dialect("postgres")
	request := dialect.
		Select("id", "name", "description", "is_official", "author", "uploaded_by").
		From(SourcesDbName).
		Where(goqu.C("id").Eq(idStringifier))
	sql, _, _ := request.ToSQL()
	var source SourceDb
	err := rep.db.Get(&source, sql)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Bad requests: %s. While getting source by id: %s\n", err.Error(), idStringifier)
		return dto.SourceDto{}, err
	}
	return rep.dbSourceToSourceDto(source), nil
}

func (rep *SourcesRepository) GetSources(userId dto.UserId, params dto.SearchSourceDto) ([]dto.SourceDto, error) {
	dialect := goqu.Dialect("postgres")
	request := dialect.
		Select("id", "name", "description", "is_official", "author", "uploaded_by").
		From(SourcesDbName)
	sql, _, _ := request.ToSQL()
	var sources []SourceDb
	err := rep.db.Select(&sources, sql)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Bad requests: %s\n", err.Error())
		return []dto.SourceDto{}, err
	}
	ret := make([]dto.SourceDto, len(sources))
	for i := range ret {
		ret[i] = rep.dbSourceToSourceDto(sources[i])
	}
	return ret, nil
}

func (rep *SourcesRepository) dbSourceToSourceDto(sourceDb SourceDb) dto.SourceDto {
	ret := dto.SourceDto{
		Id:         dto.SourceId(sourceDb.Id),
		Name:       sourceDb.Name,
		Version:    sourceDb.Version,
		Author:     sourceDb.Author,
		UploadedBy: dto.UserId(sourceDb.UploadedBy),
		IsOfficial: sourceDb.IsOfficial,
	}
	if sourceDb.Description.Valid {
		ret.Description = sourceDb.Description.String
	}
	return ret
}

type SourceDb struct {
	Id          uuid.UUID      `db:"id"`
	Name        string         `db:"name"`
	Version     int            `db:"version_number"`
	Description sql.NullString `db:"description"`
	IsOfficial  bool           `db:"is_official"`
	Author      string         `db:"author"`
	UploadedBy  uuid.UUID      `db:"uploaded_by"`
}
