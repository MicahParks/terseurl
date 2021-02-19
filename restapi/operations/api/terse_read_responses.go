// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/MicahParks/terseurl/models"
)

// TerseReadOKCode is the HTTP code returned for type TerseReadOK
const TerseReadOKCode int = 200

/*TerseReadOK The Terse data was successfully retrieved.

swagger:response terseReadOK
*/
type TerseReadOK struct {

	/*
	  In: Body
	*/
	Payload *models.Terse `json:"body,omitempty"`
}

// NewTerseReadOK creates TerseReadOK with default headers values
func NewTerseReadOK() *TerseReadOK {

	return &TerseReadOK{}
}

// WithPayload adds the payload to the terse read o k response
func (o *TerseReadOK) WithPayload(payload *models.Terse) *TerseReadOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the terse read o k response
func (o *TerseReadOK) SetPayload(payload *models.Terse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TerseReadOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*TerseReadDefault Unexpected error.

swagger:response terseReadDefault
*/
type TerseReadDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewTerseReadDefault creates TerseReadDefault with default headers values
func NewTerseReadDefault(code int) *TerseReadDefault {
	if code <= 0 {
		code = 500
	}

	return &TerseReadDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the terse read default response
func (o *TerseReadDefault) WithStatusCode(code int) *TerseReadDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the terse read default response
func (o *TerseReadDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the terse read default response
func (o *TerseReadDefault) WithPayload(payload *models.Error) *TerseReadDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the terse read default response
func (o *TerseReadDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TerseReadDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}