package starlark

import (
	"path/filepath"
	"strings"
	"testing"

	"go.starlark.net/starlark"
)

func TestContainerFunc(t *testing.T) {
	tests := []struct {
		name   string
		thread func(t *testing.T) *starlark.Thread
		args   func(t *testing.T, thread *starlark.Thread) []starlark.Tuple
		eval   func(t *testing.T, kwargs []starlark.Tuple, thread *starlark.Thread)
	}{
		{
			name: "simple build",
			thread: func(t *testing.T) *starlark.Thread {
				return newTestThreadLocal(t)
			},
			args: func(t *testing.T, thread *starlark.Thread) []starlark.Tuple {
				workdir := filepath.Dir(getTestUnitFilename(t))
				return []starlark.Tuple{
					{starlark.String("workdir"), starlark.String(workdir)},
					{starlark.String("registry"), starlark.String("docker.io")},
					{starlark.String("image"), starlark.String("simple-build")},
					{starlark.String("dockerfile"), starlark.String("test/Dockerfile")},
				}
			},
			eval: func(t *testing.T, kwargs []starlark.Tuple, thread *starlark.Thread) {
				_, err := containerFunc(thread, nil, nil, kwargs)
				if err == nil {
					if !strings.Contains(err.Error(), "error pushing container") {
						t.Fatal(err)
					}
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			thread := test.thread(t)
			test.eval(t, test.args(t, thread), thread)
		})
	}
}
