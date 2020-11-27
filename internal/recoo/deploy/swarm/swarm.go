package swarm

func Run() error {
	if err := generateCompose(); err != nil {
		return err
	}
	/*_, err := os.Exec("docker swarm deploy").Output()
	if err != nil {
		return fmt.Errorf("unable to exec: %v", err)
	}*/
	return nil
}
