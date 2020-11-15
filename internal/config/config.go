package config

import (
	"fmt"

	"github.com/saromanov/cowrow"
)

// Config defines configuration
type Config struct {
	Build  Build  `yaml:"build"`
	Deploy Deploy `yaml:"deploy"`
}

// Build defined build stage
type Build struct {
	Image     string `yaml:"image"`
	Entryfile string `yaml:"entryfile"`
}

// Deploy defines stage for deploy
type Deploy struct {
	Provider string `yaml:"provider"`
}

// Load provides loading of config
func Load(path string) (*Config, error) {
	cfg := &Config{}
	if err := cowrow.LoadByPath(path, &cfg); err != nil {
		return nil, fmt.Errorf("unable to load config: %v", err)
	}
	return cfg, nil
}
