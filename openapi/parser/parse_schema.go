package parser

import (
	"github.com/go-faster/errors"

	"github.com/istforks/ogen"
	"github.com/istforks/ogen/jsonpointer"
	"github.com/istforks/ogen/jsonschema"
)

func (p *parser) parseSchema(schema *ogen.Schema, ctx *jsonpointer.ResolveCtx) (*jsonschema.Schema, error) {
	s, err := p.schemaParser.Parse(schema.ToJSONSchema(), ctx)
	if err != nil {
		return nil, errors.Wrap(err, "parse schema")
	}
	return s, nil
}
