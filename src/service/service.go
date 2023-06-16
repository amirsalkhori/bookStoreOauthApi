package service

import (
	"bookStoreOauthApi/src/domain/accessToken"
	"bookStoreOauthApi/src/errors"
	"strings"
)

type Repository interface {
	GetById(string) (*accessToken.AccessToken, *errors.RestError)
	Create(accessToken.AccessToken) *errors.RestError
	UpdateExpirationTime(accessToken.AccessToken) *errors.RestError
}

type Service interface {
	GetById(string) (*accessToken.AccessToken, *errors.RestError)
	Create(accessToken.AccessToken) *errors.RestError
	UpdateExpirationTime(accessToken.AccessToken) *errors.RestError
}

type servie struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &servie{
		repository: repo,
	}
}

func (s *servie) GetById(accessTokenId string) (*accessToken.AccessToken, *errors.RestError) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("Invalid access token !")
	}
	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *servie) Create(at accessToken.AccessToken) *errors.RestError {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.Create(at)
}

func (s *servie) UpdateExpirationTime(at accessToken.AccessToken) *errors.RestError {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.UpdateExpirationTime(at)
}
