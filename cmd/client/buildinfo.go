package client

import (
	"fmt"

	"github.com/pismo/e2eclient/buildinfo"
	"github.com/spf13/cobra"
)

func newBuildinfoCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "version",
		Short: "prints version",
		Long:  "Prints version information for the program",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("Version: %s\nGitSHA: %s\n", buildinfo.Version, buildinfo.GitSHA)
			return nil
		},
	}
	return cmd
}
