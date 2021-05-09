package config

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/leonardogdasilva/util"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// example file
// repository:
//   git: git@github.com:pismo/api-gateway-e2e-test
//   branch: canary-api

// components:
// - name: auth-api
//   workdir: src/auth-api
//   componentdir: auth-api

var (
	// Directory path created at e2ecli runtime
	CliDir = filepath.Join(os.Getenv("HOME"), ".e2ecli")

	// Directory to store e2e repo content
	CliConfigTargetDir = filepath.Join(CliDir, "e2e-test")

	// Default config file
	DefaultConfig = "config.yaml"

	ClientConfig = &Config{}
)

type Component struct {
	Name         string `yaml:"name"`
	WorkDir      string `yaml:"workdir"`
	ComponentDir string `yaml:"componentdir"`
	Git          string `yaml:"git"`
}

type Config struct {
	// Name corresponds to the yaml field "repository".
	Repository Repository `yaml:"repository"`
	// Name corresponds to the yaml field "components".
	Components []Component `yaml:"components"`
}

type Repository struct {
	Git    string `yaml:"git"`
	Branch string `yaml:"branch"`
}

func parseConfig(s io.Reader) error {
	d := yaml.NewDecoder(s)
	if err := d.Decode(ClientConfig); err != nil {
		return err
	}
	return nil
}

// NewConfig returns a new decoded struct
func LoadConfig(configPath string) (*Config, error) {

	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("invalid config path '%s'", configPath)
	}

	defer file.Close()
	parseConfig(file)

	return ClientConfig, nil
}

// UpdateConfigDir downloads git repository containing e2e metadata and supporting files
func UpdateConfigDir() error {
	if err := util.GitClone(ClientConfig.Repository.Git, CliConfigTargetDir); err != nil {
		return errors.Wrapf(err, "error downloading e2e config dir in '%s'", CliConfigTargetDir)
	}
	if err := util.GitCheckout(ClientConfig.Repository.Git, CliConfigTargetDir, ClientConfig.Repository.Branch); err != nil {
		return errors.Wrapf(err, "error updating e2e config dir in '%s'", CliConfigTargetDir)
	}
	return nil
}

func ValidateConfig(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}
