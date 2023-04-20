package repository

import (
	"fmt"
	"spells.tvblackman1.ru/pkg/domain/dto"
)

type SourcesRepository struct {
}

func NewSourcesRepository() *SourcesRepository {
	return &SourcesRepository{}
}

func (rep *SourcesRepository) CreateSource(sourceDto dto.SourceToRepositoryDto) {
	fmt.Printf("created: %+v", sourceDto)
}

func (rep *SourcesRepository) GetById(id dto.SourceId) dto.SourceDto {
	return dto.SourceDto{}
}
