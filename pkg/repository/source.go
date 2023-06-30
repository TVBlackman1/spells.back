package repository

import (
	"database/sql"
	"fmt"
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
	request := fmt.Sprintf("INSERT INTO %s(id, name, "+
		"description, is_official, author, uploaded_by) VALUES ('%s', '%s', '%s', '%t', '%s', '%s') RETURNING id;\n",
		SourcesDbName,
		uuid.UUID(sourceDto.Id).String(),
		sourceDto.Name,
		sourceDto.Description,
		sourceDto.IsOfficial,
		sourceDto.Author,
		uuid.UUID(sourceDto.UploadedBy).String(),
	)
	var uuidStr string
	err := rep.db.Get(&uuidStr, request)
	if err != nil {
		fmt.Println(request)
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (rep *SourcesRepository) GetById(id dto.SourceId) (dto.SourceDto, error) {
	request := fmt.Sprintf("select id, name, description, version_number, is_official, author, uploaded_by from %s where id='%s';\n",
		SourcesDbName, uuid.UUID(id).String(),
	)
	var source SourceDb
	err := rep.db.Get(&source, request)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Bad requests: %s. While getting source by id: %s\n", err.Error(), uuid.UUID(id).String())
		return dto.SourceDto{}, err
	}
	return rep.dbSourceToSourceDto(source), nil
}

func (rep *SourcesRepository) GetSources(userId dto.UserId, params dto.SearchSourceDto) ([]dto.SourceDto, error) {
	request := fmt.Sprintf("select id, name, description, version_number, is_official, author, uploaded_by from %s;\n",
		SourcesDbName,
	)
	var sources []SourceDb
	err := rep.db.Select(&sources, request)
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
