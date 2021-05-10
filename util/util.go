package util

import (
	"github.com/sirupsen/logrus"
	"github.com/vladimirvivien/gexe"
)

func run(cmd string, err func(error) error) (string, error) {
	p := gexe.New().RunProc(cmd)
	if p.Err() != nil {
		return "", err(p.Err())
	} else {
		logrus.Info(p.Result())
	}
	return p.Result(), nil
}
