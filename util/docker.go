package util

import (
	"fmt"
)

func DockerBuild(workdir string, registry string, image string, dockerfile string) (string, error) {
	cmd := fmt.Sprintf(`/bin/bash -c "(cd %s && docker build -t %s/%s -f %s .)"`, workdir, registry, image, dockerfile)
	errorFunc := func(err error) error {
		imageName := fmt.Sprintf("%s/%s", registry, image)
		return fmt.Errorf("error build container '%s' using '%s'", imageName, err.Error())
	}
	return run(cmd, errorFunc)
}

func DockerPush(registry string, image string) (string, error) {
	cmd := fmt.Sprintf(`/bin/bash -c "docker push %s/%s"`, registry, image)
	errorFunc := func(err error) error {
		imageName := fmt.Sprintf("%s/%s", registry, image)
		return fmt.Errorf("error pushing container '%s:%s' ", imageName, err.Error())
	}
	return run(cmd, errorFunc)
}
