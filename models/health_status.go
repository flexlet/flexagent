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

// HealthStatus Node healthy status
//
// swagger:model HealthStatus
type HealthStatus struct {

	// Plugin health probes
	// Required: true
	Probes map[string]map[string]HealthProbe `json:"probes"`

	// Node healthy status
	// Required: true
	// Enum: [unknown healthy warning critical]
	Status string `json:"status"`
}

// Validate validates this health status
func (m *HealthStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProbes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HealthStatus) validateProbes(formats strfmt.Registry) error {

	if err := validate.Required("probes", "body", m.Probes); err != nil {
		return err
	}

	for k := range m.Probes {

		if err := validate.Required("probes"+"."+k, "body", m.Probes[k]); err != nil {
			return err
		}

		if err := validate.Required("probes"+"."+k, "body", m.Probes); err != nil {
			return err
		}

		for kk := range m.Probes[k] {

			if err := validate.Required("probes"+"."+k+"."+kk, "body", m.Probes[k][kk]); err != nil {
				return err
			}
			if val, ok := m.Probes[k][kk]; ok {
				if err := val.Validate(formats); err != nil {
					if ve, ok := err.(*errors.Validation); ok {
						return ve.ValidateName("probes" + "." + k + "." + kk)
					} else if ce, ok := err.(*errors.CompositeError); ok {
						return ce.ValidateName("probes" + "." + k + "." + kk)
					}
					return err
				}
			}

		}

	}

	return nil
}

var healthStatusTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["unknown","healthy","warning","critical"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		healthStatusTypeStatusPropEnum = append(healthStatusTypeStatusPropEnum, v)
	}
}

const (

	// HealthStatusStatusUnknown captures enum value "unknown"
	HealthStatusStatusUnknown string = "unknown"

	// HealthStatusStatusHealthy captures enum value "healthy"
	HealthStatusStatusHealthy string = "healthy"

	// HealthStatusStatusWarning captures enum value "warning"
	HealthStatusStatusWarning string = "warning"

	// HealthStatusStatusCritical captures enum value "critical"
	HealthStatusStatusCritical string = "critical"
)

// prop value enum
func (m *HealthStatus) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, healthStatusTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HealthStatus) validateStatus(formats strfmt.Registry) error {

	if err := validate.RequiredString("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this health status based on the context it is used
func (m *HealthStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProbes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HealthStatus) contextValidateProbes(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.Required("probes", "body", m.Probes); err != nil {
		return err
	}

	for k := range m.Probes {

		if err := validate.Required("probes"+"."+k, "body", m.Probes); err != nil {
			return err
		}

		for kk := range m.Probes[k] {

			if val, ok := m.Probes[k][kk]; ok {
				if err := val.ContextValidate(ctx, formats); err != nil {
					return err
				}
			}

		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HealthStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HealthStatus) UnmarshalBinary(b []byte) error {
	var res HealthStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}