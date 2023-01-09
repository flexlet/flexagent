// Code generated by go-swagger; DO NOT EDIT.

package job

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

// SubmitHandlerFunc turns a function with the right signature into a submit handler
type SubmitHandlerFunc func(SubmitParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SubmitHandlerFunc) Handle(params SubmitParams) middleware.Responder {
	return fn(params)
}

// SubmitHandler interface for that can handle valid submit params
type SubmitHandler interface {
	Handle(SubmitParams) middleware.Responder
}

// NewSubmit creates a new http.Handler for the submit operation
func NewSubmit(ctx *middleware.Context, handler SubmitHandler) *Submit {
	return &Submit{Context: ctx, Handler: handler}
}

/* Submit swagger:route POST /jobs Job submit

Submit jobs

*/
type Submit struct {
	Context *middleware.Context
	Handler SubmitHandler
}

func (o *Submit) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewSubmitParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// SubmitBadRequestBody submit bad request body
//
// swagger:model SubmitBadRequestBody
type SubmitBadRequestBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this submit bad request body
func (o *SubmitBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SubmitBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("submitBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this submit bad request body based on context it is used
func (o *SubmitBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SubmitBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SubmitBadRequestBody) UnmarshalBinary(b []byte) error {
	var res SubmitBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// SubmitForbiddenBody submit forbidden body
//
// swagger:model SubmitForbiddenBody
type SubmitForbiddenBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this submit forbidden body
func (o *SubmitForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SubmitForbiddenBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("submitForbidden"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this submit forbidden body based on context it is used
func (o *SubmitForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SubmitForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SubmitForbiddenBody) UnmarshalBinary(b []byte) error {
	var res SubmitForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// SubmitNotFoundBody submit not found body
//
// swagger:model SubmitNotFoundBody
type SubmitNotFoundBody struct {

	// error
	// Required: true
	Error *string `json:"error"`

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this submit not found body
func (o *SubmitNotFoundBody) Validate(formats strfmt.Registry) error {
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

func (o *SubmitNotFoundBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("submitNotFound"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

func (o *SubmitNotFoundBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("submitNotFound"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this submit not found body based on context it is used
func (o *SubmitNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SubmitNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SubmitNotFoundBody) UnmarshalBinary(b []byte) error {
	var res SubmitNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
