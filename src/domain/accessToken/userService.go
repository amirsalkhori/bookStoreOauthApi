package accessToken

import (
	"bookStoreOauthApi/src/domain/users"
	"bookStoreOauthApi/src/errors"
	"bookStoreOauthApi/src/repository/rest"
)

type UserRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestError)
}

type UserService interface {
	LoginUser(string, string) (*users.User, *errors.RestError)
}

type userServie struct {
	userRepository UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userServie{
		userRepository: repo,
	}
}

func (us *userServie) LoginUser(email string, password string) (*users.User, *errors.RestError) {
	result, err := rest.RestUserRepository.LoginUser(rest.NewUserRepository(), email, password)
	if err != nil {
		return nil, err
	}
	
	return result, nil
}
