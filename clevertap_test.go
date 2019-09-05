package go_sdk_clevertap

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

const (
	cleverTapUrl  = "https://api.clevertap.com"
	accountId     = "TEST-000-000-000"
	passcode      = "AAA-AAA-AAA"
	testIdentity  = "eko@kitabisa.com"
	testEventName = "Golang SDK Test Event"

	okResponse = `{
		"status": "success",
		"processed": 1,
		"unprocessed": []
	}`
)

func TestSendEvent(t *testing.T) {
	clevertapBuilder := &ClevertapBuilder{}
	service := &CleverTapService{}

	ok := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(okResponse))
	})

	httpClient, teardown := testingHTTPClient(ok)
	defer teardown()

	baseUrl, _ := url.Parse(cleverTapUrl)

	clevertapBuilder.SetBuilder(service)
	clevertapBuilder.SetHttpClient(httpClient)
	clevertapBuilder.SetBaseURL(baseUrl)
	clevertapBuilder.SetAccountID(accountId)
	clevertapBuilder.SetPasscode(passcode)
	cleverTap := clevertapBuilder.Build()

	eventData := make(map[string]interface{})
	eventData["full_name"] = "Test Name1"
	eventData["user_id_type"] = "email"
	eventData["social_media_id"] = "11111"

	_ = cleverTap.SendEvent(testIdentity, testEventName, eventData)
}

func BenchmarkSendEvent(b *testing.B) {
	ok := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(okResponse))
	})

	httpClient, teardown := testingHTTPClient(ok)
	defer teardown()

	for n := 0; n < b.N; n++ {
		clevertapBuilder := &ClevertapBuilder{}
		service := &CleverTapService{}
		baseUrl, _ := url.Parse(cleverTapUrl)

		clevertapBuilder.SetBuilder(service)
		clevertapBuilder.SetHttpClient(httpClient)
		clevertapBuilder.SetBaseURL(baseUrl)
		clevertapBuilder.SetAccountID(accountId)
		clevertapBuilder.SetPasscode(passcode)
		cleverTap := clevertapBuilder.Build()

		eventData := make(map[string]interface{})
		eventData["full_name"] = "Test Name1"
		eventData["user_id_type"] = "email"
		eventData["social_media_id"] = "11111"
		_ = cleverTap.SendEvent(testIdentity, testEventName, eventData)
	}
}

func testingHTTPClient(handler http.Handler) (*http.Client, func()) {
	s := httptest.NewTLSServer(handler)

	cli := &http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, network, _ string) (net.Conn, error) {
				return net.Dial(network, s.Listener.Addr().String())
			},
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	return cli, s.Close
}
