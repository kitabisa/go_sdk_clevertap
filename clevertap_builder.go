package clevertap

import (
	"net/http"
	"net/url"
)

// Builder ...
type Builder struct {
	option  Options
	builder BuildClevertap
}

// SetBuilder ...
func (c *Builder) SetBuilder(b BuildClevertap) {
	c.builder = b
}

// SetHTTPClient ...
func (c *Builder) SetHTTPClient(httpClient *http.Client) {
	c.option.httpClient = httpClient
}

// SetBaseURL ...
func (c *Builder) SetBaseURL(baseURL *url.URL) {
	c.option.baseURL = baseURL
}

// SetAccountID ...
func (c *Builder) SetAccountID(accountID string) {
	c.option.AccountID = accountID
}

// SetPasscode ...
func (c *Builder) SetPasscode(passcode string) {
	c.option.Passcode = passcode
}

// Build ...
func (c *Builder) Build() BuildClevertap {
	c.builder.setOptions(c.option)
	return c.builder
}
