package repository

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"os"
	"spells.tvblackman1.ru/pkg/domain/dto"
)

type UsersRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UsersRepository {
	return &UsersRepository{db}
}

func (rep *UsersRepository) CreateUser(dto dto.UserToRepositoryDto) error {
	request := fmt.Sprintf("INSERT INTO %s(id, login, "+
		"hash_password) VALUES ('%s', '%s', '%s') RETURNING id;\n",
		UsersDbName,
		uuid.UUID(dto.Id).String(),
		dto.Login,
		dto.HashedPassword,
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
		fmt.Fprintf(os.Stderr, "Bad request: %s\n", err.Error())
		return dto.UserDto{}, err
	}
	return rep.dbUserToUserDto(user), nil
}

func (rep *UsersRepository) GetUsers(params dto.SearchUserDto) ([]dto.UserDto, error) {
	request := fmt.Sprintf("select id, login, email from %s;",
		UsersDbName,
	)
	var users []UserDb
	err := rep.db.Select(&users, request)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Bad request: %s\n", err.Error())
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
