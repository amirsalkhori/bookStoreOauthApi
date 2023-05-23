package accessToken

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "Expiration shoud be 24 hours")
}


func TestAccessTokenIsExpired(t *testing.T){
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "Empty access token should be expired by default...")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "Access token expiring three hours from now should not be expired")
}