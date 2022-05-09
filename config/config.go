package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type ApiConfig struct {
	Port          string `json:"port"`
	Path          string `json:"path"`
	LogLevel      string `json:"loglevel"`
	ClientTimeout int    `json:"timeout"`
	Backend       string `json:"backend"`
}

func GetApiConfig() ApiConfig {
	var config ApiConfig
	fmt.Println(ColorCyan, "-------CONFIG----------")
	config, err := LoadConfiguration("app.json")
	if err != nil {
		fmt.Println(ColorCyan, "Loading default configuration")
		config = ApiConfig{
			Port:          "8080",
			Path:          "/iso2json",
			LogLevel:      "debug",
			ClientTimeout: 30,
			Backend:       "isoreader:8082",
		}
	}
	c, _ := json.Marshal(config)
	fmt.Println(ColorGreen, string(c))
	return config
}

func LoadConfiguration(file string) (ApiConfig, error) {
	var config ApiConfig
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(ColorRed, "Error opening properties file: "+err.Error())
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	errDecode := jsonParser.Decode(&config)
	if errDecode != nil {
		fmt.Println(ColorRed, "Error reading properties file: "+errDecode.Error())
		return config, errDecode
	}
	return config, nil
}

const (
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
)
