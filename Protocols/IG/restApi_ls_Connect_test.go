package IG

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestApi_Marketsd(t *testing.T) {
	igConfiguration, err := NewConfiguration()
	if !assert.NoError(t, err) {
		return
	}
	httpTimeout := time.Duration(5 * time.Second)
	t.Run("Between log and logout", func(t *testing.T) {
		ig := New(igConfiguration.ApiUrl, igConfiguration.ApiKey, "", igConfiguration.Identifier, igConfiguration.Password, httpTimeout)
		err := ig.Login()
		if !assert.NoError(t, err, "fail on login") {
			return
		}
		defer func() {
			err = ig.Logout()
			assert.NoError(t, err, "fail on logout")
		}()
		t.Run("Connect for prices", func(t *testing.T) {
		})
	})
}