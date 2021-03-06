package build

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/saromanov/recoo/internal/config"
)

// https://medium.com/@Frikkylikeme/controlling-docker-with-golang-code-b213d9699998

func createDockerfile(cfg config.Build, lang Language, artifactsPath, namespace, dirName string) error {
	data := generateDockerfile(cfg, getImage(cfg, lang))
	if len(data) == 0 {
		return fmt.Errorf("failed to generate Dockerfile")
	}
	if err := createModules(); err != nil {
		return fmt.Errorf("unable to create modules: %v", err)
	}
	dockerfile := filepath.Join(artifactsPath, fmt.Sprintf("%s.recoo.Dockerfile", dirName))
	if err := ioutil.WriteFile(dockerfile, []byte(data), 0644); err != nil {
		return fmt.Errorf("unable to write file: %v", err)
	}
	if err := archiveBuildContext(dirName); err != nil {
		return failCreateDockerfile(fmt.Errorf("unable to archive build context: %v", err), dockerfile)
	}
	archfile := filepath.Join(artifactsPath, fmt.Sprintf("%s.tar.gzip", dirName))
	if err := buildImage([]string{fmt.Sprintf("%s/%s", namespace, dirName)}, dockerfile, archfile); err != nil {
		return failCreateDockerfile(fmt.Errorf("unable to build image: %v", err), archfile, dockerfile)
	}
	if err := os.Remove(dockerfile); err != nil {
		return fmt.Errorf("unable to remove Dockerfile: %v", err)
	}
	if err := os.Remove(fmt.Sprintf("%s.tar.gzip", dirName)); err != nil {
		return fmt.Errorf("unable to remove archive: %v", err)
	}
	return nil
}

// archiveContext provides arhciving of directory for build context
// output should be name.tar.gz
func archiveBuildContext(archiveName string) error {
	if err := compress(".", archiveName); err != nil {
		return fmt.Errorf("unable to compress file: %v", err)
	}
	return nil
}

func buildImage(tags []string, imageName, buildContext string) error {
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
		Dockerfile: imageName,
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
func generateDockerfile(cfg config.Build, image string) string {
	data := fmt.Sprintf("FROM %s as builder\n", image)
	data += fmt.Sprintf("ADD . /app\n")
	data += "WORKDIR /app\n"
	if len(cfg.Install) > 0 {
		data += fmt.Sprintf("RUN apk add --upgrade && apk add %s\n", strings.Join(cfg.Install, " "))
	}
	data += "ENV CGO_ENABLED=0\n"
	data += "ENV GOOS=linux\n"
	data += "ENV GARCH=amd64\n"
	data += fmt.Sprintf("RUN go mod download\n")
	data += fmt.Sprintf("RUN go build -o /bin/app %s\n", cfg.Entryfile)

	data += `FROM scratch
		COPY --from=builder /bin/app .
		CMD ["./app"]`
	return data
}

func getImage(cfg config.Build, lang Language) string {
	if cfg.Image != "" {
		return cfg.Image
	}
	if lang == GO {
		return "golang:1.15-alpine"
	}
	if lang == Python {
		return "python3.9-slim"
	}
	return ""
}

// removing current artifacts on files
func failCreateDockerfile(err error, files ...string) error {
	for _, f := range files {
		os.Remove(f) // skip checking of the error
	}
	return err
}
