package db

import (
	"bookStoreOauthApi/src/domain/accessToken"
	"bookStoreOauthApi/src/errors"
)

type dbRepository struct {
}

type DbRepository interface {
	GetById(string) (*accessToken.AccessToken, *errors.RestError)
}

func (r *dbRepository) GetById(id string) (*accessToken.AccessToken, *errors.RestError) {
	return nil, nil
}

func NewRepository() DbRepository {
	return &dbRepository{}
}
