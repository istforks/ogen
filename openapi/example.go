package openapi

import (
	"github.com/istforks/ogen/jsonschema"
	"github.com/istforks/ogen/location"
)

// Example is an OpenAPI Example.
type Example struct {
	Ref Ref

	Summary       string
	Description   string
	Value         jsonschema.Example
	ExternalValue string

	location.Pointer `json:"-" yaml:"-"`
}
