package config

import "os"

type HttpServerConfig struct {
	Port string
	Env  string
}

type Configs struct {
	HttpServer *HttpServerConfig
}

func GetConfig() *Configs {
	return &Configs{
		HttpServer: &HttpServerConfig{Port: getEnv("HTTP_PORT", "8081"), Env: getEnv("HTTP_ENV", "dev")},
	}
}

func getEnv(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return value
}
