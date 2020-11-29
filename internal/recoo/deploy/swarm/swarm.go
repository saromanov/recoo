package swarm

import (
	"fmt"
	"os/exec"

	"github.com/saromanov/recoo/internal/config"
)

// Run defines execute of swarm stage
func Run(cfg config.Deploy, imageURL, imageName string) error {
	if err := generateCompose(cfg, imageURL, imageName); err != nil {
		return err
	}
	cmdStr := "docker stack deploy --compose-file docker-compose.yml recoo --with-registry-auth"
	_, err := exec.Command("/bin/bash", "-c", cmdStr).Output()
	if err != nil {
		return fmt.Errorf("unable to exec command: %v", err)
	}
	return nil
}
