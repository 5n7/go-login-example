package repository

import "github.com/km2/go-login-example/model"

type AuthRepository interface {
	GetAuthByEmail(email string) (*model.Auth, error)
}
