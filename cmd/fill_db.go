package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"spells.tvblackman1.ru/lib/config"
	importer_v2 "spells.tvblackman1.ru/lib/importer.v2"
	"spells.tvblackman1.ru/lib/pagination"
	"spells.tvblackman1.ru/lib/postgres"
	"spells.tvblackman1.ru/pkg/domain/dto"
	"spells.tvblackman1.ru/pkg/domain/usecases"
	"spells.tvblackman1.ru/pkg/repository"
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
	user, err := setUser(useCases)
	if err != nil {
		fmt.Println(err)
	}
	_, meta, err := repo.Spells.GetSpells(dto.SearchSpellDto{}, pagination.Pagination{
		Limit:      1,
		PageNumber: 0,
	})
	if err != nil || meta.All > 0 {
		os.Exit(0)
	}

	directoryPath := path.Dir("./init/seeds/spell.categories/")
	files, err := os.ReadDir(directoryPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fullFilename := path.Join(directoryPath, file.Name())
		loadFileDataToDb(user.Id, useCases, fullFilename)
	}
}

func loadFileDataToDb(userId dto.UserId, useCases *usecases.UseCases, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	data := importer_v2.GetSpellsData(file)
	uploadSourcesToDbv2(userId, useCases, data)
	uploadSpellsToDbv2(userId, useCases, data)
	println(data.ShortName)
	return nil
}

func uploadSourcesToDbv2(userId dto.UserId, useCases *usecases.UseCases, data *importer_v2.MainStructure) {
	err := useCases.Source.CreateSource(userId, dto.SourceCreateDto{
		Name:        data.ShortName,
		Description: data.Name,
		IsOfficial:  data.IsOfficial,
	})
	if err != nil {
		fmt.Println(err.Error())
	}
}

func uploadSpellsToDbv2(userId dto.UserId, useCases *usecases.UseCases, data *importer_v2.MainStructure) {
	fmt.Printf("Uploading from %s\n", data.ShortName)

	similarSources, _ := useCases.Source.GetSources(dto.UserId{}, dto.SearchSourceDto{
		Name: data.ShortName,
	})
	var sourceId dto.SourceId
	for _, source := range similarSources {
		if source.Name == data.ShortName {
			sourceId = source.Id
		}
	}
	for index, spell := range data.Spells {
		err := useCases.Spell.CreateSpell(userId, dto.CreateSpellDto{
			Name:                 spell.Name,
			Level:                spell.Level,
			Classes:              []string{},
			Description:          spell.Description,
			CastingTime:          spell.CastingTime,
			Duration:             spell.Duration,
			IsVerbal:             spell.IsVerbal,
			IsSomatic:            spell.IsSomatic,
			HasMaterialComponent: spell.HasMaterialComponent,
			MaterialComponent:    spell.Materials,
			MagicalSchool:        spell.School,
			Distance:             spell.Distance,
			IsRitual:             spell.IsRitual,
			SourceId:             sourceId,
		})
		if err != nil {
			log.Fatal(err)
		}
		if index%20 == 0 {
			fmt.Printf("%d/%d\n", index, len(data.Spells))
		}
	}
}

func setUser(useCases *usecases.UseCases) (dto.UserDto, error) {
	users, err := useCases.User.Find(dto.SearchUserDto{
		EqualsLogin: "tvblackman1",
	})
	if err != nil {
		return dto.UserDto{}, err
	}
	fmt.Println(len(users))
	for _, user := range users {
		fmt.Println(user)
		if user.Login == "tvblackman1" {
			return user, nil
		}
	}
	return useCases.User.Register(dto.UserCreateDto{
		Login:    "tvblackman1",
		Password: "120474ba",
	})
}
