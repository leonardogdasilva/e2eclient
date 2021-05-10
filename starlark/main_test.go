package starlark

import (
	"runtime"
	"testing"

	"go.starlark.net/starlark"
)

func newTestThreadLocal(t *testing.T) *starlark.Thread {
	thread := &starlark.Thread{Name: "test-e2e"}
	if err := setupLocalDefaults(thread); err != nil {
		t.Fatalf("failed to setup new thread local: %s", err)
	}
	return thread
}

func getTestUnitFilename(t *testing.T) string {
	_, filename, _, _ := runtime.Caller(0)
	return filename
}
