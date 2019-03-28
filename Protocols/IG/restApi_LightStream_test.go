package IG

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
	"time"
)

func TestLightStreamConnect(t *testing.T) {
	igConfiguration, err := NewConfiguration()
	if !assert.NoError(t, err) {
		return
	}
	httpTimeout := time.Duration(5 * time.Second)
	t.Run("Between log and logout", func(t *testing.T) {
		ig := NewIgRestApi(igConfiguration.ApiUrl, igConfiguration.ApiKey, "", igConfiguration.Identifier, igConfiguration.Password, httpTimeout)
		err := ig.Login()
		if !assert.NoError(t, err, "fail on login") {
			return
		}
		defer func() {
			err = ig.Logout()
			assert.NoError(t, err, "fail on logout")
		}()
		t.Run("Connect/Disconnect to light stream", func(t *testing.T) {
			sessionInfo, err := ig.GetSession(true)
			assert.NoError(t, err)
			assert.NotNil(t, sessionInfo)
			lightStreamConnection, _ := CreateLightStreamConnection(sessionInfo.LightstreamerEndpoint)
			ans, err := lightStreamConnection.Connect(
				func(data url.Values) {
					data.Set("LS__user", ig.AccountID)
					data.Set("LS_password", fmt.Sprintf("CST-%v|XST-%v", sessionInfo.Cst, sessionInfo.XSECURITYTOKEN))
					data.Set("LS_adapter_set", "DEFAULT")
					data.Set("LS_polling", "true")     // same parameters as streaming companion
					data.Set("LS_polling_millis", "0") // only if polling is true
					data.Set("LS_idle_millis", "0")    // only if polling is true
				})
			if !assert.NoError(t, err) {
				return
			}
			if !assert.NotNil(t, ans) {
				return
			}
			if !assert.True(t, ans.Success) {
				return
			}

			ans, err = lightStreamConnection.Disconnect()
			if !assert.NoError(t, err) {
				return
			}
			if !assert.NotNil(t, ans) {
				return
			}
			if !assert.True(t, ans.Success) {
				return
			}
		})
	})
}
