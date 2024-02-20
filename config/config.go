package config

import (
	"fmt"
	"os"
	"path"

	"github.com/zetatez/lazy/utils/sugar"
	"gopkg.in/yaml.v2"
)

type Config struct {
	CMD string `yaml:"cmd"`
}

func NewConfig() *Config {
	return &Config{}
}

func (s *Config) load(f string) (err error) {
	isExists, err := sugar.IsFileExists(f)
	if err != nil {
		return err
	}
	if !isExists {
		fmt.Errorf("config file %s not exists", f)
	}
	fbyte, err := os.ReadFile(f)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(fbyte, &s); err != nil {
		return err
	}
	return nil
}

func GetConfig(suffix string) (cfg *Config) {
	prefixes := []string{
		"./etc",
		path.Join(os.Getenv("HOME"), ".config", "lazy", "etc"),
	}
	s := NewConfig()
	for _, prefix := range prefixes {
		f := path.Join(prefix, suffix)
		err := s.load(f)
		if err != nil {
			continue
		}
	}

	return cfg
}
