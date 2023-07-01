package dto

import "github.com/google/uuid"

type UrlSetId uuid.UUID

type UrlSetDto struct {
	Id   UrlSetId
	Uri  string
	Name string
}

type UrlSetToRepositoryDto struct {
	Id  UrlSetId
	Uri string
}
