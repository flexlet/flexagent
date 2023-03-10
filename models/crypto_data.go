// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CryptoData Data to encrypt or decrypt
//
// swagger:model CryptoData
type CryptoData struct {

	// Data dictionary
	// Required: true
	Data map[string]string `json:"data"`

	// Data format
	// Required: true
	// Enum: [base64 raw]
	Format string `json:"format"`
}

// Validate validates this crypto data
func (m *CryptoData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFormat(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CryptoData) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	return nil
}

var cryptoDataTypeFormatPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["base64","raw"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cryptoDataTypeFormatPropEnum = append(cryptoDataTypeFormatPropEnum, v)
	}
}

const (

	// CryptoDataFormatBase64 captures enum value "base64"
	CryptoDataFormatBase64 string = "base64"

	// CryptoDataFormatRaw captures enum value "raw"
	CryptoDataFormatRaw string = "raw"
)

// prop value enum
func (m *CryptoData) validateFormatEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cryptoDataTypeFormatPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CryptoData) validateFormat(formats strfmt.Registry) error {

	if err := validate.RequiredString("format", "body", m.Format); err != nil {
		return err
	}

	// value enum
	if err := m.validateFormatEnum("format", "body", m.Format); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this crypto data based on context it is used
func (m *CryptoData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CryptoData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CryptoData) UnmarshalBinary(b []byte) error {
	var res CryptoData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
