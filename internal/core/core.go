package core

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/saromanov/recoo/internal/config"
	"github.com/saromanov/recoo/internal/recoo/build"
	"github.com/saromanov/recoo/internal/recoo/release"
	"github.com/saromanov/recoo/internal/recoo/deploy/swarm"
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
	if dirName == "" {
		return fmt.Errorf("unable to get dir name")
	}
	imageURL := c.getImageURL(dirName)
	if err := c.preStage(); err != nil {
		return fmt.Errorf("unable to execute pre stage: %v", err)
	}
	if err := build.Run(c.cfg.Build, dirName); err != nil {
		return fmt.Errorf("unable to execute build stage: %v", err)
	}
	if err := release.Run(c.cfg.Release, dirName); err != nil {
		return fmt.Errorf("unable to execute release stage: %v", err)
	}
	if err := swarm.Run(c.cfg.Deploy, imageURL, dirName); err != nil {
		return fmt.Errorf("unable to run swarm stage: %v", err)
	}
	return nil
}

func (c *Core) Remove(ctx context.Context) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("unable to get current dir: %v", err)
	}
	dirName := filepath.Base(currentDir)
	if dirName == "" {
		return fmt.Errorf("unable to get dir name")
	}
	
	if err := swarm.Remove(dirName); err != nil {
		return fmt.Errorf("unable to remove deploy stack: %v", err)
	}
	return nil
}

// getImageURL returns image url
func (c *Core) getImageURL(image string) string {
	if c.cfg.Release.Registry.URL == "local" {
		return fmt.Sprintf("%s/%s", c.cfg.Release.Registry.Login, image)
	}
	return fmt.Sprintf("%s/%s/%s", c.cfg.Release.Registry.URL, c.cfg.Release.Registry.Login, image)
}

// preStage probides running of prepare
func (c *Core) preStage() error {
	if _, err := os.Stat("recoo.Dockerfile"); err == nil {
		if err := os.Remove("recoo.Dockerfile"); err != nil {
			return fmt.Errorf("unable to remove file: %v", err)
		}
	} 
	if _, err := os.Stat("recoo.tar.gzip"); err == nil {
		if err := os.Remove("recoo.tar.gzip"); err != nil {
			return fmt.Errorf("unable to remove file: %v", err)
		}
	} 
	return nil
}
