package clevertap

import (
	"net/http"
	"net/url"
)

// Options ...
type Options struct {
	httpClient *http.Client
	baseURL    *url.URL
	AccountID  string
	Passcode   string
}

// SendEventRequest ...
type SendEventRequest struct {
	Identity  string                 `json:"identity"`
	Type      string                 `json:"type"`
	Timestamp int64                  `json:"ts"`
	EventName string                 `json:"evtName"`
	EventData map[string]interface{} `json:"evtData"`
}

// Response ...
type Response struct {
	Status      string        `json:"status"`
	Processed   int           `json:"processed"`
	Unprocessed []interface{} `json:"unprocessed"`
}

type SendProfileRequest struct {
	Identity    string                 `json:"identity"`
	Type        string                 `json:"type"`
	Timestamp   int64                  `json:"ts"`
	ProfileData map[string]interface{} `json:"profileData"`
}
