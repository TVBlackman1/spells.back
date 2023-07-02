package dbdto

import "github.com/google/uuid"

type UrlSetDb struct {
	Id   uuid.UUID `db:"id"`
	Name string    `db:"name"`
	Uri  string    `db:"url"`
}
