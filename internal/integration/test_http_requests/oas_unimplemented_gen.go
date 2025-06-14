// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/istforks/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// AllRequestBodies implements allRequestBodies operation.
//
// POST /allRequestBodies
func (UnimplementedHandler) AllRequestBodies(ctx context.Context, req AllRequestBodiesReq) (r AllRequestBodiesOK, _ error) {
	return r, ht.ErrNotImplemented
}

// AllRequestBodiesOptional implements allRequestBodiesOptional operation.
//
// POST /allRequestBodiesOptional
func (UnimplementedHandler) AllRequestBodiesOptional(ctx context.Context, req AllRequestBodiesOptionalReq) (r AllRequestBodiesOptionalOK, _ error) {
	return r, ht.ErrNotImplemented
}

// Base64Request implements base64Request operation.
//
// POST /base64Request
func (UnimplementedHandler) Base64Request(ctx context.Context, req Base64RequestReq) (r Base64RequestOK, _ error) {
	return r, ht.ErrNotImplemented
}

// MaskContentType implements maskContentType operation.
//
// POST /maskContentType
func (UnimplementedHandler) MaskContentType(ctx context.Context, req *MaskContentTypeReqWithContentType) (r *MaskResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// MaskContentTypeOptional implements maskContentTypeOptional operation.
//
// POST /maskContentTypeOptional
func (UnimplementedHandler) MaskContentTypeOptional(ctx context.Context, req *MaskContentTypeOptionalReqWithContentType) (r *MaskResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// StreamJSON implements streamJSON operation.
//
// POST /streamJSON
func (UnimplementedHandler) StreamJSON(ctx context.Context, req []float64) (r float64, _ error) {
	return r, ht.ErrNotImplemented
}
