// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/istforks/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// ProbeLiveness implements probeLiveness operation.
//
// Liveness probe for kubernetes.
//
// GET /healthz
func (UnimplementedHandler) ProbeLiveness(ctx context.Context) (r string, _ error) {
	return r, ht.ErrNotImplemented
}

// NewError creates *ErrorStatusCode from error returned by handler.
//
// Used for common default response.
func (UnimplementedHandler) NewError(ctx context.Context, err error) (r *ErrorStatusCode) {
	r = new(ErrorStatusCode)
	return r
}
