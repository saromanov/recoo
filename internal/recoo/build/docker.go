package build

import (
	"fmt"
	"io/ioutil"

	"github.com/saromanov/recoo/internal/config"
)

func createDockerfile(cfg config.Build, lang Language) error {
	data := generateDockerfile(cfg)
	if err := ioutil.WriteFile("Dockerfile", []byte(data), 0644); err != nil {
		return fmt.Errorf("unable to write file: %v", err)
	}
	return nil
}

// generateDockerfile provides generating of Dockerfiloe based on language
func generateDockerfile(cfg config.Build) string {
	data := fmt.Sprintf("FROM %s", cfg.Image)
	data += fmt.Sprintf("go build -o app %s", cfg.Entryfile)
	data += "WORKDIR app"
	data += "CMD app"
	return data
}
