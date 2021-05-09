package config

import (
	"reflect"
	"strings"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	yaml := `
  repository:
    git: git@github.com:pismo/api-gateway-e2e-test
    branch: canary-api
  components:
  - name: auth-api
    workdir: src/auth-api
    componentdir: auth-api
    git: git@github.com:pismo/auth-api
`
	r := strings.NewReader(yaml)
	err := parseConfig(r)
	if err != nil {
		t.Errorf("%v", err)
	}
	c := Config{
		Repository: Repository{
			Git:    "git@github.com:pismo/api-gateway-e2e-test",
			Branch: "canary-api",
		},
		Components: []Component{
			{
				Name:         "auth-api",
				WorkDir:      "src/auth-api",
				ComponentDir: "auth-api",
				Git:          "git@github.com:pismo/pismo-zuul-gateway",
			},
		},
	}
	if reflect.DeepEqual(c, ClientConfig) {
		t.Errorf("invalid configuration processing '%v' != '%v'", ClientConfig, c)
	}
}
