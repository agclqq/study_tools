package config

import (
	"github.com/agclqq/prow-framework/env/autoenv"
)

var db = map[string]map[string]string{
	"demo": {
		"alias":       autoenv.Get("DB_DEMO_DB_ALIAS"),
		"charset":     autoenv.Get("DB_DEMO_CHARSET"),
		"db":          autoenv.Get("DB_DEMO_DB"),
		"driver":      autoenv.Get("DB_DEMO_DRIVER"),
		"host":        autoenv.Get("DB_DEMO_HOST"),
		"log":         autoenv.Get("DB_DEMO_SQLLOG"),
		"maxIdle":     autoenv.Get("DB_DEMO_MAX_IDLE"),
		"maxIdleTime": autoenv.Get("DB_DEMO_MAX_IDLE_TIME"),
		"maxLife":     autoenv.Get("DB_DEMO_MAX_LIFE"),
		"maxOpen":     autoenv.Get("DB_DEMO_MAX_OPEN"),
		"password":    autoenv.Get("DB_DEMO_PASS"),
		"port":        autoenv.Get("DB_DEMO_PORT"),
		"user":        autoenv.Get("DB_DEMO_NAME"),
	},
}

func GetAllDb() map[string]map[string]string {
	return db
}
func GetDb(key string) map[string]string {
	return db[key]
}
