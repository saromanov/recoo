package build

import (
	"fmt"

	"github.com/saromanov/recoo/internal/config"
)

// Run starts build phase
func Run(cfg config.Build, artifactsPath, namespace, imageName string) error {
	if cfg.Entryfile == "" {
		return fmt.Errorf("entryfile is not defined")
	}
	lang, err := detectLanguage(cfg.Entryfile)
	if err != nil {
		return fmt.Errorf("unable to detect language: %v", err)
	}
	if err := createDockerfile(cfg, lang, artifactsPath, namespace, imageName); err != nil {
		return fmt.Errorf("unable to create dockerfile: %v", err)
	}
	return nil
}
