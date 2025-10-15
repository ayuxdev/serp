package config

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Test with a valid config file
	cfg, err := LoadConfig("config.yaml")
	if err != nil {
		t.Fatalf("LoadConfig() error = %v", err)
	}
	if cfg == nil {
		t.Fatal("LoadConfig() did not load config")
	}
	t.Logf("LoadConfig() loaded config: %+v", cfg)
}
