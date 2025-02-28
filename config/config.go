package config

import (
	"fmt"
	"os"
	"path"

	"github.com/zetatez/lazy/sugar"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Cmds []string `yaml:"cmds"`
}

func LoadConfig(suffix string) (cfg *Config, err error) {
	prefixes := []string{
		"./etc",
		path.Join(os.Getenv("HOME"), ".config", "lazy", "etc"),
	}
	cfg = &Config{}
	found := false
	for _, prefix := range prefixes {
		f := path.Join(prefix, suffix)
		if !sugar.IsFileExists(f) {
			continue
		}
		fbyte, err := os.ReadFile(f)
		if err != nil {
			return nil, err
		}
		if err = yaml.Unmarshal(fbyte, &cfg); err != nil {
			return nil, err
		}
		found = true
	}
	if !found {
		return nil, fmt.Errorf("config file %s not found", suffix)
	}
	return cfg, nil
}
