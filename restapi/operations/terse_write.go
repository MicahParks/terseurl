// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/MicahParks/terse-URL/models"
)

// TerseWriteHandlerFunc turns a function with the right signature into a terse write handler
type TerseWriteHandlerFunc func(TerseWriteParams, *models.JWTInfo) middleware.Responder

// Handle executing the request and returning a response
func (fn TerseWriteHandlerFunc) Handle(params TerseWriteParams, principal *models.JWTInfo) middleware.Responder {
	return fn(params, principal)
}

// TerseWriteHandler interface for that can handle valid terse write params
type TerseWriteHandler interface {
	Handle(TerseWriteParams, *models.JWTInfo) middleware.Responder
}

// NewTerseWrite creates a new http.Handler for the terse write operation
func NewTerseWrite(ctx *middleware.Context, handler TerseWriteHandler) *TerseWrite {
	return &TerseWrite{Context: ctx, Handler: handler}
}

/*TerseWrite swagger:route POST /api/write/{operation} terseWrite

TerseWrite terse write API

*/
type TerseWrite struct {
	Context *middleware.Context
	Handler TerseWriteHandler
}

func (o *TerseWrite) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewTerseWriteParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.JWTInfo
	if uprinc != nil {
		principal = uprinc.(*models.JWTInfo) // this is really a models.JWTInfo, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}