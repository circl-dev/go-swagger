// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/circl-dev/runtime/middleware"
)

// PutTestHandlerFunc turns a function with the right signature into a put test handler
type PutTestHandlerFunc func(PutTestParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutTestHandlerFunc) Handle(params PutTestParams) middleware.Responder {
	return fn(params)
}

// PutTestHandler interface for that can handle valid put test params
type PutTestHandler interface {
	Handle(PutTestParams) middleware.Responder
}

// NewPutTest creates a new http.Handler for the put test operation
func NewPutTest(ctx *middleware.Context, handler PutTestHandler) *PutTest {
	return &PutTest{Context: ctx, Handler: handler}
}

/* PutTest swagger:route PUT /test putTest

PutTest put test API

*/
type PutTest struct {
	Context *middleware.Context
	Handler PutTestHandler
}

func (o *PutTest) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPutTestParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
