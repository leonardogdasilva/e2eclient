package config

import (
	"encoding/json"
	"io/ioutil"
	yaml2 "sigs.k8s.io/yaml"
	"strings"
	"testing"

	"github.com/pismo/e2eclient/pkg/config/v1alpha1"
)

func TestValidateSchema(t *testing.T) {

	cfgPath := "./test_assets/config_test_simple_scenario.yaml"

	if err := ValidateSchemaFile(cfgPath, []byte(v1alpha1.JSONSchema)); err != nil {
		t.Errorf("Validation of config file %s against the default schema failed: %+v", cfgPath, err)
	}

	filecontent, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		t.Errorf("Error loading yaml %s", cfgPath)
	}

	filecontent, err = yaml2.YAMLToJSON(filecontent)

	configFile := v1alpha1.SchemaJson{}
	err = json.Unmarshal(filecontent, &configFile)
	if err != nil {
		t.Errorf("Error loading yaml for %s %s", cfgPath, err)
	}
	if strings.Compare(*configFile.Name, "foo") != 0 {
		t.Fatalf("Invalid config name %s", *configFile.Name)
	}

}

func TestValidateSchemaFail(t *testing.T) {

	cfgPath := "./test_assets/config_test_simple_invalid_scenario.yaml"

	var err error
	if err = ValidateSchemaFile(cfgPath, []byte(v1alpha1.JSONSchema)); err == nil {
		t.Errorf("Validation of config file %s against the default schema passed where we expected a failure", cfgPath)
	}

	expectedErrorText := `- name: Invalid type. Expected: string, given: integer
`

	if err.Error() != expectedErrorText {
		t.Errorf("Actual validation error\n%s\ndoes not match expected error\n%s\n", err.Error(), expectedErrorText)
	}

}
