package usecases

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"spells.tvblackman1.ru/pkg/domain/boundaries"
	"spells.tvblackman1.ru/pkg/domain/dto"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type UrlSetUseCase struct {
	repository *boundaries.Repository
}

func NewUrlSetUseCase(repository *boundaries.Repository) *UrlSetUseCase {
	return &UrlSetUseCase{repository}
}

func (usecase *UrlSetUseCase) CreateUrlSet() (string, error) {
	linkPart := usecase.generateRandomLinkPart()
	link := usecase.linkFromLinkPart(linkPart)
	err := usecase.repository.UrlSets.CreateUrlSet(dto.UrlSetToRepositoryDto{
		Id:  dto.UrlSetId(uuid.New()),
		Uri: link,
	})
	if err != nil {
		return "", err
	}
	return link, nil
}

func (usecase *UrlSetUseCase) RenameUrlSet(linkPart string, newName string) error {
	link := usecase.linkFromLinkPart(linkPart)
	urlSet, err := usecase.repository.UrlSets.GetByLink(link)
	if err != nil {
		return err
	}
	return usecase.repository.UrlSets.RenameUrlSet(urlSet.Id, newName)
}

func (usecase *UrlSetUseCase) GetUrlSet(linkPart string) (dto.UrlSetDto, error) {
	link := usecase.linkFromLinkPart(linkPart)
	return usecase.repository.UrlSets.GetByLink(link)
}

func (usecase *UrlSetUseCase) generateRandomLinkPart() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	builder := strings.Builder{}
	for i := 0; i < 8; i++ {
		letterIndex := int(rand.Uint32()) % len(letters)
		letter := letters[letterIndex]
		builder.WriteRune(letter)
	}
	return builder.String()
}

func (usecase *UrlSetUseCase) linkFromLinkPart(linkPart string) string {
	return fmt.Sprintf("http://localhost:8080/api/v1/url-sets/%s", linkPart)
}
