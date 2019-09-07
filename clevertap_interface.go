package go_sdk_clevertap

type BuildClevertap interface {
	setOptions(clevertapOptions ClevertapOptions) BuildClevertap
	SendEvent(identity string, evtName string, evtData map[string]interface{}, responseInterface interface{}) error
}
