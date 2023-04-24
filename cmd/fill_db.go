package main

import (
	"fmt"
	"github.com/gofrs/uuid"
	"golang.org/x/exp/slices"
	"log"
	"os"
	"spells.tvblackman1.ru/lib/config"
	"spells.tvblackman1.ru/lib/importer"
	"spells.tvblackman1.ru/lib/postgres"
	"spells.tvblackman1.ru/pkg/domain/dto"
	"spells.tvblackman1.ru/pkg/domain/usecases"
	"spells.tvblackman1.ru/pkg/repository"
	"strings"
)

func main() {
	config, err := config.LoadConfig("./")
	if err != nil {
		log.Fatal(err)
	}
	dbConfig := postgres.PostgresConfig{
		Dbname:   config.POSTGRES_DBNAME,
		Host:     config.POSTGRES_HOST,
		Port:     config.POSTGRES_PORT,
		User:     config.POSTGRES_USER,
		Password: config.POSTGRES_PASS,
	}
	postgresDb, err := postgres.Connect(dbConfig)
	defer postgresDb.Close()

	if err != nil {
		log.Fatal(err)
	}
	repo, err := repository.NewRepository(postgresDb)
	if err != nil {
		log.Fatalf("repo err: %s", err.Error())
	}
	useCases := usecases.NewUseCases(repo)
	file, err := os.Open("./init/seeds/dnd-spells.json")
	if err != nil {
		log.Fatal(err)
	}
	data := importer.GetSpellsData(file)
	fmt.Println(len(data.AllSpells))
	//user, err := registerDefaultUser(useCases)
	//if err != nil {
	//	fmt.Println(err)
	//}
	uuidReady, _ := uuid.FromString("4366cc02-5150-4686-ba62-954f88a112cd")
	customUserId := dto.UserId(uuidReady)
	//uploadSourcesToDb(customUserId, useCases, data)
	//uploadSourcesToDb(user.Id, useCases, data)
	//aliasSourceToId := getSourceIds(useCases, data.SourceList)
	//for name, id := range aliasSourceToId {
	//	fmt.Println(name, id)
	//}
	uploadSpellsToDb(customUserId, useCases, data)

}

func uploadSourcesToDb(userId dto.UserId, useCases *usecases.UseCases, data *importer.MainStructure) {
	for sourceName, source := range data.SourceList {
		useCases.Source.CreateSource(userId, dto.SourceCreateDto{
			Name:        sourceName,
			Description: source.Text.Ru.Title,
			IsOfficial:  source.Official,
		})
	}
}

func uploadSpellsToDb(userId dto.UserId, useCases *usecases.UseCases, data *importer.MainStructure) {
	aliasSourceToId := getSourceIds(useCases, data.SourceList)
	for ind, spell := range data.AllSpells[0:498] {
		//if spell.En.Name != "Earth Tremor" {
		//	continue
		//}
		ruSpell := spell.Ru
		spellSourceNames := func() string {
			if len(spell.Ru.Source) == 0 {
				return spell.En.Source
			} else {
				return spell.Ru.Source
			}
		}()
		spellLevel, err := spell.En.Level.Int64()
		if err != nil {
			fmt.Printf("not valid spell level: %s of spell %s\n", ruSpell.Level.String(), ruSpell.Name)
			continue
		}
		sourceIds, err := aliasSourcesToIds(strings.Split(spellSourceNames, ", "), aliasSourceToId)
		if err != nil {
			fmt.Printf(err.Error())
			continue
		}
		err = useCases.Spell.CreateSpell(userId, dto.CreateSpellDto{
			Name:                 ruSpell.Name,
			Level:                int(spellLevel),
			Classes:              []string{},
			Description:          strings.Replace(ruSpell.Text, "'", "", -1),
			CastingTime:          ruSpell.CastingTime,
			Duration:             ruSpell.Duration,
			IsVerbal:             slices.Contains(strings.Split(ruSpell.Components, ", "), "В"),
			IsSomatic:            slices.Contains(strings.Split(ruSpell.Components, ", "), "С"),
			HasMaterialComponent: slices.Contains(strings.Split(ruSpell.Components, ", "), "М"),
			MaterialComponent:    ruSpell.Materials,
			MagicalSchool:        ruSpell.School,
			Distance:             ruSpell.Range,
			IsRitual:             len(ruSpell.Ritual) > 0,
			SourceIds:            sourceIds,
		})
		if err != nil {
			log.Fatal(err)
		}
		if ind%20 == 0 {
			fmt.Printf("%d/%d\n", ind, len(data.AllSpells))
		}
	}
}

func registerDefaultUser(useCases *usecases.UseCases) (dto.UserDto, error) {
	return useCases.User.Register(dto.UserCreateDto{
		Login:    "tvblackman1",
		Password: "120474ba",
	})
}

func getSourceIds(useCases *usecases.UseCases, sources map[string]importer.Source) map[string]dto.SourceId {
	aliases := make(map[string]dto.SourceId)
	for sourceName := range sources {
		similarSources, _ := useCases.Source.GetSources(dto.UserId{}, dto.SearchSourceDto{
			Name: sourceName,
		})
		for _, source := range similarSources {
			if source.Name == sourceName {
				aliases[source.Name] = source.Id
			}
		}
	}
	return aliases
}

func aliasSourcesToIds(sources []string, alias map[string]dto.SourceId) ([]dto.SourceId, error) {
	ret := make([]dto.SourceId, len(sources))
	for i, source := range sources {
		var ok bool
		ret[i], ok = alias[source]
		if !ok {
			return []dto.SourceId{}, fmt.Errorf("bad source name: %s", source)
		}

	}
	return ret, nil
}
