// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/MicahParks/terse-URL/models"
)

// TerseDumpOKCode is the HTTP code returned for type TerseDumpOK
const TerseDumpOKCode int = 200

/*TerseDumpOK terse dump o k

swagger:response terseDumpOK
*/
type TerseDumpOK struct {

	/*
	  In: Body
	*/
	Payload map[string]models.Dump `json:"body,omitempty"`
}

// NewTerseDumpOK creates TerseDumpOK with default headers values
func NewTerseDumpOK() *TerseDumpOK {

	return &TerseDumpOK{}
}

// WithPayload adds the payload to the terse dump o k response
func (o *TerseDumpOK) WithPayload(payload map[string]models.Dump) *TerseDumpOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the terse dump o k response
func (o *TerseDumpOK) SetPayload(payload map[string]models.Dump) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TerseDumpOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty map
		payload = make(map[string]models.Dump, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*TerseDumpDefault Unexpected error.

swagger:response terseDumpDefault
*/
type TerseDumpDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewTerseDumpDefault creates TerseDumpDefault with default headers values
func NewTerseDumpDefault(code int) *TerseDumpDefault {
	if code <= 0 {
		code = 500
	}

	return &TerseDumpDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the terse dump default response
func (o *TerseDumpDefault) WithStatusCode(code int) *TerseDumpDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the terse dump default response
func (o *TerseDumpDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the terse dump default response
func (o *TerseDumpDefault) WithPayload(payload *models.Error) *TerseDumpDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the terse dump default response
func (o *TerseDumpDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TerseDumpDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
