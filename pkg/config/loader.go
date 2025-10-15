package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadConfig(path string) (*Config, error) {
	_, err := os.Stat(path)
	if err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("failed to check if config file exists: %v", err)
	} else if os.IsNotExist(err) {
		// TODO: log.infof("Config file not found, creating default config at %s", path)
		// if config file doesn't exist we create a default config and return it
		if err := saveDefaultConfig(path); err != nil {
			return nil, fmt.Errorf("failed to create default config: %v", err)
		}
		return DefaultCfg(), nil // return default config
	}
	// if config file exists we read from it's value and return it
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %v", err)
	}
	var cfg *Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %v", err)
	}
	return cfg, nil
}

func saveDefaultConfig(path string) error {
	data, err := yaml.Marshal(DefaultCfg())
	if err != nil {
		return fmt.Errorf("failed to marshal default config: %v", err)
	}
	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write default config to %s: %v", path, err)
	}
	return nil
}
