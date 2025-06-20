package config

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

// LoadConfig loads configuration from a YAML file
func LoadConfig(configPath string) error {
	if configPath == "" {
		configPath = DefaultConfigPath
	}
	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to read config file: %v", err)
	}
	if err := yaml.Unmarshal(data, &Config); err != nil {
		return fmt.Errorf("failed to parse config file: %v", err)
	}
	// Set default values if not provided
	if Config.Server.ListenPort != 0 {
		ListenPort = Config.Server.ListenPort
	}
	if Config.Server.ApiAuth != "" {
		ApiAuth = Config.Server.ApiAuth
	}
	if Config.Log.MaxSize != 0 {
		LogMaxSize = Config.Log.MaxSize
	}
	if Config.Log.MaxBackups != 0 {
		LogMaxBackups = Config.Log.MaxBackups
	}
	if Config.Log.MaxAge != 0 {
		LogMaxAge = Config.Log.MaxAge
	}
	if Config.Sqlite.DSN != "" {
		SqliteDSN = Config.Sqlite.DSN
	}else{
		return fmt.Errorf("sqlite DSN is not set in the configuration file")
	}
	return nil
}
