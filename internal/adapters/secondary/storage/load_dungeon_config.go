package storage

import (
	"os"

	"gopkg.in/yaml.v3"
)

// LoadDungeonConfig reads and parses a YAML configuration file for dungeon generation.
// The function loads the file from the specified path and unmarshals it into a Config struct.
//
// Parameters:
//   - path: string - filesystem path to the YAML configuration file
//
// Returns:
//   - *Config: pointer to the parsed configuration structure
//   - error: any error that occurred during file reading or YAML parsing
//
// Example usage:
//
//	cfg, err := LoadDungeonConfig("config/game_config.yaml")
//	if err != nil {
//	    - handle error...
//	}
func LoadDungeonConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	return &cfg, err
}
