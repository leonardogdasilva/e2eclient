package config

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// example file
// repository:
//   git: git@github.com:pismo/api-gateway-e2e-test
//   branch: canary-api

// components:
// - name: auth-api
//   src: src/auth-api

var (
	// Directory path created at e2ecli runtime
	CliDir = filepath.Join(os.Getenv("HOME"), ".e2ecli")

	// Default config file
	DefaultConfig = "config.yaml"

	ClientConfig = &Config{}
)

type Component struct {
	Name string `yaml:"name"`
	Src  string `yaml:"src"`
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
