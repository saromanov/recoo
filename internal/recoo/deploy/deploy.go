package deploy

import (
	"github.com/saromanov/recoo/internal/config"
	"github.com/saromanov/recoo/internal/deploy/k3s"
)
type Deploy interface {
	Run(*config.Deploy) error
}

type DeployFactory struct {

}

// Run provides running of the
func (d *DeployFactory) Run(dep *config.Deploy) (Deploy, error) {
	if dep.Provider == config.K3S {
		return k3s.New(), nil
	}
	return nil, nil
}