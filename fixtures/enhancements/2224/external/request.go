// Code generated by go-swagger; DO NOT EDIT.

package external

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	httpext "net/http"

	"github.com/go-openapi/strfmt"
	"github.com/circl-dev/runtime"
)

// Request request
//
// swagger:model Request
type Request struct {
	httpext.Request
}

func (m Request) Validate(formats strfmt.Registry) error {
	var f interface{} = m.Request
	if v, ok := f.(runtime.Validatable); ok {
		return v.Validate(formats)
	}
	return nil
}

func (m Request) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var f interface{} = m.Request
	if v, ok := f.(runtime.ContextValidatable); ok {
		return v.ContextValidate(ctx, formats)
	}
	return nil
}
