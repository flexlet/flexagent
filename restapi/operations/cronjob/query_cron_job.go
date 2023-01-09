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

// QueryCronJobHandlerFunc turns a function with the right signature into a query cron job handler
type QueryCronJobHandlerFunc func(QueryCronJobParams) middleware.Responder

// Handle executing the request and returning a response
func (fn QueryCronJobHandlerFunc) Handle(params QueryCronJobParams) middleware.Responder {
	return fn(params)
}

// QueryCronJobHandler interface for that can handle valid query cron job params
type QueryCronJobHandler interface {
	Handle(QueryCronJobParams) middleware.Responder
}

// NewQueryCronJob creates a new http.Handler for the query cron job operation
func NewQueryCronJob(ctx *middleware.Context, handler QueryCronJobHandler) *QueryCronJob {
	return &QueryCronJob{Context: ctx, Handler: handler}
}

/* QueryCronJob swagger:route GET /cronjobs/{id} Cronjob queryCronJob

Query cronjob

*/
type QueryCronJob struct {
	Context *middleware.Context
	Handler QueryCronJobHandler
}

func (o *QueryCronJob) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewQueryCronJobParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// QueryCronJobBadRequestBody query cron job bad request body
//
// swagger:model QueryCronJobBadRequestBody
type QueryCronJobBadRequestBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this query cron job bad request body
func (o *QueryCronJobBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QueryCronJobBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("queryCronJobBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this query cron job bad request body based on context it is used
func (o *QueryCronJobBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *QueryCronJobBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QueryCronJobBadRequestBody) UnmarshalBinary(b []byte) error {
	var res QueryCronJobBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// QueryCronJobForbiddenBody query cron job forbidden body
//
// swagger:model QueryCronJobForbiddenBody
type QueryCronJobForbiddenBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this query cron job forbidden body
func (o *QueryCronJobForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QueryCronJobForbiddenBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("queryCronJobForbidden"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this query cron job forbidden body based on context it is used
func (o *QueryCronJobForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *QueryCronJobForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QueryCronJobForbiddenBody) UnmarshalBinary(b []byte) error {
	var res QueryCronJobForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// QueryCronJobNotFoundBody query cron job not found body
//
// swagger:model QueryCronJobNotFoundBody
type QueryCronJobNotFoundBody struct {

	// error
	// Required: true
	Error *string `json:"error"`

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this query cron job not found body
func (o *QueryCronJobNotFoundBody) Validate(formats strfmt.Registry) error {
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

func (o *QueryCronJobNotFoundBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("queryCronJobNotFound"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

func (o *QueryCronJobNotFoundBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("queryCronJobNotFound"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this query cron job not found body based on context it is used
func (o *QueryCronJobNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *QueryCronJobNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QueryCronJobNotFoundBody) UnmarshalBinary(b []byte) error {
	var res QueryCronJobNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
