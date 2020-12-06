package build

import (
	"fmt"
	"os"
	"os/exec"
)

// create modules probides creating of modules for go
func createModules() error {
	_, errMod := os.Stat("go.mod")
	if errMod == nil {
		return nil
	}

	_, err := exec.Command("/bin/bash", "-c", "go mod init").Output()
	if err != nil {
		return fmt.Errorf("unable to exec command: %v", err)
	}

	_, err = exec.Command("/bin/bash", "-c", "go mod tidy").Output()
	if err != nil {
		return fmt.Errorf("unable to exec command: %v", err)
	}
	return nil
}