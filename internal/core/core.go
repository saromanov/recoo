package core

import (
	"context"
	"fmt"

	"github.com/saromanov/recoo/internal/recoo/build"
	"github.com/saromanov/recoo/internal/config"
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
	if err := build.Run(c.cfg.Build); err != nil {
		return fmt.Errorf("unable to execute build phase: %v", err)
	}
	return nil
}
