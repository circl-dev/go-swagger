// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/circl-dev/runtime/middleware"
)

// PetGetHandlerFunc turns a function with the right signature into a pet get handler
type PetGetHandlerFunc func(PetGetParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PetGetHandlerFunc) Handle(params PetGetParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PetGetHandler interface for that can handle valid pet get params
type PetGetHandler interface {
	Handle(PetGetParams, interface{}) middleware.Responder
}

// NewPetGet creates a new http.Handler for the pet get operation
func NewPetGet(ctx *middleware.Context, handler PetGetHandler) *PetGet {
	return &PetGet{Context: ctx, Handler: handler}
}

/* PetGet swagger:route GET /pet/{petId} pet petGet

Get pet by it's ID

*/
type PetGet struct {
	Context *middleware.Context
	Handler PetGetHandler
}

func (o *PetGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPetGetParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
