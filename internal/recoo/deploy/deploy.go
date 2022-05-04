package deploy

import "github.com/saromanov/recoo/internal/config"

type Deploy interface {
	Run(*config.Deploy) error
}

type DeployFactory struct {

}

// Run provides running of the
func (d *DeployFactory) Run(dep Deploy) error {

}