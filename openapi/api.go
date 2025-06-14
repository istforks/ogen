// Package openapi represents OpenAPI v3 Specification in Go.
package openapi

import (
	"github.com/istforks/ogen/jsonpointer"
	"github.com/istforks/ogen/jsonschema"
)

// Ref is a JSON Reference.
type Ref = jsonpointer.RefKey

// API represents parsed OpenAPI spec.
type API struct {
	Version    Version
	Tags       []Tag
	Servers    []Server
	Operations []*Operation
	Webhooks   []Webhook
	Components *Components
	Info       Info
}

// Components represent parsed components of OpenAPI spec.
type Components struct {
	Schemas       map[string]*jsonschema.Schema
	Responses     map[string]*Response
	Parameters    map[string]*Parameter
	Examples      map[string]*Example
	RequestBodies map[string]*RequestBody
}
