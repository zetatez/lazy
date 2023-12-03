package config

import (
	"fmt"
	"os"
	"path"
	"sync"

	"github.com/zetatez/lazy/utils/sugar"
	"gopkg.in/yaml.v2"
)

var (
	cfg  *Config
	once sync.Once
)

type Config struct {
	View struct {
		Preset map[string]string `yaml:"preset"`
		Ext    map[string]string `yaml:"ext"`
		Mime   map[string]string `yaml:"mime"`
	} `yaml:"view"`
	Open struct {
		Preset map[string]string `yaml:"preset"`
		Ext    map[string]string `yaml:"ext"`
		Mime   map[string]string `yaml:"mime"`
	} `yaml:"open"`
	Exec struct {
		Preset map[string]string `yaml:"preset"`
		Ext    map[string]string `yaml:"ext"`
		Mime   map[string]string `yaml:"mime"`
	} `yaml:"exec"`
}

func GetConfig() *Config {
	once.Do(func() {
		cfg = &Config{}
	})

	return cfg
}

func (s *Config) ReLoadCfg() (err error) {
	return s.LoadCfg()
}

func (s *Config) LoadCfg() (err error) {
	configFileList := []string{
		"lazy.yaml",
		path.Join(os.Getenv("HOME"), ".config", "lazy", "lazy.yaml"),
		path.Join(os.Getenv("HOME"), ".lazy.yaml"),
		"/etc/lazy.yaml",
	}
	for _, f := range configFileList {
		isExists, err := sugar.IsFileExists(f)
		if err != nil {
			return err
		}
		if !isExists {
			continue
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

	return fmt.Errorf("no config file was found")
}
