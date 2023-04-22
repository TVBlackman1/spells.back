package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"spells.tvblackman1.ru/pkg/domain/dto"
)

type UsersRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UsersRepository {
	return &UsersRepository{db}
}

func (rep *UsersRepository) CreateUser(dto dto.UserToRepositoryDto) {
	fmt.Printf("created: %+v", dto)
}

func (rep *UsersRepository) GetById(id dto.UserId) dto.UserDto {
	return dto.UserDto{
		Id:    dto.UserId(uuid.New()),
		Login: "some-login",
		Email: "some-email",
	}
}
