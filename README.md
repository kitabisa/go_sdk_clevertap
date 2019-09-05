# Go SDK Clevertap
Golang SDK for Clevertap

## Features
### Upload Event API
Based on `https://developer.clevertap.com/docs/upload-events-api`

#### Example
```go
package test

import (
    "net/url"
    "net/http"
    "time"
    "github.com/kitabisa/go_sdk_clevertap"
)

const (
	cleverTapUrl = "https://api.clevertap.com"
	accountId = "TEST-000-000-000"
	passcode = "AAA-AAA-AAA"
	testIdentity = "your@email.com"
	testEventName = "Golang SDK Test Event"
)

func main() {
	clevertapBuilder := &ClevertapBuilder{}
	service := &CleverTapService{}

	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}

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
```

## Installation
```bash
go get -u github.com/kitabisa/go_sdk_clevertap
```


## License
MIT License