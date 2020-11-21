package core

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/saromanov/recoo/internal/config"
	"github.com/saromanov/recoo/internal/recoo/build"
	"github.com/saromanov/recoo/internal/recoo/release"
)

// Core defines main logic
type Core struct {
	cfg *config.Config
}

// New provides initalization of Core
func New(cfg *config.Config) *Core {
	return &Core{
		cfg: cfg,
	}
}

// Start provides running of pipeline
func (c *Core) Start(ctx context.Context) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("unable to get current dir: %v", err)
	}
	dirName := filepath.Base(currentDir)

	if err := build.Run(c.cfg.Build, dirName); err != nil {
		return fmt.Errorf("unable to execute build phase: %v", err)
	}
	if err := release.Run(c.cfg.Release, dirName); err != nil {
		return fmt.Errorf("unable to execute release stage: %v", err)
	}
	return nil
}
