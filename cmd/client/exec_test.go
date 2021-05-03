package client

import (
	"strings"
	"testing"
)

func TestExecute(t *testing.T) {
	tests := []struct {
		name   string
		script string
		exec   func(t *testing.T, script string)
	}{
		{
			name:   "run_local",
			script: `result = run_local("echo 'Hello World!'")`,
			exec: func(t *testing.T, script string) {
				if err := execute("run_local", strings.NewReader(script), ArgMap{}); err != nil {
					t.Fatal(err)
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.exec(t, test.script)
		})
	}
}
