package client

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/leonardogdasilva/pkg/config/v1alpha1"
	"github.com/pismo/e2eclient/config"
	"github.com/pismo/e2eclient/starlark"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type ArgMap map[string]string

const (
	COMPONENT_KEY = "component_name"
)

func execute(name string, source io.Reader, args ArgMap) error {
	star := starlark.New()

	if args != nil {
		starStruct, err := starlark.NewGoValue(args).ToStarlarkStruct("args")
		if err != nil {
			return err
		}

		star.AddPredeclared("args", starStruct)
	}

	err := star.Exec(name, source)
	if err != nil {
		err = errors.Wrap(err, "exec failed")
	}

	return err
}

func processTaskList(defs *[]v1alpha1.ComponentDef, args *ArgMap, basepath string) error {
	for _, comp := range *defs {
		compName := *comp.Name
		tasks := &comp.Tasks
		for _, task := range *tasks {
			processTaskParams(args, task.Params)
			taskName := *task.Name
			taskScript := *task.Script
			injectComponent(compName, args)
			logrus.Infof("Executing component: %s, task: %s, script: %s ", compName, taskName, taskScript)
			err := loadFile(basepath, taskScript, *args)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func ExecuteScenario(scn v1alpha1.SchemaJson, args ArgMap, basepath string) error {
	logrus.Infof("Executing scenario '%s'", *scn.Name)

	defs := &scn.Spec.Setup
	if err := processTaskList(defs, &args, basepath); err != nil {
		return err
	}

	// reads and injects client config information
	injectConfigParams(&args, config.ClientConfig)

	defs = &scn.Spec.Components
	if err := processTaskList(defs, &args, basepath); err != nil {

		return err
	}

	defs = &scn.Spec.Teardown
	if err := processTaskList(defs, &args, basepath); err != nil {
		return err
	}

	return nil
}

func injectComponent(compName string, args *ArgMap) {
	m := map[string]string(*args)
	m[COMPONENT_KEY] = compName
}

func injectConfigParams(args *ArgMap, cfg *config.Config) error {
	m := map[string]string(*args)
	for _, c := range cfg.Components {
		n := c.Name
		kworkdir := fmt.Sprintf("%s_workdir", n)
		kcompdir := fmt.Sprintf("%s_componentdir", n)
		kgit := fmt.Sprintf("%s_git", n)
		m[strings.TrimSpace(kworkdir)] = filepath.Join(config.CliConfigTargetDir, c.WorkDir)
		if c.ComponentDir != "" {
			m[strings.TrimSpace(kcompdir)] = filepath.Join(config.CliConfigTargetDir, c.WorkDir, c.ComponentDir)
		}
		if c.Git != "" {
			m[strings.TrimSpace(kgit)] = c.Git
		}
	}
	return nil
}

func processTaskParams(args *ArgMap, params v1alpha1.ParamsDef) {
	m := map[string]string(*args)
	for _, param := range params {
		k := *param.Name
		v := *param.Value
		m[strings.TrimSpace(k)] = strings.TrimSpace(v)
	}
}

func loadFile(basepath string, script string, args ArgMap) error {
	file, err := os.Open(filepath.Join(basepath, script))
	if err != nil {
		return errors.Wrapf(err, "failed to open script file: %s", script)
	}
	defer file.Close()
	return executeFile(file, args)
}

func executeFile(file *os.File, args ArgMap) error {
	return execute(file.Name(), file, args)
}
