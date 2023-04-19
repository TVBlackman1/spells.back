package repository

import (
	"fmt"
	"spells.tvblackman1.ru/pkg/domain/dto"
)

type UsersRepository struct {
}

func NewUserRepository() *UsersRepository {
	return &UsersRepository{}
}

func (rep *UsersRepository) CreateUser(dto dto.UserToRepositoryDto) {
	fmt.Printf("%+v", dto)
}
