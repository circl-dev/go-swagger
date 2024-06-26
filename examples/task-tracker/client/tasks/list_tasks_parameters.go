// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/circl-dev/runtime"
	cr "github.com/circl-dev/runtime/client"
)

// NewListTasksParams creates a new ListTasksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListTasksParams() *ListTasksParams {
	return &ListTasksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListTasksParamsWithTimeout creates a new ListTasksParams object
// with the ability to set a timeout on a request.
func NewListTasksParamsWithTimeout(timeout time.Duration) *ListTasksParams {
	return &ListTasksParams{
		timeout: timeout,
	}
}

// NewListTasksParamsWithContext creates a new ListTasksParams object
// with the ability to set a context for a request.
func NewListTasksParamsWithContext(ctx context.Context) *ListTasksParams {
	return &ListTasksParams{
		Context: ctx,
	}
}

// NewListTasksParamsWithHTTPClient creates a new ListTasksParams object
// with the ability to set a custom HTTPClient for a request.
func NewListTasksParamsWithHTTPClient(client *http.Client) *ListTasksParams {
	return &ListTasksParams{
		HTTPClient: client,
	}
}

/* ListTasksParams contains all the parameters to send to the API endpoint
   for the list tasks operation.

   Typically these are written to a http.Request.
*/
type ListTasksParams struct {

	/* PageSize.

	   Amount of items to return in a single page

	   Format: int32
	   Default: 20
	*/
	PageSize *int32

	/* SinceID.

	   The last id that was seen.

	   Format: int64
	*/
	SinceID *int64

	/* Status.

	   the status to filter by
	*/
	Status []string

	/* Tags.

	   the tags to filter by
	*/
	Tags []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list tasks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListTasksParams) WithDefaults() *ListTasksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list tasks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListTasksParams) SetDefaults() {
	var (
		pageSizeDefault = int32(20)
	)

	val := ListTasksParams{
		PageSize: &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list tasks params
func (o *ListTasksParams) WithTimeout(timeout time.Duration) *ListTasksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list tasks params
func (o *ListTasksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list tasks params
func (o *ListTasksParams) WithContext(ctx context.Context) *ListTasksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list tasks params
func (o *ListTasksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list tasks params
func (o *ListTasksParams) WithHTTPClient(client *http.Client) *ListTasksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list tasks params
func (o *ListTasksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageSize adds the pageSize to the list tasks params
func (o *ListTasksParams) WithPageSize(pageSize *int32) *ListTasksParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list tasks params
func (o *ListTasksParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSinceID adds the sinceID to the list tasks params
func (o *ListTasksParams) WithSinceID(sinceID *int64) *ListTasksParams {
	o.SetSinceID(sinceID)
	return o
}

// SetSinceID adds the sinceId to the list tasks params
func (o *ListTasksParams) SetSinceID(sinceID *int64) {
	o.SinceID = sinceID
}

// WithStatus adds the status to the list tasks params
func (o *ListTasksParams) WithStatus(status []string) *ListTasksParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the list tasks params
func (o *ListTasksParams) SetStatus(status []string) {
	o.Status = status
}

// WithTags adds the tags to the list tasks params
func (o *ListTasksParams) WithTags(tags []string) *ListTasksParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the list tasks params
func (o *ListTasksParams) SetTags(tags []string) {
	o.Tags = tags
}

// WriteToRequest writes these params to a swagger request
func (o *ListTasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize int32

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt32(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}
	}

	if o.SinceID != nil {

		// query param sinceId
		var qrSinceID int64

		if o.SinceID != nil {
			qrSinceID = *o.SinceID
		}
		qSinceID := swag.FormatInt64(qrSinceID)
		if qSinceID != "" {

			if err := r.SetQueryParam("sinceId", qSinceID); err != nil {
				return err
			}
		}
	}

	if o.Status != nil {

		// binding items for status
		joinedStatus := o.bindParamStatus(reg)

		// query array param status
		if err := r.SetQueryParam("status", joinedStatus...); err != nil {
			return err
		}
	}

	if o.Tags != nil {

		// binding items for tags
		joinedTags := o.bindParamTags(reg)

		// query array param tags
		if err := r.SetQueryParam("tags", joinedTags...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamListTasks binds the parameter status
func (o *ListTasksParams) bindParamStatus(formats strfmt.Registry) []string {
	statusIR := o.Status

	var statusIC []string
	for _, statusIIR := range statusIR { // explode []string

		statusIIV := statusIIR // string as string
		statusIC = append(statusIC, statusIIV)
	}

	// items.CollectionFormat: "pipes"
	statusIS := swag.JoinByFormat(statusIC, "pipes")

	return statusIS
}

// bindParamListTasks binds the parameter tags
func (o *ListTasksParams) bindParamTags(formats strfmt.Registry) []string {
	tagsIR := o.Tags

	var tagsIC []string
	for _, tagsIIR := range tagsIR { // explode []string

		tagsIIV := tagsIIR // string as string
		tagsIC = append(tagsIC, tagsIIV)
	}

	// items.CollectionFormat: ""
	tagsIS := swag.JoinByFormat(tagsIC, "")

	return tagsIS
}
