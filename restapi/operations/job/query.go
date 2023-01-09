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

// QueryHandlerFunc turns a function with the right signature into a query handler
type QueryHandlerFunc func(QueryParams) middleware.Responder

// Handle executing the request and returning a response
func (fn QueryHandlerFunc) Handle(params QueryParams) middleware.Responder {
	return fn(params)
}

// QueryHandler interface for that can handle valid query params
type QueryHandler interface {
	Handle(QueryParams) middleware.Responder
}

// NewQuery creates a new http.Handler for the query operation
func NewQuery(ctx *middleware.Context, handler QueryHandler) *Query {
	return &Query{Context: ctx, Handler: handler}
}

/* Query swagger:route GET /jobs/{urn} Job query

Query job

*/
type Query struct {
	Context *middleware.Context
	Handler QueryHandler
}

func (o *Query) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewQueryParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// QueryBadRequestBody query bad request body
//
// swagger:model QueryBadRequestBody
type QueryBadRequestBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this query bad request body
func (o *QueryBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QueryBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("queryBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this query bad request body based on context it is used
func (o *QueryBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *QueryBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QueryBadRequestBody) UnmarshalBinary(b []byte) error {
	var res QueryBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// QueryForbiddenBody query forbidden body
//
// swagger:model QueryForbiddenBody
type QueryForbiddenBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this query forbidden body
func (o *QueryForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QueryForbiddenBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("queryForbidden"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this query forbidden body based on context it is used
func (o *QueryForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *QueryForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QueryForbiddenBody) UnmarshalBinary(b []byte) error {
	var res QueryForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// QueryNotFoundBody query not found body
//
// swagger:model QueryNotFoundBody
type QueryNotFoundBody struct {

	// error
	// Required: true
	Error *string `json:"error"`

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this query not found body
func (o *QueryNotFoundBody) Validate(formats strfmt.Registry) error {
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

func (o *QueryNotFoundBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("queryNotFound"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

func (o *QueryNotFoundBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("queryNotFound"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this query not found body based on context it is used
func (o *QueryNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *QueryNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QueryNotFoundBody) UnmarshalBinary(b []byte) error {
	var res QueryNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
