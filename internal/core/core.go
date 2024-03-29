package core

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/saromanov/recoo/internal/config"
	"github.com/saromanov/recoo/internal/recoo/build"
	"github.com/saromanov/recoo/internal/recoo/deploy/swarm"
	"github.com/saromanov/recoo/internal/recoo/deploy"
	"github.com/saromanov/recoo/internal/recoo/release"
)

var errNoDirName = errors.New("unable to get dir name")

// Core defines main logic
type Core struct {
	cfg *config.Config
	dep deploy.Deploy
}

// New provides initalization of Core
func New(cfg *config.Config, dep deploy.Deploy) *Core {
	return &Core{
		cfg: cfg,
		dep: dep,
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
		return errNoDirName
	}
	if c.cfg.Build.Title != "" {
		dirName = c.cfg.Build.Title
	}
	imageURL := c.getImageURL(dirName)
	if err := c.preStage(); err != nil {
		return fmt.Errorf("unable to execute pre stage: %v", err)
	}
	fmt.Println("Executing of the build stage")
	if err := build.Run(c.cfg.Build, c.cfg.ArtifactsDir, c.cfg.Release.Registry.Login, dirName); err != nil {
		return fmt.Errorf("unable to execute build stage: %v", err)
	}
	fmt.Println("Executing of the release stage")
	if err := release.Run(c.cfg.Release, dirName); err != nil {
		return fmt.Errorf("unable to execute release stage: %v", err)
	}
	fmt.Println("Executing of the deploy stage")
	if err := c.dep.Run(imageURL, dirName, c.cfg.Build.Ports); err != nil {
		return fmt.Errorf("unable to run deploy stage: %v", err)
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
