package clevertap

import (
	"context"
	"crypto/tls"
	"math/rand"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

const (
	cleverTapURL  = "https://api.clevertap.com"
	accountID     = "TEST-000-000-000"
	passcode      = "AAA-AAA-AAA"
	testIdentity  = "eko@kitabisa.com"
	testEventName = "Golang SDK Test Event"
)

var okResponse = make([]string, 0)
var unauthorizedResponse = make([]string, 0)
var nonJSONResponse = make([]string, 0)
var eventData = make(map[string]interface{})

func init() {
	ok := [...]string{
		`{
	"status": "success",
	"processed": 1,
	"unprocessed": []
}`,
	}

	okResponse = ok[:]

	unauthorized := [...]string{
		`{
			"status": "fail",
			"error": "12 digit account ID mandatory",
			"code": 401
}`,
		`{
    "status": "fail",
    "error": "Account ID is not valid",
    "code": 401
}`,
		`{
    "status": "fail",
    "error": "Passcode Invalid",
    "code": 401
}`,
	}

	unauthorizedResponse = unauthorized[:]

	nonJSON := [...]string{
		`{ hello world`,
	}

	nonJSONResponse = nonJSON[:]

	eventData["full_name"] = "Test Name1"
	eventData["user_id_type"] = "email"
	eventData["social_media_id"] = "11111"
}

func setCleverTapBuild(httpClient *http.Client) BuildClevertap {
	clevertapBuilder := &Builder{}
	service := &Service{}
	baseURL, _ := url.Parse(cleverTapURL)

	clevertapBuilder.SetBuilder(service)
	clevertapBuilder.SetHTTPClient(httpClient)
	clevertapBuilder.SetBaseURL(baseURL)
	clevertapBuilder.SetAccountID(accountID)
	clevertapBuilder.SetPasscode(passcode)

	return clevertapBuilder.Build()
}

func TestSendEvent(t *testing.T) {
	handler := anyHandler(okResponse, 200)
	httpClient, teardown := testingHTTPClient(handler)
	defer teardown()
	cleverTap := setCleverTapBuild(httpClient)
	cleverTapResponse := &Response{}

	if err := cleverTap.SendEvent(testIdentity, testEventName, eventData, cleverTapResponse); err != nil {
		t.Errorf("Fail [%v]", err)
	} else {
		t.Log("ok")
	}
}

func TestSendEventUnauthorized(t *testing.T) {
	handler := anyHandler(unauthorizedResponse, 401)
	httpClient, teardown := testingHTTPClient(handler)
	defer teardown()
	cleverTap := setCleverTapBuild(httpClient)
	cleverTapResponse := &Response{}

	if err := cleverTap.SendEvent(testIdentity, testEventName, eventData, cleverTapResponse); err != nil || cleverTapResponse.Status != "fail" {
		t.Errorf("Fail Got error : [%v] or status : [%s]", err, cleverTapResponse.Status)
	} else {
		t.Log("ok")
	}
}

func TestSendEventGotNonJsonResponse(t *testing.T) {
	handler := anyHandler(nonJSONResponse, 500)
	httpClient, teardown := testingHTTPClient(handler)
	defer teardown()
	cleverTap := setCleverTapBuild(httpClient)
	cleverTapResponse := &Response{}

	if err := cleverTap.SendEvent(testIdentity, testEventName, eventData, cleverTapResponse); err != nil {
		t.Logf("ok, Expected Error [%v]", err)
	} else {
		t.Error("Fail, Expected error not found")
	}
}

func TestSendEventErrorTimeout(t *testing.T) {
	handler := timeoutHandler()
	httpClient, teardown := testingHTTPClient(handler)
	defer teardown()
	cleverTap := setCleverTapBuild(httpClient)
	cleverTapResponse := &Response{}

	if err := cleverTap.SendEvent(testIdentity, testEventName, eventData, cleverTapResponse); err != nil {
		t.Logf("ok, Expected Error [%v]", err)
	} else {
		t.Error("Fail, Expected error not found")
	}
}

func BenchmarkSendEvent(b *testing.B) {
	handler := anyHandler(okResponse, 200)
	httpClient, teardown := testingHTTPClient(handler)
	defer teardown()
	cleverTap := setCleverTapBuild(httpClient)

	for n := 0; n < b.N; n++ {
		cleverTapResponse := &Response{}
		_ = cleverTap.SendEvent(testIdentity, testEventName, eventData, cleverTapResponse)
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
		Timeout: 1000 * time.Millisecond,
	}

	return cli, s.Close
}

func anyHandler(bodyResponse []string, httpStatus int) (handler http.Handler) {
	rand.Seed(time.Now().Unix())
	body := bodyResponse[rand.Intn(len(bodyResponse))]
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(httpStatus)
		_, _ = w.Write([]byte(body))
	})
}

func timeoutHandler() (handler http.Handler) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		w.WriteHeader(http.StatusGatewayTimeout)
	})
}
