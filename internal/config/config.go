package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Name string   `yaml:"name"`
	Env  []string `yaml:"env"`
}

func LoadConfig() (*Config, error) {
	f, err := os.ReadFile("config/dev.yaml")
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = yaml.Unmarshal(f, &cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
