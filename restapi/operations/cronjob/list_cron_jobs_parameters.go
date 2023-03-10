// Code generated by go-swagger; DO NOT EDIT.

package cronjob

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewListCronJobsParams creates a new ListCronJobsParams object
//
// There are no default values defined in the spec.
func NewListCronJobsParams() ListCronJobsParams {

	return ListCronJobsParams{}
}

// ListCronJobsParams contains all the bound params for the list cron jobs operation
// typically these are obtained from a http.Request
//
// swagger:parameters listCronJobs
type ListCronJobsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Name filter
	  Pattern: ^[A-Za-z0-9\-_.]{1,32}$
	  In: query
	*/
	Name *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewListCronJobsParams() beforehand.
func (o *ListCronJobsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qName, qhkName, _ := qs.GetOK("name")
	if err := o.bindName(qName, qhkName, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindName binds and validates parameter Name from query.
func (o *ListCronJobsParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Name = &raw

	if err := o.validateName(formats); err != nil {
		return err
	}

	return nil
}

// validateName carries on validations for parameter Name
func (o *ListCronJobsParams) validateName(formats strfmt.Registry) error {

	if err := validate.Pattern("name", "query", *o.Name, `^[A-Za-z0-9\-_.]{1,32}$`); err != nil {
		return err
	}

	return nil
}
