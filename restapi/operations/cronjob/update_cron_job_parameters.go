// Code generated by go-swagger; DO NOT EDIT.

package cronjob

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

// NewUpdateCronJobParams creates a new UpdateCronJobParams object
//
// There are no default values defined in the spec.
func NewUpdateCronJobParams() UpdateCronJobParams {

	return UpdateCronJobParams{}
}

// UpdateCronJobParams contains all the bound params for the update cron job operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateCronJob
type UpdateCronJobParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Cronjob id
	  Required: true
	  Pattern: ^[A-Za-z0-9\-]{36}$
	  In: path
	*/
	ID string
	/*Cronjob spec
	  Required: true
	  In: body
	*/
	Spec *models.CronJobSpec
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateCronJobParams() beforehand.
func (o *UpdateCronJobParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.CronJobSpec
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("spec", "body", ""))
			} else {
				res = append(res, errors.NewParseError("spec", "body", "", err))
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
				o.Spec = &body
			}
		}
	} else {
		res = append(res, errors.Required("spec", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindID binds and validates parameter ID from path.
func (o *UpdateCronJobParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ID = raw

	if err := o.validateID(formats); err != nil {
		return err
	}

	return nil
}

// validateID carries on validations for parameter ID
func (o *UpdateCronJobParams) validateID(formats strfmt.Registry) error {

	if err := validate.Pattern("id", "path", o.ID, `^[A-Za-z0-9\-]{36}$`); err != nil {
		return err
	}

	return nil
}
