package go_sdk_clevertap

import (
	"net/http"
	"net/url"
)

type ClevertapOptions struct {
	httpClient *http.Client
	baseURL    *url.URL
	AccountID  string
	Passcode   string
}

type CleverTapSendEventRequest struct {
	Identity  string                 `json:"identity"`
	Type      string                 `json:"type"`
	Timestamp int64                  `json:"ts"`
	EventName string                 `json:"evtName"`
	EventData map[string]interface{} `json:"evtData"`
}

type CleverTapResponse struct {
	Status      string `json:"status"`
	Processed   int    `json:"processed"`
	Unprocessed string `json:"unprocessed"`
}
