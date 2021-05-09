package starlark

import (
	"fmt"

	"github.com/vladimirvivien/gexe"
	"go.starlark.net/starlark"
)

// runLocalFunc is a built-in starlark function that runs a provided command on the local machine.
// It returns the result of the command as struct containing information about the executed command.
// Starlark format: run_local(<command string>)
func runLocalFunc(thread *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var cmdStr string
	err := starlark.UnpackArgs(
		identifiers.runLocal, args, kwargs,
		"cmd", &cmdStr,
	)
	if err == nil {
		p := gexe.New().RunProc(cmdStr)
		if p.Err() != nil {
			return starlark.None, fmt.Errorf("%s: %s: %s", identifiers.runLocal, p.Err(), p.Result())
		} else {
			return starlark.String(p.Result()), nil
		}
	} else {
		return starlark.None, err
	}

}
