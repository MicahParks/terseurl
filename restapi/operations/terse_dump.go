// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// TerseDumpHandlerFunc turns a function with the right signature into a terse dump handler
type TerseDumpHandlerFunc func(TerseDumpParams) middleware.Responder

// Handle executing the request and returning a response
func (fn TerseDumpHandlerFunc) Handle(params TerseDumpParams) middleware.Responder {
	return fn(params)
}

// TerseDumpHandler interface for that can handle valid terse dump params
type TerseDumpHandler interface {
	Handle(TerseDumpParams) middleware.Responder
}

// NewTerseDump creates a new http.Handler for the terse dump operation
func NewTerseDump(ctx *middleware.Context, handler TerseDumpHandler) *TerseDump {
	return &TerseDump{Context: ctx, Handler: handler}
}

/*TerseDump swagger:route GET /api/dump terseDump

TerseDump terse dump API

*/
type TerseDump struct {
	Context *middleware.Context
	Handler TerseDumpHandler
}

func (o *TerseDump) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewTerseDumpParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
