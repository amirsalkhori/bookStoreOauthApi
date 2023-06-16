package accessToken

import (
	"bookStoreOauthApi/src/errors"
	"strings"
	"time"
)

const (
	expirationTime      = 24
	grandTypePassword   = "password"
	grantTypeCredential = "client_credential"
)

type AccessTokenRequest struct {
	GrantType string `json:"grantType"`
	Scope     string `json:"scope"`

	//Used for password grant type
	Username string `json:"username"`
	Password string `json:"password"`

	//Used for client credentials grant type
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

func (acr *AccessTokenRequest) Validate() *errors.RestError {
	switch acr.GrantType {
	case grandTypePassword:
		break
	case grantTypeCredential:
		break
	default:
		return errors.NewBadRequestError("Invalid grant_type parameter!")
	}

	return nil
}

type AccessToken struct {
	AccessToken string `json:"accessToken"`
	UserId      int64  `json:"userId"`
	ClientId    int64  `json:"clientId"`
	Expires     int64  `json:"expires"`
}

func (at AccessToken) Validate() *errors.RestError {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("Invalid access token id!")
	}
	if at.UserId <= 0 {
		return errors.NewBadRequestError("Invalid userId!")
	}
	if at.ClientId <= 0 {
		return errors.NewBadRequestError("Invalid clientId!")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequestError("Invalid Expire time!")
	}

	return nil
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}
