package go_sdk_clevertap

import (
	"net/http"
	"net/url"
)

type ClevertapBuilder struct {
	option ClevertapOptions
	builder BuildClevertap
}

func (c *ClevertapBuilder) SetBuilder(b BuildClevertap) {
	c.builder = b
}

func (c *ClevertapBuilder) SetHttpClient(httpClient *http.Client) {
	c.option.httpClient = httpClient
}

func (c *ClevertapBuilder) SetBaseURL(baseURL *url.URL) {
	c.option.baseURL = baseURL
}

func (c *ClevertapBuilder) SetAccountID(accountId string) {
	c.option.AccountID = accountId
}

func (c *ClevertapBuilder) SetPasscode(passcode string) {
	c.option.Passcode = passcode
}

func (c *ClevertapBuilder) Build() (BuildClevertap) {
	c.builder.SetOptions(c.option)
	return c.builder
}