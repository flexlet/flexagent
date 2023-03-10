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

// UpdateVaultHandlerFunc turns a function with the right signature into a update vault handler
type UpdateVaultHandlerFunc func(UpdateVaultParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateVaultHandlerFunc) Handle(params UpdateVaultParams) middleware.Responder {
	return fn(params)
}

// UpdateVaultHandler interface for that can handle valid update vault params
type UpdateVaultHandler interface {
	Handle(UpdateVaultParams) middleware.Responder
}

// NewUpdateVault creates a new http.Handler for the update vault operation
func NewUpdateVault(ctx *middleware.Context, handler UpdateVaultHandler) *UpdateVault {
	return &UpdateVault{Context: ctx, Handler: handler}
}

/* UpdateVault swagger:route PUT /crypto/vault/{name} Crypto updateVault

Update vault

*/
type UpdateVault struct {
	Context *middleware.Context
	Handler UpdateVaultHandler
}

func (o *UpdateVault) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateVaultParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// UpdateVaultBadRequestBody update vault bad request body
//
// swagger:model UpdateVaultBadRequestBody
type UpdateVaultBadRequestBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this update vault bad request body
func (o *UpdateVaultBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateVaultBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("updateVaultBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update vault bad request body based on context it is used
func (o *UpdateVaultBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateVaultBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateVaultBadRequestBody) UnmarshalBinary(b []byte) error {
	var res UpdateVaultBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// UpdateVaultForbiddenBody update vault forbidden body
//
// swagger:model UpdateVaultForbiddenBody
type UpdateVaultForbiddenBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this update vault forbidden body
func (o *UpdateVaultForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateVaultForbiddenBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("updateVaultForbidden"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update vault forbidden body based on context it is used
func (o *UpdateVaultForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateVaultForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateVaultForbiddenBody) UnmarshalBinary(b []byte) error {
	var res UpdateVaultForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// UpdateVaultNotFoundBody update vault not found body
//
// swagger:model UpdateVaultNotFoundBody
type UpdateVaultNotFoundBody struct {

	// error
	// Required: true
	Error *string `json:"error"`

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this update vault not found body
func (o *UpdateVaultNotFoundBody) Validate(formats strfmt.Registry) error {
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

func (o *UpdateVaultNotFoundBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("updateVaultNotFound"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

func (o *UpdateVaultNotFoundBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("updateVaultNotFound"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update vault not found body based on context it is used
func (o *UpdateVaultNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateVaultNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateVaultNotFoundBody) UnmarshalBinary(b []byte) error {
	var res UpdateVaultNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
