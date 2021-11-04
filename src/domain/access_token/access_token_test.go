package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "expiration time should be 24 hours")
}

func TestGetNewAcceessToken(t *testing.T) {
	at := GetNewAcceessToken()
	assert.False(t, at.IsExpired(), "brand new access token should not be nil")
	assert.EqualValues(t, "", at.AccessToken, "new access token should not have defined access token id")
	assert.Truef(t, at.UserId == 0, "new access token should not have an associated user id")
}

func TestAccessToken_IsExpired(t *testing.T) {
	at := AccessToken{}
	assert.Truef(t, at.IsExpired(), "empty access token should be expired by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired())
}
