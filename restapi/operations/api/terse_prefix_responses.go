// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/MicahParks/terseurl/models"
)

// TersePrefixOKCode is the HTTP code returned for type TersePrefixOK
const TersePrefixOKCode int = 200

/*TersePrefixOK The HTTP prefix all shortened URLs have.

swagger:response tersePrefixOK
*/
type TersePrefixOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewTersePrefixOK creates TersePrefixOK with default headers values
func NewTersePrefixOK() *TersePrefixOK {

	return &TersePrefixOK{}
}

// WithPayload adds the payload to the terse prefix o k response
func (o *TersePrefixOK) WithPayload(payload string) *TersePrefixOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the terse prefix o k response
func (o *TersePrefixOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TersePrefixOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*TersePrefixDefault Unexpected error.

swagger:response tersePrefixDefault
*/
type TersePrefixDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewTersePrefixDefault creates TersePrefixDefault with default headers values
func NewTersePrefixDefault(code int) *TersePrefixDefault {
	if code <= 0 {
		code = 500
	}

	return &TersePrefixDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the terse prefix default response
func (o *TersePrefixDefault) WithStatusCode(code int) *TersePrefixDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the terse prefix default response
func (o *TersePrefixDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the terse prefix default response
func (o *TersePrefixDefault) WithPayload(payload *models.Error) *TersePrefixDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the terse prefix default response
func (o *TersePrefixDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TersePrefixDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}