package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"spells.tvblackman1.ru/pkg/domain/dto"
)

type SourcesRepository struct {
	db *sqlx.DB
}

func NewSourcesRepository(db *sqlx.DB) *SourcesRepository {
	return &SourcesRepository{db}
}

func (rep *SourcesRepository) CreateSource(sourceDto dto.SourceToRepositoryDto) {
	fmt.Printf("created: %+v", sourceDto)
}

func (rep *SourcesRepository) GetById(id dto.SourceId) dto.SourceDto {
	return dto.SourceDto{}
}
