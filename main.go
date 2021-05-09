package main

import (
	"os"
	"path/filepath"

	"github.com/pismo/e2eclient/cmd/client"
	"github.com/pismo/e2eclient/config"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(os.Stdout)
	logrus.Info("Starting e2e client...")

	_, err := config.LoadConfig(filepath.Join(config.CliDir, config.DefaultConfig))
	if err != nil {
		logrus.Warnf("error on reading config from '%s': '%s'", config.CliDir, err.Error())
		logrus.Infof("Please execute the init command")
	} else {
		logrus.Infof("Processing config from '%s:%s'", config.ClientConfig.Repository.Git,
			config.ClientConfig.Repository.Branch)
		config.UpdateConfigDir()
	}
}

func main() {

	if err := client.Run(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}
