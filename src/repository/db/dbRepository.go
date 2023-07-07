package db

import (
	"bookStoreOauthApi/src/dataSorces/mysql/authDB"
	"bookStoreOauthApi/src/domain/accessToken"
	"bookStoreOauthApi/src/errors"
	"fmt"

	"gorm.io/gorm"
)

func connection() (*gorm.DB, *errors.RestError) {
	db, err := authDB.Connect()
	if err != nil {
		fmt.Println("Error when tryin to connect to db", err)
		return nil, errors.NewInternamlServerError(err.Error())
	}

	return db, nil
}

type dbRepository struct {
}

type DbRepository interface {
	GetById(string) (*accessToken.AccessToken, *errors.RestError)
	Create(accessToken.AccessToken) *errors.RestError
	UpdateExpirationTime(accessToken.AccessToken) *errors.RestError
}

func (r *dbRepository) GetById(accessTokenId string) (*accessToken.AccessToken, *errors.RestError) {
	db, errorConnection := connection()
	if errorConnection != nil {
		return nil, errorConnection
	}
	var at accessToken.AccessToken

	err := db.Where("access_token = ?", accessTokenId).Find(&at).Error
	if err != nil {
		fmt.Println("Error when tryin to get user", err)
		return nil, errors.NewInternamlServerError(err.Error())
	}
	if at.AccessToken == "" {
		return nil, errors.NewNotFoundError("Access token is not valid...")
	}
	return &at, nil
}

func (r *dbRepository) Create(at accessToken.AccessToken) *errors.RestError {

	db, errorConnection := connection()
	if errorConnection != nil {
		return errorConnection
	}
	accessTokenItem := accessToken.AccessToken{
		AccessToken: at.AccessToken,
		UserId:      at.UserId,
		ClientId:    at.ClientId,
		Expires:     at.Expires,
	}

	fmt.Println("AccessToken is:", accessTokenItem)

	if err := db.Create(&accessTokenItem).Error; err != nil {
		fmt.Println("Error:", err.Error())
		return errors.NewInternamlServerError(err.Error())
	}

	return nil
}

func (r *dbRepository) UpdateExpirationTime(at accessToken.AccessToken) *errors.RestError {

	db, errorConnection := connection()
	if errorConnection != nil {
		return errorConnection
	}
	var accessToken accessToken.AccessToken
	err := db.Save(&accessToken).Error
	if err != nil {
		fmt.Println("Error when tryin to update accessToken", err)
		return errors.NewBadRequestError("Error when tryin to update accessToken")

	}

	return nil
}

func NewRepository() DbRepository {
	return &dbRepository{}
}
