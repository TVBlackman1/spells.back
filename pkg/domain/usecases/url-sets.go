package usecases

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/microcosm-cc/bluemonday"
	"math/rand"
	"spells.tvblackman1.ru/lib/pagination"
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
	usecases   *UseCases
}

func NewUrlSetUseCase(repository *boundaries.Repository, usecases *UseCases) *UrlSetUseCase {
	return &UrlSetUseCase{repository, usecases}
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

func (usecase *UrlSetUseCase) AddSpell(linkPart string, spellId dto.SpellId) error {
	// TODO if spell exists, if link exists
	link := usecase.linkFromLinkPart(linkPart)
	urlSet, err := usecase.repository.UrlSets.GetByLink(link)
	if err != nil {
		return err
	}
	_, meta, err := usecase.repository.UrlSets.GetSpells(urlSet.Id, dto.SearchSpellDto{
		Id: spellId,
	}, pagination.Pagination{
		Limit: 1,
	})
	if err != nil {
		fmt.Printf("can not get spell with id %s\n", uuid.UUID(spellId).String())
		return err
	}
	spellAlreadyExists := meta.AllRecords != 0
	if spellAlreadyExists {
		return errors.New("already exists")
	}
	return usecase.repository.UrlSets.AddSpell(urlSet.Id, spellId)
}

func (usecase *UrlSetUseCase) RemoveSpell(linkPart string, spellId dto.SpellId) error {
	// TODO if spell exists, if link exists
	link := usecase.linkFromLinkPart(linkPart)
	urlSet, err := usecase.repository.UrlSets.GetByLink(link)
	if err != nil {
		return err
	}
	return usecase.repository.UrlSets.RemoveSpell(urlSet.Id, spellId)
}

func (usecase *UrlSetUseCase) GetSpells(linkPart string, search dto.SearchSpellDto, pag pagination.Pagination) ([]dto.SpellDto, pagination.Meta, error) {
	link := usecase.linkFromLinkPart(linkPart)
	urlSet, err := usecase.repository.UrlSets.GetByLink(link)
	if err != nil {
		return []dto.SpellDto{}, pagination.Meta{}, err
	}
	spells, meta, err := usecase.repository.UrlSets.GetSpells(urlSet.Id, search, pag)
	p := bluemonday.StripTagsPolicy()
	for index := range spells {
		spells[index].Description = p.Sanitize(spells[index].Description)
	}
	return spells, meta, err
}

func (usecase *UrlSetUseCase) GetAllSpells(linkPart string, search dto.SearchSpellDto, pag pagination.Pagination) ([]dto.SpellMarkedDto, pagination.Meta, error) {
	link := usecase.linkFromLinkPart(linkPart)
	urlSet, err := usecase.repository.UrlSets.GetByLink(link)
	if err != nil {
		return []dto.SpellMarkedDto{}, pagination.Meta{}, err
	}
	return usecase.repository.UrlSets.GetAllSpells(urlSet.Id, search, pag)
}

func (usecase *UrlSetUseCase) CreateUrlSetWithSpells(urlSetName string, spellNames []string) (string, error) {
	link, err := usecase.CreateUrlSet()
	if err != nil {
		fmt.Println("Can not create url set")
		return "", err
	}
	usecase.RenameUrlSet(usecase.linkPartFromLink(link), urlSetName)
	for _, name := range spellNames {
		list, meta, err := usecase.usecases.Spell.GetCommonSpellList(dto.SearchSpellDto{
			EqualsName: name,
		}, pagination.Pagination{
			Limit: 1,
		})
		if err != nil {
			return "", err
		}
		if meta.AllRecords > 0 {
			fmt.Println(list[0].Name, list[0].Id)
			usecase.AddSpell(
				usecase.linkPartFromLink(link),
				list[0].Id)
		} else {
			fmt.Printf("Not found: %s\n", name)
		}
	}
	return link, nil
	//link := usecase.linkFromLinkPart(linkPart)
	//urlSet, err := usecase.repository.UrlSets.GetByLink(link)
	//if err != nil {
	//	return []dto.SpellMarkedDto{}, pagination.Meta{}, err
	//}
	//return usecase.repository.UrlSets.GetAllSpells(urlSet.Id, search, pag)
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

func (usecase *UrlSetUseCase) linkPartFromLink(link string) string {
	slashIndex := strings.LastIndex(link, "/")
	return link[slashIndex+1:]
}
