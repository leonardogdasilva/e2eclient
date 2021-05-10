package starlark

import (
	"bytes"

	"github.com/leonardogdasilva/util"
	"go.starlark.net/starlark"
)

// containerFunc is a built-in starlark function that runs a git checkout on the local machine.
// It returns the result of the command as struct containing information about the executed command.
// Starlark format: git(<command string>)
func containerFunc(thread *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var registry string
	var image string
	var dockerfile string
	var workdir string
	err := starlark.UnpackArgs(
		identifiers.git, args, kwargs,
		"registry", &registry, "image", &image, "dockerfile", &dockerfile, "workdir", &workdir,
	)
	var out bytes.Buffer
	if err == nil {
		var s string
		if s, err = util.DockerBuild(workdir, registry, image, dockerfile); err != nil {
			return starlark.None, err
		} else {
			out.WriteString(s)
		}
		if s, err = util.DockerPush(registry, image); err != nil {
			return starlark.None, err
		} else {
			out.WriteString(s)
		}

	}
	return starlark.String(out.String()), nil
}
