package main

import (
	"os"

	"github.com/pismo/e2eclient/cmd/client"
	"github.com/sirupsen/logrus"
)

func main() {

	if err := client.Run(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}
