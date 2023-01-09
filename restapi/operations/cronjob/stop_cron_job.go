// Code generated by go-swagger; DO NOT EDIT.

package cronjob

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StopCronJobHandlerFunc turns a function with the right signature into a stop cron job handler
type StopCronJobHandlerFunc func(StopCronJobParams) middleware.Responder

// Handle executing the request and returning a response
func (fn StopCronJobHandlerFunc) Handle(params StopCronJobParams) middleware.Responder {
	return fn(params)
}

// StopCronJobHandler interface for that can handle valid stop cron job params
type StopCronJobHandler interface {
	Handle(StopCronJobParams) middleware.Responder
}

// NewStopCronJob creates a new http.Handler for the stop cron job operation
func NewStopCronJob(ctx *middleware.Context, handler StopCronJobHandler) *StopCronJob {
	return &StopCronJob{Context: ctx, Handler: handler}
}

/* StopCronJob swagger:route POST /cronjobs/{id}/stop Cronjob stopCronJob

Stop cronjob

*/
type StopCronJob struct {
	Context *middleware.Context
	Handler StopCronJobHandler
}

func (o *StopCronJob) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewStopCronJobParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// StopCronJobBadRequestBody stop cron job bad request body
//
// swagger:model StopCronJobBadRequestBody
type StopCronJobBadRequestBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this stop cron job bad request body
func (o *StopCronJobBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StopCronJobBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("stopCronJobBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this stop cron job bad request body based on context it is used
func (o *StopCronJobBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StopCronJobBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StopCronJobBadRequestBody) UnmarshalBinary(b []byte) error {
	var res StopCronJobBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// StopCronJobForbiddenBody stop cron job forbidden body
//
// swagger:model StopCronJobForbiddenBody
type StopCronJobForbiddenBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this stop cron job forbidden body
func (o *StopCronJobForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StopCronJobForbiddenBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("stopCronJobForbidden"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this stop cron job forbidden body based on context it is used
func (o *StopCronJobForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StopCronJobForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StopCronJobForbiddenBody) UnmarshalBinary(b []byte) error {
	var res StopCronJobForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// StopCronJobNotFoundBody stop cron job not found body
//
// swagger:model StopCronJobNotFoundBody
type StopCronJobNotFoundBody struct {

	// error
	// Required: true
	Error *string `json:"error"`

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this stop cron job not found body
func (o *StopCronJobNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StopCronJobNotFoundBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("stopCronJobNotFound"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

func (o *StopCronJobNotFoundBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("stopCronJobNotFound"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this stop cron job not found body based on context it is used
func (o *StopCronJobNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StopCronJobNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StopCronJobNotFoundBody) UnmarshalBinary(b []byte) error {
	var res StopCronJobNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}