package dto

import "github.com/google/uuid"

type UserId uuid.UUID

type UserDto struct {
	Id    UserId
	Login string
	Email string
}

type UserCreateDto struct {
	Login    string
	Password string
}

type UserToRepositoryDto struct {
	Id             UserId
	Login          string
	HashedPassword string
	Email          string
}

type SearchUserDto struct {
	Id    UserId
	Login string
}
