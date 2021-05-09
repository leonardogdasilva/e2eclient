package client

import (
	"os"
	"path/filepath"

	"github.com/pismo/e2eclient/config"
)

var (
	// file path of the defaults args file
	ArgsFile = filepath.Join(config.CliDir, "args")
)

// This creates a e2ecli directory which can be used as a default workdir
// for script execution. It will also house the default args file.
func CreateClidDir() error {
	if _, err := os.Stat(config.CliDir); os.IsNotExist(err) {
		return os.Mkdir(config.CliDir, 0755)
	}
	return nil
}
