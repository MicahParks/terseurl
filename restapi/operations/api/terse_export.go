// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// TerseExportHandlerFunc turns a function with the right signature into a terse export handler
type TerseExportHandlerFunc func(TerseExportParams) middleware.Responder

// Handle executing the request and returning a response
func (fn TerseExportHandlerFunc) Handle(params TerseExportParams) middleware.Responder {
	return fn(params)
}

// TerseExportHandler interface for that can handle valid terse export params
type TerseExportHandler interface {
	Handle(TerseExportParams) middleware.Responder
}

// NewTerseExport creates a new http.Handler for the terse export operation
func NewTerseExport(ctx *middleware.Context, handler TerseExportHandler) *TerseExport {
	return &TerseExport{Context: ctx, Handler: handler}
}

/*TerseExport swagger:route GET /api/export api terseExport

Export all Terse and Visits data from the backend.

Depending on the underlying storage and amount of data, this may take a while.

*/
type TerseExport struct {
	Context *middleware.Context
	Handler TerseExportHandler
}

func (o *TerseExport) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewTerseExportParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}