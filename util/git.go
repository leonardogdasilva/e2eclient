package util

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/vladimirvivien/gexe"
)

func gitCommand(cmd string, err func(error) error) (string, error) {
	p := gexe.New().RunProc(cmd)
	if p.Err() != nil {
		return "", err(p.Err())
	} else {
		logrus.Info(p.Result())
	}
	return p.Result(), nil
}

func GitClone(repo string, targetDir string) error {
	cmd := fmt.Sprintf(`/bin/bash -c "git clone %s %s"`, repo, targetDir)
	errorFunc := func(err error) error {
		return fmt.Errorf("error cloning git repo %s into %s: %s", repo, targetDir, err.Error())
	}
	_, err := gitCommand(cmd, errorFunc)
	return err
}

func GitClean(repo string, targetDir string) error {
	cmd := fmt.Sprintf(`/bin/bash -c "(cd %s && git clean -fdx)"`, targetDir)
	errorFunc := func(err error) error {
		return fmt.Errorf("error cleaning git repo %s into %s: %s", repo, targetDir, err.Error())
	}
	_, err := gitCommand(cmd, errorFunc)
	return err
}

func GitCheckout(repo string, targetDir string, branch string) error {
	cmd := fmt.Sprintf(`/bin/bash -c "(cd %s && git checkout %s)"`, targetDir, branch)
	errorFunc := func(err error) error {
		return fmt.Errorf("error checking out repo %s into %s: %s", repo, targetDir, err.Error())
	}
	_, err := gitCommand(cmd, errorFunc)
	return err
}

func GitBranch(targetDir string) (string, error) {
	cmd := fmt.Sprintf(`/bin/bash -c "(cd %s && git rev-parse --abbrev-ref HEAD)"`, targetDir)
	errorFunc := func(err error) error {
		return fmt.Errorf("error retrieving branch name from repo %s: %s", targetDir, err.Error())
	}
	return gitCommand(cmd, errorFunc)
}
