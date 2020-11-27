package swarm

import (
	"fmt"
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

// Compose defines file for saving data
type Compose struct {
	Version  string   `yaml:"version"`
	Networks []string `yaml:"networks"`
}

func generateCompose() error {
	c := &Compose{
		Version:  "'3.8'",
		Networks: []string{"test"},
	}

	out, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Errorf("unable to marshal to file: %v", err)
	}
	if err := ioutil.WriteFile("docker-compose.yml", out, 644); err != nil {
		return fmt.Errorf("unable to write to file: %v", err)
	}
	return nil
}
