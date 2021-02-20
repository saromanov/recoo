package swarm

import (
	"fmt"
	"os/exec"

	"github.com/saromanov/recoo/internal/config"
)

// Run defines execute of swarm stage
func Run(cfg config.Deploy, imageURL, imageName string, ports []string) error {
	if err := generateCompose(cfg, imageURL, imageName, ports); err != nil {
		return err
	}
	cmdStr := fmt.Sprintf("docker stack deploy --compose-file docker-compose.yml recoo_%s --with-registry-auth", imageName)
	_, err := exec.Command("/bin/bash", "-c", cmdStr).Output()
	if err != nil {
		return fmt.Errorf("unable to exec command: %v", err)
	}
	return nil
}
