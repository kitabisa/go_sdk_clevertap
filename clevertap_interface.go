package clevertap

// BuildClevertap ...
type BuildClevertap interface {
	setOptions(clevertapOptions Options) BuildClevertap
	SendEvent(identity string, evtName string, evtData map[string]interface{}, responseInterface interface{}) error
	SendProfile(identity string, profileData map[string]interface{}) error
}
