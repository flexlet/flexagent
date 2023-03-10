// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewQueryParams creates a new QueryParams object
// with the default values initialized.
func NewQueryParams() QueryParams {

	var (
		// initialize parameters with default values

		outputLineLimitDefault = int32(10)
	)

	return QueryParams{
		OutputLineLimit: &outputLineLimitDefault,
	}
}

// QueryParams contains all the bound params for the query operation
// typically these are obtained from a http.Request
//
// swagger:parameters query
type QueryParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Output lines limit
	  Maximum: 100
	  Minimum: 0
	  In: query
	  Default: 10
	*/
	OutputLineLimit *int32
	/*Output line start
	  Minimum: 1
	  In: query
	*/
	OutputLineStart *int32
	/*Job urn
	  Required: true
	  Pattern: ^[A-Za-z0-9\-._]{1,32}:jobs:[A-Za-z0-9\-._]{1,32}:[A-Za-z0-9\-]{36}$
	  In: path
	*/
	Urn string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewQueryParams() beforehand.
func (o *QueryParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qOutputLineLimit, qhkOutputLineLimit, _ := qs.GetOK("outputLineLimit")
	if err := o.bindOutputLineLimit(qOutputLineLimit, qhkOutputLineLimit, route.Formats); err != nil {
		res = append(res, err)
	}

	qOutputLineStart, qhkOutputLineStart, _ := qs.GetOK("outputLineStart")
	if err := o.bindOutputLineStart(qOutputLineStart, qhkOutputLineStart, route.Formats); err != nil {
		res = append(res, err)
	}

	rUrn, rhkUrn, _ := route.Params.GetOK("urn")
	if err := o.bindUrn(rUrn, rhkUrn, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindOutputLineLimit binds and validates parameter OutputLineLimit from query.
func (o *QueryParams) bindOutputLineLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewQueryParams()
		return nil
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("outputLineLimit", "query", "int32", raw)
	}
	o.OutputLineLimit = &value

	if err := o.validateOutputLineLimit(formats); err != nil {
		return err
	}

	return nil
}

// validateOutputLineLimit carries on validations for parameter OutputLineLimit
func (o *QueryParams) validateOutputLineLimit(formats strfmt.Registry) error {

	if err := validate.MinimumInt("outputLineLimit", "query", int64(*o.OutputLineLimit), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("outputLineLimit", "query", int64(*o.OutputLineLimit), 100, false); err != nil {
		return err
	}

	return nil
}

// bindOutputLineStart binds and validates parameter OutputLineStart from query.
func (o *QueryParams) bindOutputLineStart(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("outputLineStart", "query", "int32", raw)
	}
	o.OutputLineStart = &value

	if err := o.validateOutputLineStart(formats); err != nil {
		return err
	}

	return nil
}

// validateOutputLineStart carries on validations for parameter OutputLineStart
func (o *QueryParams) validateOutputLineStart(formats strfmt.Registry) error {

	if err := validate.MinimumInt("outputLineStart", "query", int64(*o.OutputLineStart), 1, false); err != nil {
		return err
	}

	return nil
}

// bindUrn binds and validates parameter Urn from path.
func (o *QueryParams) bindUrn(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Urn = raw

	if err := o.validateUrn(formats); err != nil {
		return err
	}

	return nil
}

// validateUrn carries on validations for parameter Urn
func (o *QueryParams) validateUrn(formats strfmt.Registry) error {

	if err := validate.Pattern("urn", "path", o.Urn, `^[A-Za-z0-9\-._]{1,32}:jobs:[A-Za-z0-9\-._]{1,32}:[A-Za-z0-9\-]{36}$`); err != nil {
		return err
	}

	return nil
}
