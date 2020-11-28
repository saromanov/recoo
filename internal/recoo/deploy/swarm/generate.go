package swarm

import (
	"fmt"
	"io/ioutil"

	"github.com/go-yaml/yaml"
	"github.com/saromanov/recoo/internal/config"
)

// Compose defines file for saving data
type Compose struct {
	Version  string             `yaml:"version"`
	Networks map[string]Network `yaml:"networks"`
	Services map[string]Service `yaml:"services"`
}

type Network struct {
}

type Service struct {
	Image    string   `yaml:"image"`
	Networks []string `yaml:"networks"`
}

func generateCompose(cfg config.Deploy, imageURL, imageName string) error {
	c := &Compose{
		Version:  "3.3",
		Networks: map[string]Network{"test": Network{}},
	}
	services := map[string]Service{}
	for _, s := range cfg.Services {
		services[fmt.Sprintf("%s-service", s.Image)] = Service{
			Image:    s.Image,
			Networks: []string{"test"},
		}
	}

	services[fmt.Sprintf("%s-service", imageName)] = Service{
		Image:    imageURL,
		Networks: []string{"test"},
	}
	c.Services = services

	out, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Errorf("unable to marshal to file: %v", err)
	}
	if err := ioutil.WriteFile("docker-compose.yml", out, 777); err != nil {
		return fmt.Errorf("unable to write to file: %v", err)
	}
	return nil
}
