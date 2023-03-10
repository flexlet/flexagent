// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"

	"flexagent/models"
)

// NewInputParams creates a new InputParams object
//
// There are no default values defined in the spec.
func NewInputParams() InputParams {

	return InputParams{}
}

// InputParams contains all the bound params for the input operation
// typically these are obtained from a http.Request
//
// swagger:parameters input
type InputParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Input data
	  Required: true
	  In: body
	*/
	Input *models.JobInput
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
// To ensure default values, the struct must have been initialized with NewInputParams() beforehand.
func (o *InputParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.JobInput
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("input", "body", ""))
			} else {
				res = append(res, errors.NewParseError("input", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Input = &body
			}
		}
	} else {
		res = append(res, errors.Required("input", "body", ""))
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

// bindUrn binds and validates parameter Urn from path.
func (o *InputParams) bindUrn(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *InputParams) validateUrn(formats strfmt.Registry) error {

	if err := validate.Pattern("urn", "path", o.Urn, `^[A-Za-z0-9\-._]{1,32}:jobs:[A-Za-z0-9\-._]{1,32}:[A-Za-z0-9\-]{36}$`); err != nil {
		return err
	}

	return nil
}
