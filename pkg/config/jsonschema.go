package config

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/xeipuuv/gojsonschema"
	"sigs.k8s.io/yaml"
)

// ValidateSchemaFile takes a filepath, reads the file and validates it against a JSON schema
func ValidateSchemaFile(filepath string, schema []byte) error {
	logrus.Debugf("Validating file %s against default JSONSchema...", filepath)

	fileContents, err := ioutil.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %+v", filepath, err)
	}

	var content map[string]interface{}
	if err := yaml.Unmarshal(fileContents, &content); err != nil {
		return fmt.Errorf("failed to unmarshal the content of %s to a map: %+v", filepath, err)
	}

	return ValidateSchema(content, schema)
}

// ValidateSchema validates a YAML construct (non-struct representation) against a JSON Schema
func ValidateSchema(content map[string]interface{}, schemaJSON []byte) error {

	contentYaml, err := yaml.Marshal(content)
	if err != nil {
		return err
	}
	contentJSON, err := yaml.YAMLToJSON(contentYaml)
	if err != nil {
		return err
	}

	if bytes.Equal(contentJSON, []byte("null")) {
		contentJSON = []byte("{}") // non-json yaml struct
	}

	configLoader := gojsonschema.NewBytesLoader(contentJSON)
	schemaLoader := gojsonschema.NewBytesLoader(schemaJSON)

	result, err := gojsonschema.Validate(schemaLoader, configLoader)
	if err != nil {
		return err
	}

	logrus.Debugf("JSON Schema Validation Result: %+v", result)

	if !result.Valid() {
		var sb strings.Builder
		for _, desc := range result.Errors() {
			sb.WriteString(fmt.Sprintf("- %s\n", desc))
		}
		return errors.New(sb.String())
	}

	return nil
}
