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

// StartCronJobHandlerFunc turns a function with the right signature into a start cron job handler
type StartCronJobHandlerFunc func(StartCronJobParams) middleware.Responder

// Handle executing the request and returning a response
func (fn StartCronJobHandlerFunc) Handle(params StartCronJobParams) middleware.Responder {
	return fn(params)
}

// StartCronJobHandler interface for that can handle valid start cron job params
type StartCronJobHandler interface {
	Handle(StartCronJobParams) middleware.Responder
}

// NewStartCronJob creates a new http.Handler for the start cron job operation
func NewStartCronJob(ctx *middleware.Context, handler StartCronJobHandler) *StartCronJob {
	return &StartCronJob{Context: ctx, Handler: handler}
}

/* StartCronJob swagger:route POST /cronjobs/{id}/start Cronjob startCronJob

Start cronjob

*/
type StartCronJob struct {
	Context *middleware.Context
	Handler StartCronJobHandler
}

func (o *StartCronJob) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewStartCronJobParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// StartCronJobBadRequestBody start cron job bad request body
//
// swagger:model StartCronJobBadRequestBody
type StartCronJobBadRequestBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this start cron job bad request body
func (o *StartCronJobBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StartCronJobBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("startCronJobBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this start cron job bad request body based on context it is used
func (o *StartCronJobBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartCronJobBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartCronJobBadRequestBody) UnmarshalBinary(b []byte) error {
	var res StartCronJobBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// StartCronJobForbiddenBody start cron job forbidden body
//
// swagger:model StartCronJobForbiddenBody
type StartCronJobForbiddenBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this start cron job forbidden body
func (o *StartCronJobForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StartCronJobForbiddenBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("startCronJobForbidden"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this start cron job forbidden body based on context it is used
func (o *StartCronJobForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartCronJobForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartCronJobForbiddenBody) UnmarshalBinary(b []byte) error {
	var res StartCronJobForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// StartCronJobNotFoundBody start cron job not found body
//
// swagger:model StartCronJobNotFoundBody
type StartCronJobNotFoundBody struct {

	// error
	// Required: true
	Error *string `json:"error"`

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this start cron job not found body
func (o *StartCronJobNotFoundBody) Validate(formats strfmt.Registry) error {
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

func (o *StartCronJobNotFoundBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("startCronJobNotFound"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

func (o *StartCronJobNotFoundBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("startCronJobNotFound"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this start cron job not found body based on context it is used
func (o *StartCronJobNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartCronJobNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartCronJobNotFoundBody) UnmarshalBinary(b []byte) error {
	var res StartCronJobNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
