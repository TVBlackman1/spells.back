package repository

import (
	"database/sql"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
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

func (rep *UsersRepository) CreateUser(createDto dto.UserToRepositoryDto) error {
	dialect := goqu.Dialect("postgres")
	request := dialect.Insert(UsersDbName).
		Rows(goqu.Record{
			"id":            uuid.UUID(createDto.Id).String(),
			"login":         createDto.Login,
			"hash_password": createDto.HashedPassword,
		}).
		Returning("id")
	sql, _, _ := request.ToSQL()
	var uuidStr string
	err := rep.db.Get(&uuidStr, sql)
	if err != nil {
		fmt.Println(sql)
		fmt.Println(err)
		return err
	}
	return nil
}

func (rep *UsersRepository) GetById(id dto.UserId) (dto.UserDto, error) {
	idStringifier := uuid.UUID(id).String()
	dialect := goqu.Dialect("postgres")
	request := dialect.
		Select("id", "login", "email").
		From(UsersDbName).
		Where(goqu.C("id").Eq(idStringifier)).
		Limit(1)
	sql, _, _ := request.ToSQL()
	var user UserDb
	err := rep.db.Get(&user, sql)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Bad requests: %s\n", err.Error())
		return dto.UserDto{}, err
	}
	return rep.dbUserToUserDto(user), nil
}

func (rep *UsersRepository) GetUsers(params dto.SearchUserDto) ([]dto.UserDto, error) {
	dialect := goqu.Dialect("postgres")
	request := dialect.Select("id", "login", "email").
		From(UsersDbName)
	if uuid.UUID(params.Id) != uuid.Nil {
		idStringifier := uuid.UUID(params.Id).String()
		request.Where(goqu.C("id").Eq(idStringifier))
	} else if len(params.EqualsLogin) > 0 {
		request.Where(goqu.C("login").Eq(params.EqualsLogin))
	} else if len(params.LikeLogin) > 0 {
		request.Where(goqu.C("login").ILike(params.LikeLogin))
	}
	var users []UserDb
	sql, _, _ := request.ToSQL()
	err := rep.db.Select(&users, sql)
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
