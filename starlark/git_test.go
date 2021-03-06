package starlark

import (
	"path/filepath"
	"testing"

	"github.com/leonardogdasilva/util"
	"go.starlark.net/starlark"
)

func TestGitFunc(t *testing.T) {
	tests := []struct {
		name     string
		args     func(t *testing.T) []starlark.Tuple
		eval     func(t *testing.T, kwargs []starlark.Tuple)
		teardown func(t *testing.T)
	}{
		{
			name: "simple clone",
			args: func(t *testing.T) []starlark.Tuple {
				return []starlark.Tuple{
					{starlark.String("repo"), starlark.String("https://github.com/git-fixtures/basic.git")},
					{starlark.String("branch"), starlark.String("branch")},
					{starlark.String("targetDir"), starlark.String("/tmp/foo")},
				}
			},
			eval: func(t *testing.T, kwargs []starlark.Tuple) {
				val, err := gitFunc(newTestThreadLocal(t), nil, nil, kwargs)
				if err != nil {
					t.Fatal(err)
				}
				result := ""
				if r, ok := val.(starlark.String); ok {
					result = string(r)
				}
				if result != "branch" {
					t.Errorf("unexpected result: %s", result)
				}
			},
			teardown: func(t *testing.T) {
				path, _ := filepath.Abs("/tmp/foo")
				if err := util.RemoveContents(path); err != nil {
					t.Errorf("error removing folder /tmp/foo %s", err)
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.eval(t, test.args(t))
			test.teardown(t)
		})
	}
}
