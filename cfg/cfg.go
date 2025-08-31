package cfg

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	View map[string]map[string][]string `yaml:"view"`
	Open map[string]map[string][]string `yaml:"open"`
	Exec map[string]map[string][]string `yaml:"exec"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
