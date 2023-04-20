package dto

import "github.com/google/uuid"

type SetId uuid.UUID

type SetDto struct {
	Id          SetId
	Name        string
	UserId      UserId
	Description string
	Sources     []string
}

type CreateSetDto struct {
	Name        string
	Description string
	Sources     []string
}

type SetToRepositoryDto struct {
	Id          SetId
	Name        string
	UserId      UserId
	Description string
	Sources     []string
}
