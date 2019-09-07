package go_sdk_clevertap

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type CleverTapService struct {
	cO ClevertapOptions
}

func (c *CleverTapService) setOptions(clevertapOptions ClevertapOptions) BuildClevertap {
	c.cO = clevertapOptions
	return c
}

func (c *CleverTapService) SendEvent(identity string, evtName string, evtData map[string]interface{}, responseInterface interface{}) error {
	sendEventReq := []CleverTapSendEventRequest{
		{
			Identity:  identity,
			EventName: evtName,
			Type:      EVENT,
			Timestamp: time.Now().Unix(),
			EventData: evtData,
		},
	}

	payload := make(map[string]interface{})
	payload["d"] = sendEventReq

	if req, err := c.newRequest(POST, CLEVERTAP_SEND_EVENT_URL, payload); err != nil {
		return err
	} else if _, err = c.do(req, responseInterface); err != nil {
		return err
	}

	return nil
}

func (c *CleverTapService) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.cO.baseURL.ResolveReference(rel)
	var buf io.ReadWriter

	if body != nil {
		buf = new(bytes.Buffer)
		if err := json.NewEncoder(buf).Encode(body); err != nil {
			return nil, err
		}
	}

	if req, err := http.NewRequest(method, u.String(), buf); err != nil {
		return nil, err
	} else {
		req.Header.Set(CONTENT_TYPE, APPLICATION_JSON_CHARSET_UTF8)
		req.Header.Set(X_CLEVERTAP_ACCOUNT_ID, c.cO.AccountID)
		req.Header.Set(X_CLEVERTAP_PASSCODE, c.cO.Passcode)
		return req, nil
	}
}

func (c *CleverTapService) do(req *http.Request, v interface{}) (*http.Response, error) {
	if resp, err := c.cO.httpClient.Do(req); err != nil {
		return nil, err
	} else {
		defer func() {
			_ = resp.Body.Close()
		}()

		if bodyByte, err := ioutil.ReadAll(resp.Body); err != nil {
			return resp, err
		} else {
			return resp, json.Unmarshal(bodyByte, &v)
		}
	}
}
