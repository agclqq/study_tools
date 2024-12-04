package config

import (
	"github.com/agclqq/prow-framework/env/autoenv"
)

var event = map[string]map[string]string{
	"test": {
		"capacity": autoenv.Get("EVENT_TEST_CAPACITY"),
		"name":     autoenv.Get("EVENT_TEST"),
	},
}

func GetAllEvent() map[string]map[string]string {
	return event
}
func GetEvent(key string) map[string]string {
	return event[key]
}
