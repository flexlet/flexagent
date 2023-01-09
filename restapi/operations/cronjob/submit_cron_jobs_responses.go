// Code generated by go-swagger; DO NOT EDIT.

package cronjob

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"flexagent/models"
)

// SubmitCronJobsOKCode is the HTTP code returned for type SubmitCronJobsOK
const SubmitCronJobsOKCode int = 200

/*SubmitCronJobsOK Submit cronjob succeeded

swagger:response submitCronJobsOK
*/
type SubmitCronJobsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.CronJob `json:"body,omitempty"`
}

// NewSubmitCronJobsOK creates SubmitCronJobsOK with default headers values
func NewSubmitCronJobsOK() *SubmitCronJobsOK {

	return &SubmitCronJobsOK{}
}

// WithPayload adds the payload to the submit cron jobs o k response
func (o *SubmitCronJobsOK) WithPayload(payload []*models.CronJob) *SubmitCronJobsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the submit cron jobs o k response
func (o *SubmitCronJobsOK) SetPayload(payload []*models.CronJob) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SubmitCronJobsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.CronJob, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// SubmitCronJobsBadRequestCode is the HTTP code returned for type SubmitCronJobsBadRequest
const SubmitCronJobsBadRequestCode int = 400

/*SubmitCronJobsBadRequest submit cron jobs bad request

swagger:response submitCronJobsBadRequest
*/
type SubmitCronJobsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *SubmitCronJobsBadRequestBody `json:"body,omitempty"`
}

// NewSubmitCronJobsBadRequest creates SubmitCronJobsBadRequest with default headers values
func NewSubmitCronJobsBadRequest() *SubmitCronJobsBadRequest {

	return &SubmitCronJobsBadRequest{}
}

// WithPayload adds the payload to the submit cron jobs bad request response
func (o *SubmitCronJobsBadRequest) WithPayload(payload *SubmitCronJobsBadRequestBody) *SubmitCronJobsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the submit cron jobs bad request response
func (o *SubmitCronJobsBadRequest) SetPayload(payload *SubmitCronJobsBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SubmitCronJobsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SubmitCronJobsUnauthorizedCode is the HTTP code returned for type SubmitCronJobsUnauthorized
const SubmitCronJobsUnauthorizedCode int = 401

/*SubmitCronJobsUnauthorized submit cron jobs unauthorized

swagger:response submitCronJobsUnauthorized
*/
type SubmitCronJobsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewSubmitCronJobsUnauthorized creates SubmitCronJobsUnauthorized with default headers values
func NewSubmitCronJobsUnauthorized() *SubmitCronJobsUnauthorized {

	return &SubmitCronJobsUnauthorized{}
}

// WithPayload adds the payload to the submit cron jobs unauthorized response
func (o *SubmitCronJobsUnauthorized) WithPayload(payload interface{}) *SubmitCronJobsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the submit cron jobs unauthorized response
func (o *SubmitCronJobsUnauthorized) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SubmitCronJobsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// SubmitCronJobsForbiddenCode is the HTTP code returned for type SubmitCronJobsForbidden
const SubmitCronJobsForbiddenCode int = 403

/*SubmitCronJobsForbidden submit cron jobs forbidden

swagger:response submitCronJobsForbidden
*/
type SubmitCronJobsForbidden struct {

	/*
	  In: Body
	*/
	Payload *SubmitCronJobsForbiddenBody `json:"body,omitempty"`
}

// NewSubmitCronJobsForbidden creates SubmitCronJobsForbidden with default headers values
func NewSubmitCronJobsForbidden() *SubmitCronJobsForbidden {

	return &SubmitCronJobsForbidden{}
}

// WithPayload adds the payload to the submit cron jobs forbidden response
func (o *SubmitCronJobsForbidden) WithPayload(payload *SubmitCronJobsForbiddenBody) *SubmitCronJobsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the submit cron jobs forbidden response
func (o *SubmitCronJobsForbidden) SetPayload(payload *SubmitCronJobsForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SubmitCronJobsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SubmitCronJobsNotFoundCode is the HTTP code returned for type SubmitCronJobsNotFound
const SubmitCronJobsNotFoundCode int = 404

/*SubmitCronJobsNotFound submit cron jobs not found

swagger:response submitCronJobsNotFound
*/
type SubmitCronJobsNotFound struct {

	/*
	  In: Body
	*/
	Payload *SubmitCronJobsNotFoundBody `json:"body,omitempty"`
}

// NewSubmitCronJobsNotFound creates SubmitCronJobsNotFound with default headers values
func NewSubmitCronJobsNotFound() *SubmitCronJobsNotFound {

	return &SubmitCronJobsNotFound{}
}

// WithPayload adds the payload to the submit cron jobs not found response
func (o *SubmitCronJobsNotFound) WithPayload(payload *SubmitCronJobsNotFoundBody) *SubmitCronJobsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the submit cron jobs not found response
func (o *SubmitCronJobsNotFound) SetPayload(payload *SubmitCronJobsNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SubmitCronJobsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SubmitCronJobsInternalServerErrorCode is the HTTP code returned for type SubmitCronJobsInternalServerError
const SubmitCronJobsInternalServerErrorCode int = 500

/*SubmitCronJobsInternalServerError submit cron jobs internal server error

swagger:response submitCronJobsInternalServerError
*/
type SubmitCronJobsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewSubmitCronJobsInternalServerError creates SubmitCronJobsInternalServerError with default headers values
func NewSubmitCronJobsInternalServerError() *SubmitCronJobsInternalServerError {

	return &SubmitCronJobsInternalServerError{}
}

// WithPayload adds the payload to the submit cron jobs internal server error response
func (o *SubmitCronJobsInternalServerError) WithPayload(payload interface{}) *SubmitCronJobsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the submit cron jobs internal server error response
func (o *SubmitCronJobsInternalServerError) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SubmitCronJobsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}