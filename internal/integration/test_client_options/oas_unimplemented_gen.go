// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/istforks/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// Foo implements Foo operation.
//
// GET /foo
func (UnimplementedHandler) Foo(ctx context.Context, params FooParams) (r string, _ error) {
	return r, ht.ErrNotImplemented
}
