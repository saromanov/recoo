package deploy

import "github.com/saromanov/recoo/internal/config"

type Deploy interface {
	Run(*config.Deploy) error
}
