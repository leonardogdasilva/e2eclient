package starlark

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"go.starlark.net/starlark"
	"go.starlark.net/starlarkstruct"
)

var (
	identifiers = struct {
		os        string
		cliDir    string
		cliCfg    string
		scriptCtx string
		runLocal  string
		git       string
	}{
		os:        "os",
		cliDir:    filepath.Join(os.Getenv("HOME"), ".e2ecli"),
		cliCfg:    "cli_config",
		scriptCtx: "script_context",
		runLocal:  "run_local",
		git:       "git",
	}

	defaults = struct {
		workdir string
	}{
		workdir: "/tmp/e2ecli",
	}
)

func getWorkdirFromThread(thread *starlark.Thread) (string, error) {
	val := thread.Local(identifiers.cliCfg)
	if val == nil {
		return "", fmt.Errorf("%s not found in threard", identifiers.cliCfg)
	}
	var result string
	if valStruct, ok := val.(*starlarkstruct.Struct); ok {
		if valStr, err := valStruct.Attr("workdir"); err == nil {
			if str, ok := valStr.(starlark.String); ok {
				result = string(str)
			}
		}
	}

	if len(result) == 0 {
		result = defaults.workdir
	}
	return result, nil
}

func getUsername() string {
	usr, err := user.Current()
	if err != nil {
		return ""
	}
	return usr.Username
}
