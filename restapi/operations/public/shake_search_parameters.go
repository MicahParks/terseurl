// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewShakeSearchParams creates a new ShakeSearchParams object
// with the default values initialized.
func NewShakeSearchParams() ShakeSearchParams {

	var (
		// initialize parameters with default values

		maxResultsDefault = int64(20)
	)

	return ShakeSearchParams{
		MaxResults: &maxResultsDefault,
	}
}

// ShakeSearchParams contains all the bound params for the shake search operation
// typically these are obtained from a http.Request
//
// swagger:parameters shakeSearch
type ShakeSearchParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The maximum number of results to return.
	  In: query
	  Default: 20
	*/
	MaxResults *int64
	/*The search query.
	  Required: true
	  In: query
	*/
	Q string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewShakeSearchParams() beforehand.
func (o *ShakeSearchParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qMaxResults, qhkMaxResults, _ := qs.GetOK("maxResults")
	if err := o.bindMaxResults(qMaxResults, qhkMaxResults, route.Formats); err != nil {
		res = append(res, err)
	}

	qQ, qhkQ, _ := qs.GetOK("q")
	if err := o.bindQ(qQ, qhkQ, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindMaxResults binds and validates parameter MaxResults from query.
func (o *ShakeSearchParams) bindMaxResults(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewShakeSearchParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("maxResults", "query", "int64", raw)
	}
	o.MaxResults = &value

	return nil
}

// bindQ binds and validates parameter Q from query.
func (o *ShakeSearchParams) bindQ(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("q", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("q", "query", raw); err != nil {
		return err
	}

	o.Q = raw

	return nil
}
