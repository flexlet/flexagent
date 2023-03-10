// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"flexagent/models"
)

// NewSubmitParams creates a new SubmitParams object
// with the default values initialized.
func NewSubmitParams() SubmitParams {

	var (
		// initialize parameters with default values

		waitDefault = bool(false)
	)

	return SubmitParams{
		Wait: &waitDefault,
	}
}

// SubmitParams contains all the bound params for the submit operation
// typically these are obtained from a http.Request
//
// swagger:parameters submit
type SubmitParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Job spec
	  Required: true
	  In: body
	*/
	Spec []*models.JobSpec
	/*Wait until job finished
	  In: query
	  Default: false
	*/
	Wait *bool
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewSubmitParams() beforehand.
func (o *SubmitParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body []*models.JobSpec
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("spec", "body", ""))
			} else {
				res = append(res, errors.NewParseError("spec", "body", "", err))
			}
		} else {

			// validate array of body objects
			for i := range body {
				if body[i] == nil {
					continue
				}
				if err := body[i].Validate(route.Formats); err != nil {
					res = append(res, err)
					break
				}
			}

			if len(res) == 0 {
				o.Spec = body
			}
		}
	} else {
		res = append(res, errors.Required("spec", "body", ""))
	}

	qWait, qhkWait, _ := qs.GetOK("wait")
	if err := o.bindWait(qWait, qhkWait, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindWait binds and validates parameter Wait from query.
func (o *SubmitParams) bindWait(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewSubmitParams()
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("wait", "query", "bool", raw)
	}
	o.Wait = &value

	return nil
}
