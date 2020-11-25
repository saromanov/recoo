package swarm

import (
	"fmt"
	"os"

	"github.com/saromanov/recoo/internal/deploy"
)

// Service defines deploy to docker swarm
type Service struct {
}

// New provides initialization on swarm
func New() deploy.Deploy {
	return &Service{}
}

func (s *Service) Run() error {
	_, err := os.Exec("docker swarm deploy").Output()
	if err != nil {
		return fmt.Errorf("unable to exec: %v", err)
	}
	return nil
}
