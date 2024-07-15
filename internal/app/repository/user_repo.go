package repository

import (
	"Gophermart/internal/app/domain/model"
	"Gophermart/internal/app/repository/entity"
	"context"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

//func (r *UserRepository) FindByLoginAndHash(
//	ctx context.Context,
//	login model.Username,
//	hash model.PasswordHash) (*model.User, error) {
//	usr := &model.User{}
//	err := r.db.QueryRowxContext(ctx, "SELECT login, password as hash FROM users WHERE login = $1 AND password = $2", login, hash).StructScan(usr)
//	if err != nil {
//		return nil, err
//	}
//
//	return usr, nil
//}

func (r *UserRepository) FindByLogin(
	ctx context.Context,
	login model.Username) (*entity.User, error) {
	usr := &entity.User{}
	err := r.db.QueryRowxContext(ctx, "SELECT id as ID, login as UserName, password as Hash FROM users WHERE login = $1", login).StructScan(usr)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (r *UserRepository) AddNew(ctx context.Context, user *model.User) (int, error) {
	id := 0
	err := r.db.QueryRowxContext(ctx, "INSERT INTO users (login, password) VALUES ($1, $2) RETURNING id", user.Username, user.Hash).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}
