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
	return nil
}

func buildImage(client *client.Client, tags []string, dockerfile string) error {
	ctx := context.Background()

	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()

	dockerFileReader, err := os.Open(dockerfile)
	if err != nil {
		return fmt.Errorf("unable to open Dockerfile: %v", err)
	}

	readDockerFile, err := ioutil.ReadAll(dockerFileReader)
	if err != nil {
		return fmt.Errorf("unable to read Dockerfile: %v", err)
	}

	tarHeader := &tar.Header{
		Name: dockerfile,
		Size: int64(len(readDockerFile)),
	}

	err = tw.WriteHeader(tarHeader)
	if err != nil {
		return err
	}

	_, err = tw.Write(readDockerFile)
	if err != nil {
		return err
	}

	dockerFileTarReader := bytes.NewReader(buf.Bytes())

	buildOptions := types.ImageBuildOptions{
		Context:    dockerFileTarReader,
		Dockerfile: dockerfile,
		Remove:     true,
		Tags:       tags,
	}

	imageBuildResponse, err := client.ImageBuild(
		ctx,
		dockerFileTarReader,
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
	data += fmt.Sprintf("ADD . /app\n")
	data += "WORKDIR /app\n"
	data += "RUN ls -la\n"
	data += fmt.Sprintf("RUN go mod download\n")
	data += fmt.Sprintf("RUN go build -o /bin/app %s\n", cfg.Entryfile)
	data += "ENTRYPOINT [ /bin/app ]"
	return data
}
