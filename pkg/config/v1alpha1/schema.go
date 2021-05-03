package v1alpha1

import _ "embed"

// JSONSchema describes the schema used to validate config files
//go:embed schema.json
var JSONSchema string
