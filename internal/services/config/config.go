// Package config provides functionality for loading and parsing the game's configuration file.
// It supports YAML format and includes settings for logging and terminal window dimensions.
package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Config represents the full application configuration loaded from a YAML file.
type Config struct {
	LoggerInfo `yaml:"logger"` // Logging configuration
	Width      int             `yaml:"width"`  // Terminal window width (in characters)
	Height     int             `yaml:"height"` // Terminal window height (in lines)
}

// LoggerInfo contains settings related to logging behavior.
type LoggerInfo struct {
	Level      string `yaml:"level"`  // Log level (e.g., "debug", "info", "warn", "error")
	OutputFile string `yaml:"output"` // Path to the log output file
}

// LoadConfig reads and parses the YAML configuration file from the given path.
// Returns a Config pointer and nil error on success, or nil and an error on failure.
func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

// MustLoadConfig is a convenience wrapper around LoadConfig that panics if the config cannot be loaded.
// Useful for cases where missing or invalid config should immediately stop the application.
func MustLoadConfig(path string) *Config {
	cfg, err := LoadConfig(path)
	if err != nil {
		panic(err)
	}

	return cfg
}
