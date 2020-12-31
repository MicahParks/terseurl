// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/MicahParks/terseurl/models"
)

// NewTerseDeleteParams creates a new TerseDeleteParams object
// no default values defined in spec.
func NewTerseDeleteParams() TerseDeleteParams {

	return TerseDeleteParams{}
}

// TerseDeleteParams contains all the bound params for the terse delete operation
// typically these are obtained from a http.Request
//
// swagger:parameters terseDelete
type TerseDeleteParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*A JSON object containing the deletion information. If Terse or Visits data is marked for deletion, it will all be deleted.
	  Required: true
	  In: body
	*/
	Delete *models.Delete
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewTerseDeleteParams() beforehand.
func (o *TerseDeleteParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Delete
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("delete", "body", ""))
			} else {
				res = append(res, errors.NewParseError("delete", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Delete = &body
			}
		}
	} else {
		res = append(res, errors.Required("delete", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
