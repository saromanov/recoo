package main

import (
	"context"
	"os"
	"github.com/saromanov/recoo/internal/config"
	"github.com/saromanov/recoo/internal/core"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func exec(){
app := &cli.App{
	Name:  "mystery",
	Usage: "Starting of the app",
	Flags: []cli.Flag{
		
	},
	Commands: []*cli.Command{
		{
			Name:   "run",
			Usage:  "running of pipeline",
			Action: run,
		},
	},
}

err := app.Run(os.Args)
if err != nil {
	return
}
}

func run(c *cli.Context) error {
	cfg, err := config.Load("config.yml")
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
func main() {
	exec()
}
