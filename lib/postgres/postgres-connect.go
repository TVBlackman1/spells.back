package postgres

import (
	"errors"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type PostgresConfig struct {
	Dbname   string
	Host     string
	Port     int
	User     string
	Password string
}

func Connect(config PostgresConfig) (*sqlx.DB, error) {
	dbConnectUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		config.User, config.Password, config.Host, config.Port, config.Dbname,
	)
	conn, err := sqlx.Connect("pgx", dbConnectUrl)
	if err != nil {
		errorText := fmt.Sprintf("Unable to connect to database: %v\n", err)
		return nil, errors.New(errorText)
	}
	fmt.Printf("Connected to database '%s' on %s:%d\n", config.Dbname, config.Host, config.Port)
	return conn, nil
}
