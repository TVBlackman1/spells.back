package usecases

import (
	"errors"
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

func (usecase *SourceUseCase) CreateSource(userId dto.UserId, sourceDto dto.SourceCreateDto) error {
	user, _ := usecase.repository.Users.GetById(userId)
	if !usecase.isNewNameOfSource(sourceDto.Name, userId) {
		return errors.New("not unique name of spell in source")
	}
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
	return nil
}

func (usecase *SourceUseCase) GetSources(userId dto.UserId, params dto.SearchSourceDto) ([]dto.SourceDto, error) {
	return usecase.repository.Sources.GetSources(userId, params)
}

func (usecase *SourceUseCase) isNewNameOfSource(sourceName string, userId dto.UserId) bool {
	sources, _ := usecase.repository.Sources.GetSources(userId, dto.SearchSourceDto{
		Name: sourceName,
	})
	for _, source := range sources {
		if source.Name == sourceName {
			return false
		}
	}
	return true

}
