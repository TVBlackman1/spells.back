package main

import (
	"fmt"
	"log"
	"spells.tvblackman1.ru/lib/config"
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
	//sourceId, _ := uuid.FromString("21ce153f-ca9f-4f69-b125-86a03468f8fa")
	spells, err := useCases.Spell.GetSpellList(dto.UserId{}, dto.SearchSpellDto{
		MagicalSchools: []string{"иллюзия", "преобразование"},
	}, pagination.Pagination{
		Limit:      10,
		PageNumber: 1,
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, spell := range spells {
		fmt.Println(spell.Name, spell.SourceName)
	}
}
