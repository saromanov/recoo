package build

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/saromanov/recoo/internal/config"
)

func createDockerfile(cfg config.Build, lang Language) error {
	data := generateDockerfile(cfg)
	if err := ioutil.WriteFile("Dockerfile", []byte(data), 0644); err != nil {
		return fmt.Errorf("unable to write file: %v", err)
	}

	if err := archiveBuildContext(); err != nil {
		return fmt.Errorf("unable to archive build context: %v", err)
	}
	if err := buildImage([]string{"1.0"}, "recoo.tar.gz"); err != nil {
		return fmt.Errorf("unable to build image: %v", err)
	}
	if err := os.Remove("Dockerfile"); err != nil {
		return fmt.Errorf("unable to remove Dockerfile: %v", err)
	}
	return nil
}

// archiveContext provides arhciving of directory for build context
// output should be name.tar.gz
func archiveBuildContext() error {
	return nil
}

func buildImage(tags []string, buildContext string) error {
	client, err := client.NewEnvClient()
	if err != nil {
		return fmt.Errorf("unable to init Docker client")
	}

	ctx := context.Background()

	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()

	buildContextReader, err := os.Open(buildContext)
	if err != nil {
		return fmt.Errorf("unable to open build context file: %v", err)
	}

	buildOptions := types.ImageBuildOptions{
		Context:    buildContextReader,
		Dockerfile: "Dockerfile",
		Remove:     true,
		Tags:       tags,
		NoCache:    true,
	}

	imageBuildResponse, err := client.ImageBuild(
		ctx,
		buildContextReader,
		buildOptions,
	)

	if err != nil {
		return fmt.Errorf("unable to build Dockerfile: %v", err)
	}

	defer imageBuildResponse.Body.Close()
	_, err = io.Copy(os.Stdout, imageBuildResponse.Body)
	if err != nil {
		return fmt.Errorf("unable to apply copy: %v", err)
	}

	return nil
}

// generateDockerfile provides generating of Dockerfiloe based on language
// https://semaphoreci.com/community/tutorials/how-to-deploy-a-go-web-application-with-docker
func generateDockerfile(cfg config.Build) string {
	data := fmt.Sprintf("FROM %s\n", cfg.Image)
	data += "RUN ls -la\n"
	data += fmt.Sprintf("ADD . /app\n")
	data += "WORKDIR /app\n"
	data += "RUN ls -la\n"
	data += fmt.Sprintf("RUN go mod download\n")
	data += fmt.Sprintf("RUN go build -o /bin/app %s\n", cfg.Entryfile)
	data += "ENTRYPOINT [ /bin/app ]"
	return data
}
