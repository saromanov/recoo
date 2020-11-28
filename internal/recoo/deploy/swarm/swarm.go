package swarm

import "github.com/saromanov/recoo/internal/config"

func Run(cfg config.Deploy, imageURL, imageName string) error {
	if err := generateCompose(cfg, imageURL, imageName); err != nil {
		return err
	}
	/*_, err := os.Exec("docker swarm deploy").Output()
	if err != nil {
		return fmt.Errorf("unable to exec: %v", err)
	}*/
	return nil
}
