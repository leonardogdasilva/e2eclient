package client

import (
	"os"
	"path/filepath"
)

var (
	// Directory path created at e2ecli runtime
	CliDir = filepath.Join(os.Getenv("HOME"), ".e2ecli")
	// file path of the defaults args file
	ArgsFile = filepath.Join(CliDir, "args")
)

// This creates a e2ecli directory which can be used as a default workdir
// for script execution. It will also house the default args file.
func CreateClidDir() error {
	if _, err := os.Stat(CliDir); os.IsNotExist(err) {
		return os.Mkdir(CliDir, 0755)
	}
	return nil
}
