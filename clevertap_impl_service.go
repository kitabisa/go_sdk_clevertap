package go_sdk_clevertap

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type CleverTapService struct {
	cO ClevertapOptions
}

func (c *CleverTapService) SetOptions(clevertapOptions ClevertapOptions) BuildClevertap {
	c.cO = clevertapOptions
	return c
}

func (c *CleverTapService) SendEvent(identity string, evtName string, evtData map[string]interface{}) error {
	sendEventReq := []CleverTapSendEventRequest{
		CleverTapSendEventRequest {
			Identity:  identity,
			EventName: evtName,
			Type:      EVENT,
			Timestamp: time.Now().Unix(),
			EventData: evtData,
		},
	}

	payload := make(map[string]interface{})
	payload["d"] = sendEventReq

	req, err := c.newRequest(POST, CLEVERTAP_SEND_EVENT_URL, payload)
	if err != nil {
		return err
	}
	cleverTapResponse := &CleverTapResponse{}
	_, err = c.do(req, cleverTapResponse)

	return err
}

func (c *CleverTapService) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.cO.baseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set(CONTENT_TYPE, APPLICATION_JSON_CHARSET_UTF8)
	req.Header.Set(X_CLEVERTAP_ACCOUNT_ID, c.cO.AccountID)
	req.Header.Set(X_CLEVERTAP_PASSCODE, c.cO.Passcode)
	return req, nil
}

func (c *CleverTapService) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.cO.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != 200 {
		bodyByte, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			log.Printf("clevertap response: %v", string(bodyByte))
		}
	}

	err = json.NewDecoder(resp.Body).Decode(v)
	x, err := ioutil.ReadAll(resp.Body)
	log.Printf("%v", x)
	return resp, err
}
