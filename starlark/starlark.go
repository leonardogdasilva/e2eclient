package starlark

import (
	"context"
	"errors"
	"fmt"
	"io"

	"go.starlark.net/starlark"
)

type Executor struct {
	thread  *starlark.Thread
	predecs starlark.StringDict
	result  starlark.StringDict
}

func New() *Executor {
	return &Executor{
		thread:  &starlark.Thread{Name: "e2eclient"},
		predecs: newPredeclareds(),
	}
}

// AddPredeclared predeclared
func (e *Executor) AddPredeclared(name string, value starlark.Value) {
	if e.predecs != nil {
		e.predecs[name] = value
	}
}

func (e *Executor) Exec(name string, source io.Reader) error {
	if err := setupLocalDefaults(e.thread); err != nil {
		return fmt.Errorf("failed to setup defaults: %s", err)
	}

	result, err := starlark.ExecFile(e.thread, name, source, e.predecs)
	if err != nil {
		if evalErr, ok := err.(*starlark.EvalError); ok {
			return fmt.Errorf(evalErr.Backtrace())
		}
		return err
	}
	e.result = result

	return nil
}

// setupLocalDefaults populates the provided execution thread
// with default configuration values.
func setupLocalDefaults(thread *starlark.Thread) error {
	if thread == nil {
		return errors.New("thread local is nil")
	}
	// add script context starlark thread
	ctx := context.Background()
	thread.SetLocal(identifiers.scriptCtx, ctx)

	return nil
}

// newPredeclareds creates string dictionary containing the
// global built-ins values and functions available to the
// running script.
func newPredeclareds() starlark.StringDict {
	return starlark.StringDict{
		identifiers.os:       setupOSStruct(),
		identifiers.runLocal: starlark.NewBuiltin(identifiers.runLocal, runLocalFunc),
		identifiers.git:      starlark.NewBuiltin(identifiers.git, gitFunc),
	}
}
