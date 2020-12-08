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

// Network provides definition of configuration
type Network struct {
}

// Service provides definition of service configurarion
type Service struct {
	Image    string   `yaml:"image"`
	Networks []string `yaml:"networks"`
	Ports    []string `yaml:"ports"`
}

func generateCompose(cfg config.Deploy, imageURL, imageName string, ports []string) error {
	networkName := "test-recoo"
	c := &Compose{
		Version:  "3.3",
		Networks: map[string]Network{networkName: Network{}},
	}
	services := map[string]Service{}
	for _, s := range cfg.Services {
		services[fmt.Sprintf("%s-service", s.Image)] = Service{
			Image:    s.Image,
			Networks: []string{networkName},
		}
	}

	services[fmt.Sprintf("%s-service", imageName)] = Service{
		Image:    imageURL,
		Networks: []string{networkName},
		Ports: ports,
	}
	c.Services = services

	out, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Errorf("unable to marshal to file: %v", err)
	}
	if err := ioutil.WriteFile("docker-compose.yml", out, 0644); err != nil {
		return fmt.Errorf("unable to write to file: %v", err)
	}
	return nil
}