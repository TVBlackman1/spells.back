package repository

import (
	"fmt"
	"github.com/google/uuid"
	"spells.tvblackman1.ru/pkg/domain/dto"
)

type UsersRepository struct {
}

func NewUserRepository() *UsersRepository {
	return &UsersRepository{}
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
