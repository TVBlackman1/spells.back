package usecases

import (
	"errors"
	"github.com/google/uuid"
	"spells.tvblackman1.ru/lib/hash"
	"spells.tvblackman1.ru/pkg/domain/boundaries"
	"spells.tvblackman1.ru/pkg/domain/dto"
	"strings"
)

type UserUseCase struct {
	repository *boundaries.Repository
}

func NewUserUseCase(repository *boundaries.Repository) *UserUseCase {
	return &UserUseCase{repository}
}

func (useCase *UserUseCase) Register(innerDto dto.UserCreateDto) error {
	if !useCase.validatePasswordWithRules(innerDto.Password) {
		return errors.New("bad password")
	}
	if !useCase.validateLoginWithRules(innerDto.Login) {
		return errors.New("bad login")
	}
	encryptedPassword, err := hash.HashPassword(innerDto.Password)
	if err != nil {
		return err
	}
	userId := dto.UserId(uuid.New())
	dataToWrite := dto.UserToRepositoryDto{
		Id:             userId,
		Login:          innerDto.Login,
		HashedPassword: encryptedPassword,
	}
	useCase.repository.Users.CreateUser(dataToWrite)
	return nil
}

func (useCase *UserUseCase) validatePasswordWithRules(password string) bool {
	return len(password) != 0
}

func (useCase *UserUseCase) validateLoginWithRules(login string) bool {
	ADMIN_LIKE_LOGIN := "tvblackman"
	ADMIN_LOGIN := "tvblackman1"
	isEmpty := len(login) == 0
	isAdminLike := strings.Contains(login, ADMIN_LIKE_LOGIN) && login != ADMIN_LOGIN
	return !isEmpty && !isAdminLike
}
