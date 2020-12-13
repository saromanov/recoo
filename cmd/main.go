package main

import (
	"context"
	"github.com/saromanov/recoo/internal/config"
	"github.com/saromanov/recoo/internal/core"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
)

func exec() {
	app := &cli.App{
		Name:  "recoo",
		Usage: "Starting of the app",
		Flags: []cli.Flag{},
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
	cfg, err := config.Load("recoo-config.yml")
	if err != nil {
		logrus.WithError(err).Fatalf("unable to load config")
	}
	if cfg == nil {
		logrus.Fatalf("unable to load config")
	}

	c := core.New(cfg)
	if err := c.Start(context.Background()); err != nil {
		logrus.WithError(err).Fatalf("unable to execute pipeline")
	}
	return nil
}

func stop(ctx *cli.Context) error {
	cfg, err := config.Load("recoo-config.yml")
	if err != nil {
		logrus.WithError(err).Fatalf("unable to load config")
	}
	c := core.New(cfg)
	if err := c.Remove(context.Background()); err != nil {
		logrus.WithError(err).Fatalf("unable to remove pipeline")
	}
	return nil
}

func main() {
	exec()
}
