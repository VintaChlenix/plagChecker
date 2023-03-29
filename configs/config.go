package configs

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Server struct {
		URL string `yaml:"SERVER_URL"`
	} `yaml:"server"`
	Database struct {
		ConnectionString string `yaml:"DATABASE_URL"`
	} `yaml:"database"`
	ReferenceValues struct {
		DiffValue    float64 `yaml:"DIFF_VALUE"`
		TokensValue  float64 `yaml:"TOKENS_VALUE"`
		MetricsValue float64 `yaml:"METRICS_VALUE"`
	} `yaml:"referenceValues"`
}

func GetConfig() (*Config, error) {
	f, err := os.Open("configs/config.yml")
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &cfg, nil
}
