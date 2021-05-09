package client

import (
	"github.com/pismo/e2eclient/buildinfo"
	"github.com/pismo/e2eclient/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	defaultLogLevel = logrus.InfoLevel
	CliName         = "e2ecli"
)

// globalFlags flags for the command
type globalFlags struct {
	debug bool
}

func preRun(flags *globalFlags) error {
	level := defaultLogLevel
	if flags.debug {
		level = logrus.DebugLevel
	}
	logrus.SetLevel(level)

	return CreateClidDir()
}

// clientDiagnosticsCommand creates a main cli command
func clientDiagnosticsCommand() *cobra.Command {
	flags := &globalFlags{debug: false}
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   CliName,
		Short: "runs the e2ecli program",
		Long:  "Runs the e2ecli program to execute script that interacts with e2e environment",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return preRun(flags)
		},
		SilenceUsage: true,
		Version:      buildinfo.Version,
	}

	cmd.PersistentFlags().BoolVar(
		&flags.debug,
		"debug",
		flags.debug,
		"sets log level to debug",
	)

	// cmd.AddCommand(newRunCommand())
	cmd.AddCommand(newBuildinfoCommand())
	cmd.AddCommand(InitConfig(config.CliDir))
	cmd.AddCommand(RunScenario())
	return cmd
}

func Run() error {
	return clientDiagnosticsCommand().Execute()
}
