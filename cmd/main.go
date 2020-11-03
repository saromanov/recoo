package main

import (
	"github.com/saromanov/recoo/internal/config"
	"github.com/saromanov/recoo/internal/core"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.Load("config.yml")
	if err != nil {
		logrus.WithError(err).Fatalf("unable to load config")
	}
	if cfg == nil {
		logrus.Fatalf("unable to load config")
	}

	core.New(cfg)

}
