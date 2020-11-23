package swarm

import "github.com/saromanov/recoo/internal/deploy"

// Service defines deploy to docker swarm
type Service struct {
}

// New provides initialization on swarm
func New() deploy.Deploy {
	return &Service{}
}
