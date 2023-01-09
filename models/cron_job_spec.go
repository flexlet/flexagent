// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CronJobSpec Cronjob spec
//
// swagger:model CronJobSpec
type CronJobSpec struct {

	// Job spec
	// Required: true
	Jobspec *JobSpec `json:"jobspec"`

	// Job name
	// Required: true
	// Pattern: ^[A-Za-z0-9\-._]{1,32}$
	Name string `json:"name"`

	// Cronjob schedule
	// Required: true
	Schedule string `json:"schedule"`
}

// Validate validates this cron job spec
func (m *CronJobSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJobspec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CronJobSpec) validateJobspec(formats strfmt.Registry) error {

	if err := validate.Required("jobspec", "body", m.Jobspec); err != nil {
		return err
	}

	if m.Jobspec != nil {
		if err := m.Jobspec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("jobspec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("jobspec")
			}
			return err
		}
	}

	return nil
}

func (m *CronJobSpec) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", m.Name, `^[A-Za-z0-9\-._]{1,32}$`); err != nil {
		return err
	}

	return nil
}

func (m *CronJobSpec) validateSchedule(formats strfmt.Registry) error {

	if err := validate.RequiredString("schedule", "body", m.Schedule); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this cron job spec based on the context it is used
func (m *CronJobSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateJobspec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CronJobSpec) contextValidateJobspec(ctx context.Context, formats strfmt.Registry) error {

	if m.Jobspec != nil {
		if err := m.Jobspec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("jobspec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("jobspec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CronJobSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CronJobSpec) UnmarshalBinary(b []byte) error {
	var res CronJobSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
