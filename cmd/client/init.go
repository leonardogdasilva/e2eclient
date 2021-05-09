package client

import (
	_ "embed"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Embeds the current supported config
//go:embed e2e/config.yaml
var e2eClientConfig string

const configFilename = "config.yaml"

func InitConfig(folder string) *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "init",
		Short: "init local client",
		Long:  fmt.Sprintf("Initializes client local HOME folder %s", folder),
		RunE: func(cmd *cobra.Command, args []string) error {
			filename := filepath.Join(folder, filepath.Base(configFilename))
			if _, err := os.Stat(filename); !os.IsNotExist(err) {
				logrus.Warn(fmt.Sprintf("Config file %s already exists at %s. Ignoring...", configFilename, filename))
				return nil
			}
			err := ioutil.WriteFile(filename, []byte(e2eClientConfig), 0644)
			logrus.Info(fmt.Sprintf("Saving %s to %s", configFilename, filename))
			if err != nil {
				panic(err)
			}
			return nil
		},
	}
	return cmd
}
