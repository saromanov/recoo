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
	APIVersion string            `yaml:"apiVersion"`
	Kind       string            `yaml:"kind"`
	Metadata   map[string]string `yaml:"metadata"`
	Spec       Spec              `yaml:"spec"`
}

type Spec struct {
	Replicas uint
	Selector Selector `yaml:"selector"`
	Template Template `yaml:"template"`
}

type SpecContainers struct {
	Containers []Container `yaml:"containers"`
}

type Container struct {
	Name string `yaml:"name"`
	Image string `yaml:"image"`
}

type Template struct {
	Metadata Metadata `yaml:"metadata"`
	Spec SpecContainers `yaml:"spec"`
}

type Metadata struct {
	Labels Labels `yaml:"labels"`
}

type Labels struct {
	App string `yaml:"app"`
}
type Selector struct {
	MatchLabels MatchLabels `yaml:"matchLabels"`
}

type MatchLabels struct {
	App string `yaml:"app"`
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
			Replicas: cfg.Replicas,
			Selector: Selector{
				MatchLabels: MatchLabels{
					App: "k3s-demo",
				},
			},
			Template: Template{
				Metadata: Metadata{
					Labels: Labels {
						App: "k3s-demo",
					},
				},
				Spec: SpecContainers{
					Containers: []Container{
						{
							Name: imageName,
							Image: imageURL,
						},
					},
				},
			},

		},
	}
	out, err := yaml.Marshal(k)
	if err != nil {
		return fmt.Errorf("unable to marshal to file: %v", err)
	}
	if err := ioutil.WriteFile("k3s.yml", out, 0644); err != nil {
		return fmt.Errorf("unable to write to file: %v", err)
	}
	return nil
}
