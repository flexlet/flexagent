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

// UpdateCronJobHandlerFunc turns a function with the right signature into a update cron job handler
type UpdateCronJobHandlerFunc func(UpdateCronJobParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateCronJobHandlerFunc) Handle(params UpdateCronJobParams) middleware.Responder {
	return fn(params)
}

// UpdateCronJobHandler interface for that can handle valid update cron job params
type UpdateCronJobHandler interface {
	Handle(UpdateCronJobParams) middleware.Responder
}

// NewUpdateCronJob creates a new http.Handler for the update cron job operation
func NewUpdateCronJob(ctx *middleware.Context, handler UpdateCronJobHandler) *UpdateCronJob {
	return &UpdateCronJob{Context: ctx, Handler: handler}
}

/* UpdateCronJob swagger:route PUT /cronjobs/{id} Cronjob updateCronJob

Update cronjob

*/
type UpdateCronJob struct {
	Context *middleware.Context
	Handler UpdateCronJobHandler
}

func (o *UpdateCronJob) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateCronJobParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// UpdateCronJobBadRequestBody update cron job bad request body
//
// swagger:model UpdateCronJobBadRequestBody
type UpdateCronJobBadRequestBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this update cron job bad request body
func (o *UpdateCronJobBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateCronJobBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("updateCronJobBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update cron job bad request body based on context it is used
func (o *UpdateCronJobBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateCronJobBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateCronJobBadRequestBody) UnmarshalBinary(b []byte) error {
	var res UpdateCronJobBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// UpdateCronJobForbiddenBody update cron job forbidden body
//
// swagger:model UpdateCronJobForbiddenBody
type UpdateCronJobForbiddenBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this update cron job forbidden body
func (o *UpdateCronJobForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateCronJobForbiddenBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("updateCronJobForbidden"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update cron job forbidden body based on context it is used
func (o *UpdateCronJobForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateCronJobForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateCronJobForbiddenBody) UnmarshalBinary(b []byte) error {
	var res UpdateCronJobForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// UpdateCronJobNotFoundBody update cron job not found body
//
// swagger:model UpdateCronJobNotFoundBody
type UpdateCronJobNotFoundBody struct {

	// error
	// Required: true
	Error *string `json:"error"`

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this update cron job not found body
func (o *UpdateCronJobNotFoundBody) Validate(formats strfmt.Registry) error {
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

func (o *UpdateCronJobNotFoundBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("updateCronJobNotFound"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

func (o *UpdateCronJobNotFoundBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("updateCronJobNotFound"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update cron job not found body based on context it is used
func (o *UpdateCronJobNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateCronJobNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateCronJobNotFoundBody) UnmarshalBinary(b []byte) error {
	var res UpdateCronJobNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
