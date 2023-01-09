// Code generated by go-swagger; DO NOT EDIT.

package crypto

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

// NewDeleteVaultParams creates a new DeleteVaultParams object
//
// There are no default values defined in the spec.
func NewDeleteVaultParams() DeleteVaultParams {

	return DeleteVaultParams{}
}

// DeleteVaultParams contains all the bound params for the delete vault operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteVault
type DeleteVaultParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Key filter
	  In: query
	*/
	Keys []string
	/*Vault name
	  Required: true
	  Pattern: ^[A-Za-z0-9\-_.]{1,32}$
	  In: path
	*/
	Name string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteVaultParams() beforehand.
func (o *DeleteVaultParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qKeys, qhkKeys, _ := qs.GetOK("keys")
	if err := o.bindKeys(qKeys, qhkKeys, route.Formats); err != nil {
		res = append(res, err)
	}

	rName, rhkName, _ := route.Params.GetOK("name")
	if err := o.bindName(rName, rhkName, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindKeys binds and validates array parameter Keys from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *DeleteVaultParams) bindKeys(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvKeys string
	if len(rawData) > 0 {
		qvKeys = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	keysIC := swag.SplitByFormat(qvKeys, "")
	if len(keysIC) == 0 {
		return nil
	}

	var keysIR []string
	for _, keysIV := range keysIC {
		keysI := keysIV

		keysIR = append(keysIR, keysI)
	}

	o.Keys = keysIR

	return nil
}

// bindName binds and validates parameter Name from path.
func (o *DeleteVaultParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Name = raw

	if err := o.validateName(formats); err != nil {
		return err
	}

	return nil
}

// validateName carries on validations for parameter Name
func (o *DeleteVaultParams) validateName(formats strfmt.Registry) error {

	if err := validate.Pattern("name", "path", o.Name, `^[A-Za-z0-9\-_.]{1,32}$`); err != nil {
		return err
	}

	return nil
}
