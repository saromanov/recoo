package config

import (
	"fmt"
	"strings"
	"os"

	"github.com/saromanov/cowrow"
)

type Provider int

const (
	Swarm Provider = iota
	K3S
)

// Config defines configuration
type Config struct {
	ArtifactsDir string  `yaml:"artifacts_dir"`
	Build        Build   `yaml:"build"`
	Deploy       Deploy  `yaml:"deploy"`
	Release      Release `yaml:"release"`
}

// Build defined build stage
type Build struct {
	Image   string   `yaml:"image"`
	Entry   string   `yaml:"entry"`
	Ports   []string `yaml:"ports"`
	Install []string `yaml:"install"`
	Title   string   `yaml:"title"`
	Tag     string   `yaml:"tag"`
}

// Deploy defines stage for deploy
type Deploy struct {
	Provider Provider    `yaml:"provider"`
	Services []Service `yaml:"services"`
}

// Service defines configuration for service
type Service struct {
	Image string `yaml:"image"`
}

// Release defines release stage
type Release struct {
	Registry Registry `yaml:"registry"`
}

// Registry defines configuration for registry
type Registry struct {
	Login    string `yaml:"login"`
	Password string `yaml:"password"`
	URL      string `yaml:"url"`
}

// Load provides loading of config
func Load(path string) (*Config, error) {
	cfg := &Config{}
	if err := cowrow.LoadByPath(path, &cfg); err != nil {
		return nil, fmt.Errorf("unable to load config file: %v", err)
	}
	return cfg, nil
}

// SetEnvVariables provides filling of config
// from environment variables
func SetEnvVariables(c *Config) {
	c.Build.Image = getVariable(c.Build.Image)
	c.Build.Tag = getVariable(c.Build.Tag)
}

func getVariable(data string) string {
	if strings.HasPrefix(data, "$") {
		return os.Getenv(data[1:])
	}
	return data
}

