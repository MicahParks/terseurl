// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// URLGetHandlerFunc turns a function with the right signature into a url get handler
type URLGetHandlerFunc func(URLGetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn URLGetHandlerFunc) Handle(params URLGetParams) middleware.Responder {
	return fn(params)
}

// URLGetHandler interface for that can handle valid url get params
type URLGetHandler interface {
	Handle(URLGetParams) middleware.Responder
}

// NewURLGet creates a new http.Handler for the url get operation
func NewURLGet(ctx *middleware.Context, handler URLGetHandler) *URLGet {
	return &URLGet{Context: ctx, Handler: handler}
}

/*URLGet swagger:route GET /{shortened} urlGet

Use the shortened URL. It will redirect to the full URL if it has not expired.

*/
type URLGet struct {
	Context *middleware.Context
	Handler URLGetHandler
}

func (o *URLGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewURLGetParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}