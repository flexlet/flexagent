// Code generated by go-swagger; DO NOT EDIT.

package crypto

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

// CreateVaultHandlerFunc turns a function with the right signature into a create vault handler
type CreateVaultHandlerFunc func(CreateVaultParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateVaultHandlerFunc) Handle(params CreateVaultParams) middleware.Responder {
	return fn(params)
}

// CreateVaultHandler interface for that can handle valid create vault params
type CreateVaultHandler interface {
	Handle(CreateVaultParams) middleware.Responder
}

// NewCreateVault creates a new http.Handler for the create vault operation
func NewCreateVault(ctx *middleware.Context, handler CreateVaultHandler) *CreateVault {
	return &CreateVault{Context: ctx, Handler: handler}
}

/* CreateVault swagger:route POST /crypto/vault/{name} Crypto createVault

Create vault

*/
type CreateVault struct {
	Context *middleware.Context
	Handler CreateVaultHandler
}

func (o *CreateVault) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateVaultParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// CreateVaultBadRequestBody create vault bad request body
//
// swagger:model CreateVaultBadRequestBody
type CreateVaultBadRequestBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this create vault bad request body
func (o *CreateVaultBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateVaultBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("createVaultBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create vault bad request body based on context it is used
func (o *CreateVaultBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateVaultBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateVaultBadRequestBody) UnmarshalBinary(b []byte) error {
	var res CreateVaultBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// CreateVaultForbiddenBody create vault forbidden body
//
// swagger:model CreateVaultForbiddenBody
type CreateVaultForbiddenBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this create vault forbidden body
func (o *CreateVaultForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateVaultForbiddenBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("createVaultForbidden"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create vault forbidden body based on context it is used
func (o *CreateVaultForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateVaultForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateVaultForbiddenBody) UnmarshalBinary(b []byte) error {
	var res CreateVaultForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// CreateVaultNotFoundBody create vault not found body
//
// swagger:model CreateVaultNotFoundBody
type CreateVaultNotFoundBody struct {

	// error
	// Required: true
	Error *string `json:"error"`

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this create vault not found body
func (o *CreateVaultNotFoundBody) Validate(formats strfmt.Registry) error {
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

func (o *CreateVaultNotFoundBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("createVaultNotFound"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

func (o *CreateVaultNotFoundBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("createVaultNotFound"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create vault not found body based on context it is used
func (o *CreateVaultNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateVaultNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateVaultNotFoundBody) UnmarshalBinary(b []byte) error {
	var res CreateVaultNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
