package config

import "time"

type ApiConfig struct {
	Port          string
	Path          string
	LogLevel      string
	ClientTimeout time.Duration
	BackendURI    string
}

func GetApiConfig() ApiConfig {
	return ApiConfig{
		Port:          "8080",
		Path:          "/iso2json",
		LogLevel:      "debug",
		ClientTimeout: 30,
		BackendURI:    "isoreader:8082",
	}
}
