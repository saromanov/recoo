package swarm

import "github.com/saromanov/recoo/internal/config"

func Run(cfg config.Deploy) error {
	if err := generateCompose(cfg); err != nil {
		return err
	}
	/*_, err := os.Exec("docker swarm deploy").Output()
	if err != nil {
		return fmt.Errorf("unable to exec: %v", err)
	}*/
	return nil
}
