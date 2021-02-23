// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/MicahParks/terseurl/models"
)

// ShortenedPrefixOKCode is the HTTP code returned for type ShortenedPrefixOK
const ShortenedPrefixOKCode int = 200

/*ShortenedPrefixOK The HTTP prefix all shortened URLs have.

swagger:response shortenedPrefixOK
*/
type ShortenedPrefixOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewShortenedPrefixOK creates ShortenedPrefixOK with default headers values
func NewShortenedPrefixOK() *ShortenedPrefixOK {

	return &ShortenedPrefixOK{}
}

// WithPayload adds the payload to the shortened prefix o k response
func (o *ShortenedPrefixOK) WithPayload(payload string) *ShortenedPrefixOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the shortened prefix o k response
func (o *ShortenedPrefixOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ShortenedPrefixOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*ShortenedPrefixDefault Unexpected error.

swagger:response shortenedPrefixDefault
*/
type ShortenedPrefixDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewShortenedPrefixDefault creates ShortenedPrefixDefault with default headers values
func NewShortenedPrefixDefault(code int) *ShortenedPrefixDefault {
	if code <= 0 {
		code = 500
	}

	return &ShortenedPrefixDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the shortened prefix default response
func (o *ShortenedPrefixDefault) WithStatusCode(code int) *ShortenedPrefixDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the shortened prefix default response
func (o *ShortenedPrefixDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the shortened prefix default response
func (o *ShortenedPrefixDefault) WithPayload(payload *models.Error) *ShortenedPrefixDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the shortened prefix default response
func (o *ShortenedPrefixDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ShortenedPrefixDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
