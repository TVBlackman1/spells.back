package main

import (
	"github.com/google/uuid"
	"log"
	"spells.tvblackman1.ru/lib/config"
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
	err = useCases.Spell.CreateSpell(dto.UserId(uuid.New()), dto.CreateSpellDto{
		Name:                 "first",
		Level:                0,
		Classes:              []string{"Warrior", "Druid"},
		Description:          "About this",
		CastingTime:          "1 action",
		Duration:             "30 sec",
		IsVerbal:             false,
		IsSomatic:            true,
		HasMaterialComponent: false,
		MaterialComponent:    "asd",
		MagicalSchool:        "illusion",
		Distance:             "30 feet",
		IsRitual:             true,
		SourceId:             dto.SourceId{},
	})
	if err != nil {
		log.Fatal(err)
	}

}
