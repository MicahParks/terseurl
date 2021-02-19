// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	"github.com/MicahParks/terseurl/models"
)

// NewImportParams creates a new ImportParams object
//
// There are no default values defined in the spec.
func NewImportParams() ImportParams {

	return ImportParams{}
}

// ImportParams contains all the bound params for the import operation
// typically these are obtained from a http.Request
//
// swagger:parameters import
type ImportParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*A JSON object containing the deletion information. If Terse or Visits data is marked for deletion, it will all be deleted. An object matching shortened URLs to their previously exported data is also required.
	  Required: true
	  In: body
	*/
	Import map[string]models.Export
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewImportParams() beforehand.
func (o *ImportParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body map[string]models.Export
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("import", "body", ""))
			} else {
				res = append(res, errors.NewParseError("import", "body", "", err))
			}
		} else {
			// validate map of body objects
			for k := range body {
				if err := validate.Required(fmt.Sprintf("%s.%v", "import", k), "body", body[k]); err != nil {
					return err
				}
				if val, ok := body[k]; ok {
					if err := val.Validate(route.Formats); err != nil {
						res = append(res, err)
						break
					}
				}
			}

			if len(res) == 0 {
				o.Import = body
			}
		}
	} else {
		res = append(res, errors.Required("import", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}