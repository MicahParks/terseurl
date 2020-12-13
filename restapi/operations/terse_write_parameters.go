// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"

	"github.com/MicahParks/terse-URL/models"
)

// NewTerseWriteParams creates a new TerseWriteParams object
// no default values defined in spec.
func NewTerseWriteParams() TerseWriteParams {

	return TerseWriteParams{}
}

// TerseWriteParams contains all the bound params for the terse write operation
// typically these are obtained from a http.Request
//
// swagger:parameters terseWrite
type TerseWriteParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	Operation string
	/*
	  Required: true
	  In: body
	*/
	Terse *models.TerseOptionalShortened
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewTerseWriteParams() beforehand.
func (o *TerseWriteParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rOperation, rhkOperation, _ := route.Params.GetOK("operation")
	if err := o.bindOperation(rOperation, rhkOperation, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.TerseOptionalShortened
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("terse", "body", ""))
			} else {
				res = append(res, errors.NewParseError("terse", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Terse = &body
			}
		}
	} else {
		res = append(res, errors.Required("terse", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindOperation binds and validates parameter Operation from path.
func (o *TerseWriteParams) bindOperation(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Operation = raw

	if err := o.validateOperation(formats); err != nil {
		return err
	}

	return nil
}

// validateOperation carries on validations for parameter Operation
func (o *TerseWriteParams) validateOperation(formats strfmt.Registry) error {

	if err := validate.EnumCase("operation", "path", o.Operation, []interface{}{"insert", "update", "upsert"}, true); err != nil {
		return err
	}

	return nil
}
