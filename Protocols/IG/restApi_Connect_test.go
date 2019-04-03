package IG

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestApi_Login(t *testing.T) {
	igConfiguration, err := NewConfiguration()
	if !assert.NoError(t, err) {
		return
	}
	httpTimeout := time.Duration(5 * time.Second)
	t.Run("Correct Login", func(t *testing.T) {
		ig := NewIgRestApi(
			os.Stdout,
			igConfiguration.ApiUrl,
			igConfiguration.ApiKey,
			"",
			igConfiguration.Identifier,
			igConfiguration.Password,
			httpTimeout)
		err := ig.Login()
		if !assert.NoError(t, err, "fail on login") {
			return
		}
		err = ig.Logout()
		assert.NoError(t, err, "fail on logout")
	})

	t.Run("Incorrect Login", func(t *testing.T) {
		ig := NewIgRestApi(
			os.Stdout,
			igConfiguration.ApiUrl,
			igConfiguration.ApiKey,
			"",
			"XXXX",
			"YYYY",
			httpTimeout)
		err := ig.Login()
		assert.Error(t, err)
		_, ok := err.(*ErrorMessage)
		assert.True(t, ok)
		err = ig.Logout()

		if !assert.Error(t, err, "should throw error") {
			return
		}
		errorCode, _ := err.(*ErrorMessage)
		assert.Equal(t, "error.security.client-token-missing", errorCode.ErrorCode)
	})

	t.Run("Login and replace token and logout", func(t *testing.T) {
		ig := NewIgRestApi(
			os.Stdout,
			igConfiguration.ApiUrl,
			igConfiguration.ApiKey,
			"",
			igConfiguration.Identifier,
			igConfiguration.Password,
			httpTimeout)
		err := ig.Login()
		if !assert.NoError(t, err, "fail on login") {
			return
		}

		err = ig.RefreshToken()
		if !assert.NoError(t, err, "fail on refresh token") {
			return
		}
		err = ig.Logout()
		assert.NoError(t, err, "fail on logout")
	})

	t.Run("Between log and logout", func(t *testing.T) {
		ig := NewIgRestApi(
			os.Stdout,
			igConfiguration.ApiUrl,
			igConfiguration.ApiKey,
			"",
			igConfiguration.Identifier,
			igConfiguration.Password,
			httpTimeout)
		err := ig.Login()
		if !assert.NoError(t, err, "fail on login") {
			return
		}
		defer func() {
			err = ig.Logout()
			assert.NoError(t, err, "fail on logout")
		}()
		t.Run("GetSession", func(t *testing.T) {
			session, err := ig.GetSession(true)
			assert.NoError(t, err)
			assert.NotNil(t, session)
			fmt.Println(session)
		})
	})
}
