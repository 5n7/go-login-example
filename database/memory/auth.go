package memory

import (
	"errors"

	"github.com/km2/go-login-example/model"
	"github.com/km2/go-login-example/repository"
)

var ErrAuthNotFound = errors.New("auth not found")

type AuthMemoryDB struct {
	auths []*model.Auth
}

func NewAuthMemoryDB() repository.AuthRepository {
	// for debug
	defaultAuths := []*model.Auth{
		{Email: "user@example.com", Password: "password"},
	}

	return &AuthMemoryDB{auths: defaultAuths}
}

func (db *AuthMemoryDB) GetAuthByEmail(email string) (*model.Auth, error) {
	for _, auth := range db.auths {
		if auth.Email == email {
			return auth, nil
		}
	}

	return nil, ErrAuthNotFound
}
