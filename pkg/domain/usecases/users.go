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

func (useCase *UserUseCase) Register(innerDto dto.UserCreateDto) (dto.UserDto, error) {
	if !useCase.validatePasswordWithRules(innerDto.Password) {
		return dto.UserDto{}, errors.New("bad password")
	}
	if !useCase.validateLoginWithRules(innerDto.Login) {
		return dto.UserDto{}, errors.New("bad login")
	}
	encryptedPassword, err := hash.HashPassword(innerDto.Password)
	if err != nil {
		return dto.UserDto{}, err
	}
	userId := dto.UserId(uuid.New())
	dataToWrite := dto.UserToRepositoryDto{
		Id:             userId,
		Login:          innerDto.Login,
		HashedPassword: encryptedPassword,
	}
	useCase.repository.Users.CreateUser(dataToWrite)
	ret := dto.UserDto{
		Id:    userId,
		Login: innerDto.Login,
	}
	return ret, nil
}

func (useCase *UserUseCase) Find(innerDto dto.SearchUserDto) ([]dto.UserDto, error) {
	return useCase.repository.Users.GetUsers(innerDto)
}

func (useCase *UserUseCase) validatePasswordWithRules(password string) bool {
	return len(password) != 0
}

func (useCase *UserUseCase) validateLoginWithRules(login string) bool {
	return useCase.isLoginUnique(login) && newLoginChecker(login).basicCheck()
}

func (useCase *UserUseCase) isLoginUnique(login string) bool {
	found, _ := useCase.Find(dto.SearchUserDto{
		EqualsLogin: login,
	})
	for _, foundUser := range found {
		if foundUser.Login == login {
			return false
		}
	}
	return true
}

func (useCase *UserUseCase) isLoginSimilarWithAdmin(login string) bool {
	ADMIN_LIKE_LOGIN := "tvblackman"
	ADMIN_LOGIN := "tvblackman1"
	isEmpty := len(login) == 0
	isAdminLike := strings.Contains(login, ADMIN_LIKE_LOGIN) && login != ADMIN_LOGIN
	return !isEmpty && !isAdminLike
}

type loginChecker struct {
	login string
}

func newLoginChecker(login string) *loginChecker {
	return &loginChecker{login}
}

func (checker *loginChecker) basicCheck() bool {
	return !checker.isSimilarWithAdmin() && !checker.isEmpty()
}

func (checker *loginChecker) isSimilarWithAdmin() bool {
	ADMIN_LIKE_LOGIN := "tvblackman"
	ADMIN_LOGIN := "tvblackman1"
	isAdminLike := strings.Contains(checker.login, ADMIN_LIKE_LOGIN) && checker.login != ADMIN_LOGIN
	return isAdminLike
}

func (checker *loginChecker) isEmpty() bool {
	return len(checker.login) == 0
}
