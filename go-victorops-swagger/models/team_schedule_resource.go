// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TeamScheduleResource team schedule resource
// swagger:model TeamScheduleResource
type TeamScheduleResource struct {

	// The user scheduled on call (if any)
	Oncall string `json:"oncall,omitempty"`

	// The user overriding the scheduled on call user (if any)
	Overrideoncall string `json:"overrideoncall,omitempty"`

	// policy type
	// Required: true
	PolicyType *string `json:"policyType"`

	// rolls
	// Required: true
	Rolls []*TeamScheduleRollResource `json:"rolls"`

	// rotation name
	RotationName string `json:"rotationName,omitempty"`

	// shift name
	ShiftName string `json:"shiftName,omitempty"`

	// shift roll
	ShiftRoll float64 `json:"shiftRoll,omitempty"`
}

// Validate validates this team schedule resource
func (m *TeamScheduleResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePolicyType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRolls(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TeamScheduleResource) validatePolicyType(formats strfmt.Registry) error {

	if err := validate.Required("policyType", "body", m.PolicyType); err != nil {
		return err
	}

	return nil
}

func (m *TeamScheduleResource) validateRolls(formats strfmt.Registry) error {

	if err := validate.Required("rolls", "body", m.Rolls); err != nil {
		return err
	}

	for i := 0; i < len(m.Rolls); i++ {
		if swag.IsZero(m.Rolls[i]) { // not required
			continue
		}

		if m.Rolls[i] != nil {
			if err := m.Rolls[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rolls" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TeamScheduleResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TeamScheduleResource) UnmarshalBinary(b []byte) error {
	var res TeamScheduleResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}