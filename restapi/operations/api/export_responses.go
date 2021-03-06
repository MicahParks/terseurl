// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/MicahParks/terseurl/models"
)

// ExportOKCode is the HTTP code returned for type ExportOK
const ExportOKCode int = 200

/*ExportOK The export was successfully retrieved.

swagger:response exportOK
*/
type ExportOK struct {

	/*
	  In: Body
	*/
	Payload map[string]*models.Export `json:"body,omitempty"`
}

// NewExportOK creates ExportOK with default headers values
func NewExportOK() *ExportOK {

	return &ExportOK{}
}

// WithPayload adds the payload to the export o k response
func (o *ExportOK) WithPayload(payload map[string]*models.Export) *ExportOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the export o k response
func (o *ExportOK) SetPayload(payload map[string]*models.Export) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ExportOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty map
		payload = make(map[string]*models.Export, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*ExportDefault Unexpected error.

swagger:response exportDefault
*/
type ExportDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewExportDefault creates ExportDefault with default headers values
func NewExportDefault(code int) *ExportDefault {
	if code <= 0 {
		code = 500
	}

	return &ExportDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the export default response
func (o *ExportDefault) WithStatusCode(code int) *ExportDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the export default response
func (o *ExportDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the export default response
func (o *ExportDefault) WithPayload(payload *models.Error) *ExportDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the export default response
func (o *ExportDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ExportDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
