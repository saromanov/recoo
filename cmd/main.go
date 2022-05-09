package main

import (
	"context"
	"fmt"
	"os"

	"github.com/saromanov/recoo/internal/config"
	"github.com/saromanov/recoo/internal/core"
	"github.com/saromanov/recoo/internal/recoo/deploy"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func exec() {
	app := &cli.App{
		Name:  "recoo",
		Usage: "Starting of the app",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "path to config",
			},
		},
		Commands: []*cli.Command{
			{
				Name:   "run",
				Usage:  "running of pipeline",
				Action: run,
			},
			{
				Name:   "rm",
				Usage:  "stopping of pipeline and removing of services",
				Action: stop,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		return
	}
}

func run(ctx *cli.Context) error {
	cfg, err := getConfig(ctx)
	if err != nil {
		logrus.WithError(err).Fatalf("unable to load config")
	}
	if cfg == nil {
		logrus.Fatalf("unable to load config")
	}

	depFactory := deploy.DeployFactory{}
	dep, err := depFactory.Run(cfg.Deploy)
	if err != nil {
		logrus.WithError(err).Fatalf("unable to create deploy")
	}
	c := core.New(cfg, dep)
	if err := c.Start(context.Background()); err != nil {
		logrus.WithError(err).Fatalf("unable to execute pipeline")
	}
	return nil
}

func stop(ctx *cli.Context) error {
	cfg, err := getConfig(ctx)
	if err != nil {
		logrus.WithError(err).Fatalf("unable to load config")
	}
	depFactory := deploy.DeployFactory{}
	dep, err := depFactory.Run(cfg.Deploy)
	if err != nil {
		logrus.WithError(err).Fatalf("unable to create deploy")
	}
	c := core.New(cfg, dep)
	if err := c.Remove(context.Background()); err != nil {
		logrus.WithError(err).Fatalf("unable to remove pipeline")
	}
	return nil
}

func getConfig(ctx *cli.Context) (*config.Config, error) {
	cfgPath := ".recoo-config.yml"
	cfgFlags := ctx.String("config")
	if cfgFlags != "" {
		cfgPath = cfgFlags
	}
	cfg, err := config.Load(cfgPath)
	if err != nil {
		return nil, fmt.Errorf("unable to load config: %v", err)
	}
	return cfg, nil
}

func main() {
	exec()
}
