package starlark

import (
	"os"

	"github.com/pismo/e2eclient/util"
	"go.starlark.net/starlark"
)

// gitFunc is a built-in starlark function that runs a git checkout on the local machine.
// It returns the result of the command as struct containing information about the executed command.
// Starlark format: git(<command string>)
func gitFunc(thread *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var repo string
	var branch string
	var targetDir string
	err := starlark.UnpackArgs(
		identifiers.git, args, kwargs,
		"repo", &repo, "targetDir", &targetDir, "branch?", &branch,
	)
	var out string
	if err == nil {
		// directory doesnt exists
		if _, err := os.Stat(targetDir); os.IsNotExist(err) {
			err := util.GitClone(repo, targetDir)
			if err != nil {
				return starlark.None, err
			}
		}
		err := util.GitClean(repo, targetDir)
		if err != nil {
			return starlark.None, err
		}
		err = util.GitCheckout(repo, targetDir, branch)
		if err != nil {
			return starlark.None, err
		}
		out, err = util.GitBranch(targetDir)
		if err != nil {
			return starlark.None, err
		}
	}
	return starlark.String(out), nil
}
