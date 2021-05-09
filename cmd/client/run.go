package client

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/leonardogdasilva/pkg/config"
	"github.com/leonardogdasilva/pkg/config/v1alpha1"
	"github.com/pismo/e2eclient/util"
	yaml2 "sigs.k8s.io/yaml"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type runFlags struct {
	args     map[string]string
	argsFile string
}

func defaultRunFlags() *runFlags {
	return &runFlags{
		argsFile: ArgsFile,
	}
}

func RunScenario() *cobra.Command {
	flags := defaultRunFlags()

	cmd := &cobra.Command{
		Args:  cobra.ExactArgs(1),
		Use:   "run <file-name>",
		Short: "runs a scenario file",
		Long:  "Parses and executes the specified scenario file",
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(flags, args[0])
		},
	}
	cmd.Flags().StringVar(&flags.argsFile, "args-file", flags.argsFile, "path to a file containing key=value argument pairs that are passed to the script file")
	return cmd
}

func run(flags *runFlags, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return errors.Wrapf(err, "failed to open scenario file: %s", path)
	}
	defer file.Close()

	if err = config.ValidateSchemaFile(path, []byte(v1alpha1.JSONSchema)); err != nil {
		return errors.Errorf("Validation of config file %s against the default schema contains errors", path)
	}

	var content []byte
	if content, err = ioutil.ReadFile(path); err != nil {
		return errors.Errorf("Invalid yaml loaded from %s", path)
	}
	content, _ = yaml2.YAMLToJSON(content)
	scenario := &v1alpha1.SchemaJson{}
	if err = json.Unmarshal(content, scenario); err != nil {
		return errors.Errorf("Invalid scenario configuration %s %#v", path, err)
	}

	scriptArgs, err := processScriptArguments(flags)
	if err != nil {
		return err
	}

	dir := filepath.Dir(path)

	if err := ExecuteScenario(*scenario, scriptArgs, dir); err != nil {
		return errors.Wrapf(err, "execution failed for %s", file.Name())
	}

	return nil
}

// prepares a map of key-value strings to be passed to the execution script
// It builds the map from the args-file as well as the args flag passed to
// the run command.
func processScriptArguments(flags *runFlags) (map[string]string, error) {
	scriptArgs := map[string]string{}

	// get args from script args file
	err := util.ReadArgsFile(flags.argsFile, scriptArgs)
	if err != nil && flags.argsFile != ArgsFile {
		return nil, errors.Wrapf(err, "failed to parse scriptArgs file: %s", flags.argsFile)
	}

	// any value specified by the args flag overrides
	// value with same key in the args-file
	for k, v := range flags.args {
		scriptArgs[strings.TrimSpace(k)] = strings.TrimSpace(v)
	}

	return scriptArgs, nil
}
