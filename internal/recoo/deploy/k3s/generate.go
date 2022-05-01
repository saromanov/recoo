package k3s

import (
	"fmt"
	"io/ioutil"

	"github.com/go-yaml/yaml"
	"github.com/saromanov/recoo/internal/config"
)

// Kuber defines struct for generation k3s
// (kubernetes) file
type Kuber struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
}

func generateK3S(cfg config.Deploy, imageURL, imageName string, ports []string) error {
	k := &Kuber{}
	out, err := yaml.Marshal(k)
	if err != nil {
		return fmt.Errorf("unable to marshal to file: %v", err)
	}
	if err := ioutil.WriteFile("docker-compose.yml", out, 0644); err != nil {
		return fmt.Errorf("unable to write to file: %v", err)
	}
	return nil
}
