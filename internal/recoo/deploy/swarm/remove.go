package swarm

import (
	"fmt"
	"os/exec"
)

// Remove provides removing of service
func Remove(stackName string) error {
	cmdStr := fmt.Sprintf("docker stack rm recoo_%s", stackName)
	_, err := exec.Command("/bin/bash", "-c", cmdStr).Output()
	if err != nil {
		return fmt.Errorf("unable to exec command: %v", err)
	}
	return nil
}
