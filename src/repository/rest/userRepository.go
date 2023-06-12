package rest

import (
	"bookStoreOauthApi/src/domain/users"
	"bookStoreOauthApi/src/errors"
	"encoding/json"
	"fmt"
	"time"

	"github.com/mercadolibre/golang-restclient/rest"
)

type RestUserRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestError)
}

type userRepository struct{}

var (
	userRestClient = rest.RequestBuilder{
		BaseURL: "localhost:8081",
		Timeout: 100 * time.Millisecond,
	}
)

func (r *userRepository) LoginUser(email string, password string) (*users.User, *errors.RestError) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	response := userRestClient.Post("/users/login", request)

	fmt.Println("Response:", response.Response)
	if response == nil || response.Response == nil {
		return nil, errors.NewInternamlServerError("Invalid restClient response when trying to login user!")
	}

	if response.StatusCode > 299 {
		var restErr errors.RestError
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, errors.NewInternamlServerError("Invalid error interface when trying login user!")
		}
		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternamlServerError("Error when trying to unmarshal users response!")
	}

	return &user, nil
}

func NewUserRepository() RestUserRepository {
	return &userRepository{}
}
