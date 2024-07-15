package model

import (
	"Gophermart/internal/app/repository/entity"
	"golang.org/x/crypto/bcrypt"
)

type Username string
type PasswordHash string
type User struct {
	Username Username
	Hash     PasswordHash
}

func newPasswordHash(password string) (PasswordHash, error) {
	//if len(password) < 8 || len(password) > 16 {
	//	return "", errors.New("password must be between 8 and 16")
	//}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return PasswordHash(hash), nil
}

func (u *User) CheckPassword(password []byte) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Hash), password)
}

func NewUserFromRealPass(username string, password string) (*User, error) {
	u := new(User)
	u.Username = Username(username)
	hash, err := newPasswordHash(password)
	if err != nil {
		return nil, err
	}
	u.Hash = hash

	return u, nil
}

func NewUserFromEntity(user *entity.User) *User {
	return &User{Username: Username(user.Username), Hash: PasswordHash(user.Hash)}

}
