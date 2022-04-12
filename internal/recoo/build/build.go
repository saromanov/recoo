package build

import (
	"errors"
	"fmt"

	"github.com/saromanov/recoo/internal/config"
)

var (
	errNoEntryfile = errors.New("entryfile is not defined")
)

// Run starts build phase
func Run(cfg config.Build, artifactsPath, namespace, imageName string) error {
	if cfg.Entry == "" {
		return errNoEntryfile
	}
	lang, err := detectLanguage(cfg.Entry)
	if err != nil {
		return fmt.Errorf("unable to detect language: %v", err)
	}
	if err := buildDockerfile(cfg, lang, artifactsPath, namespace, imageName); err != nil {
		return fmt.Errorf("unable to create dockerfile: %v", err)
	}
	return nil
}
