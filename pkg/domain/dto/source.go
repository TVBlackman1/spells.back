package dto

import "github.com/google/uuid"

type SourceId uuid.UUID

type SourceDto struct {
	Id          SourceId
	Name        string
	Description string
	Version     int
	Author      string
	UploadedBy  UserId
	IsOfficial  bool
}

type SourceCreateDto struct {
	Name        string
	Description string
	IsOfficial  bool
}

type SourceToRepositoryDto struct {
	Id          SourceId
	Name        string
	Description string
	Version     int
	UploadedBy  UserId
	IsOfficial  bool
	Author      string
}

type SearchSourceDto struct {
	Name                  string
	CurrentUser           UserId
	IsOfficial            bool
	UploadedByCurrentUser bool
	ExternLibs            bool
}
