package main

import (
	"log"
	"net/http"
	"spells.tvblackman1.ru/lib/config"
	"spells.tvblackman1.ru/lib/postgres"
	"spells.tvblackman1.ru/pkg/domain/usecases"
	"spells.tvblackman1.ru/pkg/handler"
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
	handler := handler.NewHandler(useCases)
	http.ListenAndServe(":8080", handler)

}
