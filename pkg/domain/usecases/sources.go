package usecases

import (
	"github.com/google/uuid"
	"spells.tvblackman1.ru/pkg/domain/boundaries"
	"spells.tvblackman1.ru/pkg/domain/dto"
)

type SourceUseCase struct {
	repository *boundaries.Repository
}

func NewSourceUseCase(repository *boundaries.Repository) *SourceUseCase {
	return &SourceUseCase{repository}
}

func (usecase *SourceUseCase) CreateSource(userId dto.UserId, sourceDto dto.SourceCreateDto) {
	user := usecase.repository.Users.GetById(userId)
	dataToWrite := dto.SourceToRepositoryDto{
		Id:          dto.SourceId(uuid.New()),
		Name:        sourceDto.Name,
		Description: sourceDto.Description,
		Version:     1,
		UploadedBy:  user.Id,
		IsOfficial:  sourceDto.IsOfficial,
		Author:      user.Login,
	}
	usecase.repository.Sources.CreateSource(dataToWrite)
}
