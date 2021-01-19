// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// TersePrefixHandlerFunc turns a function with the right signature into a terse prefix handler
type TersePrefixHandlerFunc func(TersePrefixParams) middleware.Responder

// Handle executing the request and returning a response
func (fn TersePrefixHandlerFunc) Handle(params TersePrefixParams) middleware.Responder {
	return fn(params)
}

// TersePrefixHandler interface for that can handle valid terse prefix params
type TersePrefixHandler interface {
	Handle(TersePrefixParams) middleware.Responder
}

// NewTersePrefix creates a new http.Handler for the terse prefix operation
func NewTersePrefix(ctx *middleware.Context, handler TersePrefixHandler) *TersePrefix {
	return &TersePrefix{Context: ctx, Handler: handler}
}

/* TersePrefix swagger:route GET /api/prefix api tersePrefix

Client's web browser is requesting what HTTP prefix all shortened URLs have.

Provides the HTTP prefix all shortened URLs have.

*/
type TersePrefix struct {
	Context *middleware.Context
	Handler TersePrefixHandler
}

func (o *TersePrefix) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewTersePrefixParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
