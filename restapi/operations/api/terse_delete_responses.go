// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/MicahParks/terseurl/models"
)

// TerseDeleteOKCode is the HTTP code returned for type TerseDeleteOK
const TerseDeleteOKCode int = 200

/*TerseDeleteOK terse delete o k

swagger:response terseDeleteOK
*/
type TerseDeleteOK struct {
}

// NewTerseDeleteOK creates TerseDeleteOK with default headers values
func NewTerseDeleteOK() *TerseDeleteOK {

	return &TerseDeleteOK{}
}

// WriteResponse to the client
func (o *TerseDeleteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*TerseDeleteDefault Unexpected error.

swagger:response terseDeleteDefault
*/
type TerseDeleteDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewTerseDeleteDefault creates TerseDeleteDefault with default headers values
func NewTerseDeleteDefault(code int) *TerseDeleteDefault {
	if code <= 0 {
		code = 500
	}

	return &TerseDeleteDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the terse delete default response
func (o *TerseDeleteDefault) WithStatusCode(code int) *TerseDeleteDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the terse delete default response
func (o *TerseDeleteDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the terse delete default response
func (o *TerseDeleteDefault) WithPayload(payload *models.Error) *TerseDeleteDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the terse delete default response
func (o *TerseDeleteDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TerseDeleteDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
