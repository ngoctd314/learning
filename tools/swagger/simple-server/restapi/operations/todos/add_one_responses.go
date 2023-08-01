// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"learn/models"
)

// AddOneCreatedCode is the HTTP code returned for type AddOneCreated
const AddOneCreatedCode int = 201

/*
AddOneCreated Created

swagger:response addOneCreated
*/
type AddOneCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Item `json:"body,omitempty"`
}

// NewAddOneCreated creates AddOneCreated with default headers values
func NewAddOneCreated() *AddOneCreated {

	return &AddOneCreated{}
}

// WithPayload adds the payload to the add one created response
func (o *AddOneCreated) WithPayload(payload *models.Item) *AddOneCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add one created response
func (o *AddOneCreated) SetPayload(payload *models.Item) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddOneCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
AddOneDefault error

swagger:response addOneDefault
*/
type AddOneDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddOneDefault creates AddOneDefault with default headers values
func NewAddOneDefault(code int) *AddOneDefault {
	if code <= 0 {
		code = 500
	}

	return &AddOneDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the add one default response
func (o *AddOneDefault) WithStatusCode(code int) *AddOneDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the add one default response
func (o *AddOneDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the add one default response
func (o *AddOneDefault) WithPayload(payload *models.Error) *AddOneDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add one default response
func (o *AddOneDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddOneDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
