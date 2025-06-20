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
	return nil
}
