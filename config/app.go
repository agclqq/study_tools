package config

import (
	"github.com/agclqq/prow-framework/env/autoenv"
)

var app = map[string]string{
	"appEnv":         autoenv.Get("APP_ENV"),
	"easyCryptIv":    autoenv.Get("EASY_CRYPT_IV"),
	"easyCryptKey":   autoenv.Get("EASY_CRYPT_KEY"),
	"easyCryptType":  autoenv.Get("EASY_CRYPT_TYPE"),
	"grpcServerPort": autoenv.Get("GRPC_SERVER_PORT"),
	"httpServerPort": autoenv.Get("HTTP_SERVER_PORT"),
	"logFile":        autoenv.Get("LOG_FILE"),
	"logRetain":      autoenv.Get("LOG_RETAIN"),
}

func GetApp(key string) string {
	return app[key]
}
