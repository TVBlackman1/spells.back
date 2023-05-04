package main

import (
	"fmt"
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
	user, err := setUser(useCases)
	if err != nil {
		fmt.Println(err)
	}
	uploadSourcesToDb(user.Id, useCases, data)
	uploadSpellsToDb(user.Id, useCases, data)

}

func uploadSourcesToDb(userId dto.UserId, useCases *usecases.UseCases, data *importer.MainStructure) {
	for sourceName, source := range data.SourceList {
		err := useCases.Source.CreateSource(userId, dto.SourceCreateDto{
			Name:        sourceName,
			Description: source.Text.Ru.Title,
			IsOfficial:  source.Official,
		})
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func uploadSpellsToDb(userId dto.UserId, useCases *usecases.UseCases, data *importer.MainStructure) {
	aliasSourceToId := getSourceIds(useCases, data.SourceList)
	for ind, spell := range data.AllSpells {
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
		//if ruSpell.Name != "Святилище" {
		//	continue
		//}
		hasMaterialComponent := hasComponent(spell.En.Components, 'M')
		isVerbal := hasComponent(spell.En.Components, 'V')
		isSomatic := hasComponent(spell.En.Components, 'S')
		sourceIds, err := aliasSourcesToIds(strings.Split(spellSourceNames, ", "), aliasSourceToId)
		if err != nil {
			fmt.Printf(err.Error(), "\n")
			continue
		}
		for _, sourceId := range sourceIds {
			err = useCases.Spell.CreateSpell(userId, dto.CreateSpellDto{
				Name:                 ruSpell.Name,
				Level:                int(spellLevel),
				Classes:              []string{},
				Description:          strings.Replace(ruSpell.Text, "'", "", -1),
				CastingTime:          ruSpell.CastingTime,
				Duration:             ruSpell.Duration,
				IsVerbal:             isVerbal,
				IsSomatic:            isSomatic,
				HasMaterialComponent: hasMaterialComponent,
				MaterialComponent:    ruSpell.Materials,
				MagicalSchool:        ruSpell.School,
				Distance:             ruSpell.Range,
				IsRitual:             len(ruSpell.Ritual) > 0,
				SourceId:             sourceId,
			})
			if err != nil {
				log.Fatal(err)
			}
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

func setUser(useCases *usecases.UseCases) (dto.UserDto, error) {
	users, err := useCases.User.Find(dto.SearchUserDto{
		Login: "tvblackman1",
	})
	if err != nil {
		return dto.UserDto{}, err
	}
	for _, user := range users {
		if user.Login == "tvblackman1" {
			return user, nil
		}
	}
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

func hasComponent(components string, component rune) bool {
	return strings.IndexRune(components, component) != -1
}
