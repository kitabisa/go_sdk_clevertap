# Go SDK Clevertap
Golang SDK for Clevertap

## Features
### Upload Event API
Based on `https://developer.clevertap.com/docs/upload-events-api`

#### Example
```go
package main

import (
	"fmt"
	"github.com/kitabisa/go_sdk_clevertap"
	"net/http"
	"net/url"
	"time"
)
const (
	cleverTapUrl = "https://api.clevertap.com"
	accountId = "TEST-000-000-000"
	passcode = "AAA-AAA-AAA"
	testIdentity = "your@email.com"
	testEventName = "Golang SDK Test Event"
)

func main()  {
	clevertapBuilder := &go_sdk_clevertap.ClevertapBuilder{}
	service := &go_sdk_clevertap.CleverTapService{}

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

	cleverTapResponse := &go_sdk_clevertap.CleverTapResponse{}

	_ = cleverTap.SendEvent(testIdentity, testEventName, eventData, cleverTapResponse)

	fmt.Printf("%v", cleverTapResponse)
}
```

## Test & Benchmark
```bash
go test ./... -bench=.
```

## Installation
```bash
go get -u github.com/kitabisa/go_sdk_clevertap
```


## License
MIT License