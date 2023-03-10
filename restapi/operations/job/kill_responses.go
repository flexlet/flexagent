// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"flexagent/models"
)

// KillOKCode is the HTTP code returned for type KillOK
const KillOKCode int = 200

/*KillOK Kill job succeeded

swagger:response killOK
*/
type KillOK struct {

	/*
	  In: Body
	*/
	Payload *models.Job `json:"body,omitempty"`
}

// NewKillOK creates KillOK with default headers values
func NewKillOK() *KillOK {

	return &KillOK{}
}

// WithPayload adds the payload to the kill o k response
func (o *KillOK) WithPayload(payload *models.Job) *KillOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the kill o k response
func (o *KillOK) SetPayload(payload *models.Job) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *KillOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// KillBadRequestCode is the HTTP code returned for type KillBadRequest
const KillBadRequestCode int = 400

/*KillBadRequest kill bad request

swagger:response killBadRequest
*/
type KillBadRequest struct {

	/*
	  In: Body
	*/
	Payload *KillBadRequestBody `json:"body,omitempty"`
}

// NewKillBadRequest creates KillBadRequest with default headers values
func NewKillBadRequest() *KillBadRequest {

	return &KillBadRequest{}
}

// WithPayload adds the payload to the kill bad request response
func (o *KillBadRequest) WithPayload(payload *KillBadRequestBody) *KillBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the kill bad request response
func (o *KillBadRequest) SetPayload(payload *KillBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *KillBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// KillUnauthorizedCode is the HTTP code returned for type KillUnauthorized
const KillUnauthorizedCode int = 401

/*KillUnauthorized kill unauthorized

swagger:response killUnauthorized
*/
type KillUnauthorized struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewKillUnauthorized creates KillUnauthorized with default headers values
func NewKillUnauthorized() *KillUnauthorized {

	return &KillUnauthorized{}
}

// WithPayload adds the payload to the kill unauthorized response
func (o *KillUnauthorized) WithPayload(payload interface{}) *KillUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the kill unauthorized response
func (o *KillUnauthorized) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *KillUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// KillForbiddenCode is the HTTP code returned for type KillForbidden
const KillForbiddenCode int = 403

/*KillForbidden kill forbidden

swagger:response killForbidden
*/
type KillForbidden struct {

	/*
	  In: Body
	*/
	Payload *KillForbiddenBody `json:"body,omitempty"`
}

// NewKillForbidden creates KillForbidden with default headers values
func NewKillForbidden() *KillForbidden {

	return &KillForbidden{}
}

// WithPayload adds the payload to the kill forbidden response
func (o *KillForbidden) WithPayload(payload *KillForbiddenBody) *KillForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the kill forbidden response
func (o *KillForbidden) SetPayload(payload *KillForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *KillForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// KillNotFoundCode is the HTTP code returned for type KillNotFound
const KillNotFoundCode int = 404

/*KillNotFound kill not found

swagger:response killNotFound
*/
type KillNotFound struct {

	/*
	  In: Body
	*/
	Payload *KillNotFoundBody `json:"body,omitempty"`
}

// NewKillNotFound creates KillNotFound with default headers values
func NewKillNotFound() *KillNotFound {

	return &KillNotFound{}
}

// WithPayload adds the payload to the kill not found response
func (o *KillNotFound) WithPayload(payload *KillNotFoundBody) *KillNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the kill not found response
func (o *KillNotFound) SetPayload(payload *KillNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *KillNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// KillInternalServerErrorCode is the HTTP code returned for type KillInternalServerError
const KillInternalServerErrorCode int = 500

/*KillInternalServerError kill internal server error

swagger:response killInternalServerError
*/
type KillInternalServerError struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewKillInternalServerError creates KillInternalServerError with default headers values
func NewKillInternalServerError() *KillInternalServerError {

	return &KillInternalServerError{}
}

// WithPayload adds the payload to the kill internal server error response
func (o *KillInternalServerError) WithPayload(payload interface{}) *KillInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the kill internal server error response
func (o *KillInternalServerError) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *KillInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
