package release

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/saromanov/recoo/internal/config"
)

// Run starts of execution of release pipeline
func Run(cfg config.Release, image string) error {
	if cfg.Registry.Login == "" && cfg.Registry.Password == "" {
		return fmt.Errorf("login or password is not defined")
	}
	if cfg.Registry.URL == "" {
		return fmt.Errorf("url to the registry is not defined")
	}

	cfg.Registry.Login = readEnv(cfg.Registry.Login)
	cfg.Registry.Password = readEnv(cfg.Registry.Password)
	return imagePush(cfg, image)
}

func imagePush(cfg config.Release, image string) error {
	ctx := context.Background()

	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	err = cli.ImageTag(ctx, "alpine:3", "docker.io/jerrymannel/alpine:3")
	if err != nil {
		return err
	}

	authConfig := types.AuthConfig{
		Username: cfg.Registry.Login,
		Password: cfg.Registry.Password,
	}
	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		return err
	}
	authStr := base64.URLEncoding.EncodeToString(encodedJSON)

	_, err = cli.ImagePush(ctx, fmt.Sprintf("%s/%s", cfg.Registry.URL, image), types.ImagePushOptions{
		RegistryAuth: authStr,
	})
	if err != nil {
		return err
	}
	return nil

}
func readEnv(data string) string {
	return os.Getenv(data)
}
