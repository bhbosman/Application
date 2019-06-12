package IG

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/url"
	"os"
	"testing"
	"time"
)

func TestLightStreamConnect2(t *testing.T) {
	igConfiguration, err := NewConfiguration()
	if !assert.NoError(t, err) {
		return
	}
	httpTimeout := time.Duration(5 * time.Second)

	ig := NewIgRestApi(
		os.Stdout,
		igConfiguration.ApiUrl,
		igConfiguration.ApiKey,
		"",
		igConfiguration.Identifier,
		igConfiguration.Password,
		httpTimeout)

	err = ig.Login()
	if !assert.NoError(t, err, "fail on login") {
		return
	}
	defer func() {
		err = ig.Logout()
		assert.NoError(t, err, "fail on logout")
	}()

	sessionInfo, err := ig.GetSession(true)
	assert.NoError(t, err)
	assert.NotNil(t, sessionInfo)
	lightStreamConnection, _ := CreateLightStreamConnection(os.Stdout, sessionInfo.LightstreamerEndpoint)
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

	defer func() {
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
	}()

	t.Run("sss", func(t *testing.T) {
		subscription := &QuoteAdapterSubscription{
			Items: []string{
				//"MARKET:IX.D.SAF.IFD.IP",
				//"MARKET:IX.D.SAF.IFM.IP",
				"MARKET:AR.D.STX40.CASH.IP",
			},
			Fields: []string{
				"BID",
				"OFFER",
				"HIGH",
				"LOW",
				"MID_OPEN",
				"CHANGE",
				"CHANGE_PCT",
				"MARKET_DELAY",
				"MARKET_STATE",
				"UPDATE_TIME",
			},
			Mode:        "MERGE",
			DataAdapter: "DEFAULT",
		}

		ca, _, _ := lightStreamConnection.Subscribe(subscription)
		fmt.Println(ca)
		time.Sleep(time.Hour * 24)
	})
}
