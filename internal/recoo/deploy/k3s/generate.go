package k3s

import (
	"fmt"
	"io/ioutil"

	"github.com/go-yaml/yaml"
	"github.com/saromanov/recoo/internal/config"
)

type K3S struct {
	cfg config.Deploy
}

// Kuber defines struct for generation k3s
// (kubernetes) file
type Kuber struct {
	APIVersion string            `yaml:"apiVersion"`
	Kind       string            `yaml:"kind"`
	Metadata   map[string]string `yaml:"metadata"`
	Spec       Spec              `yaml:"spec"`
}

type Spec struct {
	Replicas uint
	Selector Selector `yaml:"selector"`
}

type Selector struct {
	MatchLabels MatchLabels `yaml:"matchLabels"`
}

type MatchLabels struct {
	App string `yaml:"app"`
}

func New(cfg config.Deploy)*K3S {
	return &K3S{
		cfg: cfg,
	}
}


func (k *K3S) Run(imageURL, imageName string, ports []string) error {
	return nil
}

func generateK3S(cfg config.Deploy, imageURL, imageName string, ports []string) error {
	k := &Kuber{
		APIVersion: "extensions/v1beta1",
		Kind:       "Deployment",
		Metadata: map[string]string{
			"name":      "recoo",
			"namespace": "default",
		},
		Spec: Spec {
			Replicas: 1,
		},
	}
	out, err := yaml.Marshal(k)
	if err != nil {
		return fmt.Errorf("unable to marshal to file: %v", err)
	}
	if err := ioutil.WriteFile("docker-compose.yml", out, 0644); err != nil {
		return fmt.Errorf("unable to write to file: %v", err)
	}
	return nil
}
