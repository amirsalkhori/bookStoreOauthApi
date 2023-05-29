package db

import (
	"bookStoreOauthApi/src/client/cassandra"
	"bookStoreOauthApi/src/domain/accessToken"
	"bookStoreOauthApi/src/errors"
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

type dbRepository struct {
}

type DbRepository interface {
	GetById(string) (*accessToken.AccessToken, *errors.RestError)
	Create(accessToken.AccessToken) *errors.RestError
	UpdateExpirationTime(accessToken.AccessToken) *errors.RestError
}

func (r *dbRepository) GetById(id string) (*accessToken.AccessToken, *errors.RestError) {
	session, err := cassandra.GetSession()
	if err != nil {
		return nil, errors.NewInternamlServerError(err.Error())
	}
	defer session.Close()

	var result accessToken.AccessToken
	insertStmt := `SELECT * FROM access_token WHERE access_token=?`
	if err := session.Query(insertStmt, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("No access token found!")
		}
	}
	return &result, nil
}

func (r *dbRepository) Create(at accessToken.AccessToken) *errors.RestError {
	session, err := cassandra.GetSession()
	if err != nil {
		return errors.NewInternamlServerError(err.Error())
	}
	defer session.Close()

	insertStmt := "INSERT INTO oauth.access_token (access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?)"

	// Convert the expires value to time.Time
	// at.Expires = time.Date(2023, time.May, 30, 12, 0, 0, 0, time.UTC)

	// Execute the INSERT statement
	err = session.Query(insertStmt, at.AccessToken, at.UserId, at.ClientId, at.Expires).Exec()
	if err != nil {
		fmt.Println("error:", err.Error())
		return errors.NewInternamlServerError(err.Error())
	}
	log.Println("Data inserted successfully!")
	return nil
}

func (r *dbRepository) UpdateExpirationTime(at accessToken.AccessToken) *errors.RestError {
	session, err := cassandra.GetSession()
	if err != nil {
		return errors.NewInternamlServerError(err.Error())
	}
	defer session.Close()

	updateStmt := "UPDATE oauth.access_token SET expires = ? where access_token = ?"
	err = session.Query(updateStmt, at.Expires, at.AccessToken).Exec()
	if err != nil {
		fmt.Println("error:", err.Error())
		return errors.NewInternamlServerError(err.Error())
	}
	log.Println("Data Updated successfully!")
	return nil
}

func NewRepository() DbRepository {
	return &dbRepository{}
}
