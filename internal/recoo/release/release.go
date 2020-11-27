package release

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"

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

	cfg.Registry.Login = getVariable(cfg.Registry.Login)
	cfg.Registry.Password = getVariable(cfg.Registry.Password)
	return imagePush(cfg, image)
}

// imagePush provides pushing of images
func imagePush(cfg config.Release, image string) error {
	ctx := context.Background()

	cli, err := client.NewEnvClient()
	if err != nil {
		return fmt.Errorf("unabel to make new env client: %v", err)
	}

	err = cli.ImageTag(ctx, image, "docker.hub.com/recoo")
	if err != nil {
		return fmt.Errorf("unable to make image tag: %v", err)
	}

	authConfig := types.AuthConfig{
		Username: cfg.Registry.Login,
		Password: cfg.Registry.Password,
	}
	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		return fmt.Errorf("unable to marshal auth config: %v", err)
	}
	authStr := base64.URLEncoding.EncodeToString(encodedJSON)

	_, err = cli.ImagePush(ctx, fmt.Sprintf("%s/%s/%s", cfg.Registry.URL, "motorcode", image), types.ImagePushOptions{
		RegistryAuth: authStr,
	})
	if err != nil {
		return fmt.Errorf("unable to push image: %v", err)
	}
	return nil

}

func getVariable(data string) string {
	if strings.HasPrefix(data, "$") {
		return os.Getenv(data)
	}
	return data
}
