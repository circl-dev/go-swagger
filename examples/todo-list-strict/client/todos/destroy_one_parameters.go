// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/circl-dev/runtime"
	cr "github.com/circl-dev/runtime/client"
)

// NewDestroyOneParams creates a new DestroyOneParams object
// with the default values initialized.
func NewDestroyOneParams() *DestroyOneParams {
	var ()
	return &DestroyOneParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDestroyOneParamsWithTimeout creates a new DestroyOneParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDestroyOneParamsWithTimeout(timeout time.Duration) *DestroyOneParams {
	var ()
	return &DestroyOneParams{

		timeout: timeout,
	}
}

// NewDestroyOneParamsWithContext creates a new DestroyOneParams object
// with the default values initialized, and the ability to set a context for a request
func NewDestroyOneParamsWithContext(ctx context.Context) *DestroyOneParams {
	var ()
	return &DestroyOneParams{

		Context: ctx,
	}
}

// NewDestroyOneParamsWithHTTPClient creates a new DestroyOneParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDestroyOneParamsWithHTTPClient(client *http.Client) *DestroyOneParams {
	var ()
	return &DestroyOneParams{
		HTTPClient: client,
	}
}

/*DestroyOneParams contains all the parameters to send to the API endpoint
for the destroy one operation typically these are written to a http.Request
*/
type DestroyOneParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the destroy one params
func (o *DestroyOneParams) WithTimeout(timeout time.Duration) *DestroyOneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the destroy one params
func (o *DestroyOneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the destroy one params
func (o *DestroyOneParams) WithContext(ctx context.Context) *DestroyOneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the destroy one params
func (o *DestroyOneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the destroy one params
func (o *DestroyOneParams) WithHTTPClient(client *http.Client) *DestroyOneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the destroy one params
func (o *DestroyOneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the destroy one params
func (o *DestroyOneParams) WithID(id string) *DestroyOneParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the destroy one params
func (o *DestroyOneParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DestroyOneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
