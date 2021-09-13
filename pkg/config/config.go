package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Image
type Image struct {
	Name string `yaml:"name"`
	Tag  string `yaml:"tag"`
}

// Charts
type Charts struct{}

// CloudProvider
type CloudProvider struct{}

// Networking
type Networking struct {
	DefaultCNI string   `yaml:"default_cni"`
	CNIs       []string `yaml:"cnis"`
}

// Config
type Config struct {
	SkipValidate   bool            `yaml:"skip_validate"`
	Image          *Image          `yaml:"image"`
	Charts         *Charts         `yaml:"charts"`
	CloudProviders []CloudProvider `yaml:"cloud_providers"`
	Networking     *Networking     `yaml:"networking"`
}

// Validate
func (c *Config) Validate() error {
	return nil
}

// Load
func Load(configFile string) (*Config, error) {
	b, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	var c Config
	if err := yaml.Unmarshal(b, &c); err != nil {
		return nil, err
	}

	return &c, nil
}
