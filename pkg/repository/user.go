package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"os"
	"spells.tvblackman1.ru/lib/requests"
	"spells.tvblackman1.ru/pkg/domain/dto"
)

type UsersRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UsersRepository {
	return &UsersRepository{db}
}

func (rep *UsersRepository) CreateUser(createDto dto.UserToRepositoryDto) error {
	users, _ := rep.GetUsers(dto.SearchUserDto{
		EqualsLogin: createDto.Login,
	})
	if len(users) > 0 {
		return errors.New("user already exists")
	}
	request := fmt.Sprintf("INSERT INTO %s(id, login, "+
		"hash_password) VALUES ('%s', '%s', '%s') RETURNING id;\n",
		UsersDbName,
		uuid.UUID(createDto.Id).String(),
		createDto.Login,
		createDto.HashedPassword,
	)
	var uuidStr string
	err := rep.db.Get(&uuidStr, request)
	if err != nil {
		fmt.Println(request)
		fmt.Println(err)
		return err
	}
	return nil
}

func (rep *UsersRepository) GetById(id dto.UserId) (dto.UserDto, error) {
	request := fmt.Sprintf("select id, login, email from %s where id='%s' limit 1;",
		UsersDbName,
		uuid.UUID(id).String(),
	)
	var user UserDb
	err := rep.db.Get(&user, request)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Bad requests: %s\n", err.Error())
		return dto.UserDto{}, err
	}
	return rep.dbUserToUserDto(user), nil
}

func (rep *UsersRepository) GetUsers(params dto.SearchUserDto) ([]dto.UserDto, error) {
	request := requests.NewRequest(UsersDbName)
	request.Select("id, login, email")
	if uuid.UUID(params.Id) != uuid.Nil {
		request.Where(fmt.Sprintf("id='%s'", uuid.UUID(params.Id).String()))
	} else if len(params.EqualsLogin) > 0 {
		request.Where(fmt.Sprintf("login='%s'", params.EqualsLogin))
	} else if len(params.LikeLogin) > 0 {
		request.Where(fmt.Sprintf("login like '%%%s%%'", params.LikeLogin))
	}
	var users []UserDb
	fmt.Println(request.String())
	err := rep.db.Select(&users, request.String())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Bad requests: %s\n", err.Error())
		return []dto.UserDto{}, err
	}
	ret := make([]dto.UserDto, len(users))
	for i := range ret {
		ret[i] = rep.dbUserToUserDto(users[i])
	}
	return ret, nil
}

func (rep *UsersRepository) dbUserToUserDto(userDb UserDb) dto.UserDto {
	ret := dto.UserDto{
		Id:    dto.UserId(userDb.Id),
		Login: userDb.Login,
	}
	if userDb.Email.Valid {
		ret.Email = userDb.Email.String
	}
	return ret
}

type UserDb struct {
	Id    uuid.UUID      `db:"id"`
	Login string         `db:"login"`
	Email sql.NullString `db:"email"`
}
