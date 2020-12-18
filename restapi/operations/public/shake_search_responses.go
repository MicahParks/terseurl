// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"pulley.com/shakesearch/models"
)

// ShakeSearchOKCode is the HTTP code returned for type ShakeSearchOK
const ShakeSearchOKCode int = 200

/*ShakeSearchOK The fuzzy search was a success.

swagger:response shakeSearchOK
*/
type ShakeSearchOK struct {
	/*An array of strings that contain a match. Sorted by their score.

	 */
	Location []string `json:"Location"`
}

// NewShakeSearchOK creates ShakeSearchOK with default headers values
func NewShakeSearchOK() *ShakeSearchOK {

	return &ShakeSearchOK{}
}

// WithLocation adds the location to the shake search o k response
func (o *ShakeSearchOK) WithLocation(location []string) *ShakeSearchOK {
	o.Location = location
	return o
}

// SetLocation sets the location to the shake search o k response
func (o *ShakeSearchOK) SetLocation(location []string) {
	o.Location = location
}

// WriteResponse to the client
func (o *ShakeSearchOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Location

	var locationIR []string
	for _, locationI := range o.Location {
		locationIS := locationI
		if locationIS != "" {
			locationIR = append(locationIR, locationIS)
		}
	}
	location := swag.JoinByFormat(locationIR, "")
	if len(location) > 0 {
		hv := location[0]
		if hv != "" {
			rw.Header().Set("Location", hv)
		}
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*ShakeSearchDefault Unexpected error.

swagger:response shakeSearchDefault
*/
type ShakeSearchDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewShakeSearchDefault creates ShakeSearchDefault with default headers values
func NewShakeSearchDefault(code int) *ShakeSearchDefault {
	if code <= 0 {
		code = 500
	}

	return &ShakeSearchDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the shake search default response
func (o *ShakeSearchDefault) WithStatusCode(code int) *ShakeSearchDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the shake search default response
func (o *ShakeSearchDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the shake search default response
func (o *ShakeSearchDefault) WithPayload(payload *models.Error) *ShakeSearchDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the shake search default response
func (o *ShakeSearchDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ShakeSearchDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
