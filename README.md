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
	cleverTapURL = "https://api.clevertap.com"
	accountID = "TEST-000-000-000"
	passcode = "AAA-AAA-AAA"
	testIdentity = "your@email.com"
	testEventName = "Golang SDK Test Event"
)

func main()  {
	clevertapBuilder := &clevertap.Builder{}
	service := &clevertap.Service{}

	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}

	baseUrl, _ := url.Parse(cleverTapURL)

	clevertapBuilder.SetBuilder(service)
	clevertapBuilder.SetHTTPClient(httpClient)
	clevertapBuilder.SetBaseURL(baseUrl)
	clevertapBuilder.SetAccountID(accountID)
	clevertapBuilder.SetPasscode(passcode)
	ct := clevertapBuilder.Build()

	eventData := make(map[string]interface{})
	eventData["full_name"] = "Test Name1"
	eventData["user_id_type"] = "email"
	eventData["social_media_id"] = "11111"

	cleverTapResponse := &clevertap.Response{}

	err := ct.SendEvent(testIdentity, testEventName, eventData, cleverTapResponse)

	fmt.Printf("%v - %v", *cleverTapResponse, err)
}
```

## Test, Code Coverage & Benchmark
```bash
go test -v -cover ./... -bench=.
```

## Installation
```bash
go get -u github.com/kitabisa/go_sdk_clevertap
```


## License
MIT License