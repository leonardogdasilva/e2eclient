package starlark

import (
	"testing"

	"go.starlark.net/starlark"
)

func newTestThreadLocal(t *testing.T) *starlark.Thread {
	thread := &starlark.Thread{Name: "test-crashd"}
	if err := setupLocalDefaults(thread); err != nil {
		t.Fatalf("failed to setup new thread local: %s", err)
	}
	return thread
}
