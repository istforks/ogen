// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/istforks/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// AddPet implements addPet operation.
//
// Creates a new pet in the store. Duplicates are allowed.
//
// POST /pets
func (UnimplementedHandler) AddPet(ctx context.Context, req *NewPet) (r *Pet, _ error) {
	return r, ht.ErrNotImplemented
}

// DeletePet implements deletePet operation.
//
// Deletes a single pet based on the ID supplied.
//
// DELETE /pets/{id}
func (UnimplementedHandler) DeletePet(ctx context.Context, params DeletePetParams) error {
	return ht.ErrNotImplemented
}

// FindPetByID implements find pet by id operation.
//
// Returns a user based on a single ID, if the user does not have access to the pet.
//
// GET /pets/{id}
func (UnimplementedHandler) FindPetByID(ctx context.Context, params FindPetByIDParams) (r *Pet, _ error) {
	return r, ht.ErrNotImplemented
}

// FindPets implements findPets operation.
//
// Returns all pets from the system that the user has access to
// Nam sed condimentum est. Maecenas tempor sagittis sapien, nec rhoncus sem sagittis sit amet.
// Aenean at gravida augue, ac iaculis sem. Curabitur odio lorem, ornare eget elementum nec, cursus
// id lectus. Duis mi turpis, pulvinar ac eros ac, tincidunt varius justo. In hac habitasse platea
// dictumst. Integer at adipiscing ante, a sagittis ligula. Aenean pharetra tempor ante molestie
// imperdiet. Vivamus id aliquam diam. Cras quis velit non tortor eleifend sagittis. Praesent at enim
// pharetra urna volutpat venenatis eget eget mauris. In eleifend fermentum facilisis. Praesent enim
// enim, gravida ac sodales sed, placerat id erat. Suspendisse lacus dolor, consectetur non augue vel,
//
//	vehicula interdum libero. Morbi euismod sagittis libero sed lacinia.
//
// Sed tempus felis lobortis leo pulvinar rutrum. Nam mattis velit nisl, eu condimentum ligula luctus
// nec. Phasellus semper velit eget aliquet faucibus. In a mattis elit. Phasellus vel urna viverra,
// condimentum lorem id, rhoncus nibh. Ut pellentesque posuere elementum. Sed a varius odio. Morbi
// rhoncus ligula libero, vel eleifend nunc tristique vitae. Fusce et sem dui. Aenean nec scelerisque
// tortor. Fusce malesuada accumsan magna vel tempus. Quisque mollis felis eu dolor tristique, sit
// amet auctor felis gravida. Sed libero lorem, molestie sed nisl in, accumsan tempor nisi. Fusce
// sollicitudin massa ut lacinia mattis. Sed vel eleifend lorem. Pellentesque vitae felis pretium,
// pulvinar elit eu, euismod sapien.
//
// GET /pets
func (UnimplementedHandler) FindPets(ctx context.Context, params FindPetsParams) (r []Pet, _ error) {
	return r, ht.ErrNotImplemented
}

// NewError creates *ErrorStatusCode from error returned by handler.
//
// Used for common default response.
func (UnimplementedHandler) NewError(ctx context.Context, err error) (r *ErrorStatusCode) {
	r = new(ErrorStatusCode)
	return r
}
