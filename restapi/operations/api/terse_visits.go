// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// TerseVisitsHandlerFunc turns a function with the right signature into a terse visits handler
type TerseVisitsHandlerFunc func(TerseVisitsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn TerseVisitsHandlerFunc) Handle(params TerseVisitsParams) middleware.Responder {
	return fn(params)
}

// TerseVisitsHandler interface for that can handle valid terse visits params
type TerseVisitsHandler interface {
	Handle(TerseVisitsParams) middleware.Responder
}

// NewTerseVisits creates a new http.Handler for the terse visits operation
func NewTerseVisits(ctx *middleware.Context, handler TerseVisitsHandler) *TerseVisits {
	return &TerseVisits{Context: ctx, Handler: handler}
}

/*TerseVisits swagger:route GET /api/visits/{shortened} api terseVisits

Get the Visit data for a single shortened URL.

*/
type TerseVisits struct {
	Context *middleware.Context
	Handler TerseVisitsHandler
}

func (o *TerseVisits) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewTerseVisitsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}