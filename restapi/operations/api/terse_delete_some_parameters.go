// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"
)

// NewTerseDeleteSomeParams creates a new TerseDeleteSomeParams object
//
// There are no default values defined in the spec.
func NewTerseDeleteSomeParams() TerseDeleteSomeParams {

	return TerseDeleteSomeParams{}
}

// TerseDeleteSomeParams contains all the bound params for the terse delete some operation
// typically these are obtained from a http.Request
//
// swagger:parameters terseDeleteSome
type TerseDeleteSomeParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Indicate if Terse and or Visits data should be deleted and for which shortened URLs.
	  Required: true
	  In: body
	*/
	Info TerseDeleteSomeBody
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewTerseDeleteSomeParams() beforehand.
func (o *TerseDeleteSomeParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body TerseDeleteSomeBody
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("info", "body", ""))
			} else {
				res = append(res, errors.NewParseError("info", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Info = body
			}
		}
	} else {
		res = append(res, errors.Required("info", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
