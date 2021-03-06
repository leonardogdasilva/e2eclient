module github.com/leonardogdasilva

replace github.com/pismo/e2eclient => ./

go 1.16

require (
	github.com/onsi/ginkgo v1.16.1
	github.com/onsi/gomega v1.11.0
	github.com/pismo/e2eclient v0.0.0-00010101000000-000000000000
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.1.3
	github.com/vladimirvivien/gexe v0.1.0
	github.com/xeipuuv/gojsonschema v1.2.0
	go.starlark.net v0.0.0-20201006213952-227f4aabceb5
	gopkg.in/yaml.v2 v2.4.0
	sigs.k8s.io/yaml v1.2.0
)
