// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/MicahParks/terseurl/models"
)

// TerseExportSomeOKCode is the HTTP code returned for type TerseExportSomeOK
const TerseExportSomeOKCode int = 200

/*TerseExportSomeOK The export was successfully retrieved.

swagger:response terseExportSomeOK
*/
type TerseExportSomeOK struct {

	/*
	  In: Body
	*/
	Payload map[string]models.Export `json:"body,omitempty"`
}

// NewTerseExportSomeOK creates TerseExportSomeOK with default headers values
func NewTerseExportSomeOK() *TerseExportSomeOK {

	return &TerseExportSomeOK{}
}

// WithPayload adds the payload to the terse export some o k response
func (o *TerseExportSomeOK) WithPayload(payload map[string]models.Export) *TerseExportSomeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the terse export some o k response
func (o *TerseExportSomeOK) SetPayload(payload map[string]models.Export) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TerseExportSomeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty map
		payload = make(map[string]models.Export, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*TerseExportSomeDefault Unexpected error.

swagger:response terseExportSomeDefault
*/
type TerseExportSomeDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewTerseExportSomeDefault creates TerseExportSomeDefault with default headers values
func NewTerseExportSomeDefault(code int) *TerseExportSomeDefault {
	if code <= 0 {
		code = 500
	}

	return &TerseExportSomeDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the terse export some default response
func (o *TerseExportSomeDefault) WithStatusCode(code int) *TerseExportSomeDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the terse export some default response
func (o *TerseExportSomeDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the terse export some default response
func (o *TerseExportSomeDefault) WithPayload(payload *models.Error) *TerseExportSomeDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the terse export some default response
func (o *TerseExportSomeDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TerseExportSomeDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}