package config

import (
	"accounting/accounting/data/database"
	"encoding/json"
	"io/ioutil"
)

// AppConfig is type which represents application config
type AppConfig struct {
	Database database.Database
}

// Config is singleton instance of AppConfig type
var Config *AppConfig

// NewAppConfig creates only one instance of AppConfig type
func NewAppConfig() *AppConfig {
	if Config == nil {
		Config = &AppConfig{}
		file, _ := ioutil.ReadFile("config/config.json")
		json.Unmarshal(file, &Config)
	}

	return Config
}
