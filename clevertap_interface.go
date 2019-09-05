package go_sdk_clevertap

type BuildClevertap interface {
	SetOptions(clevertapOptions ClevertapOptions) BuildClevertap
	SendEvent(identity string, evtName string, evtData map[string]interface{}) error
}
