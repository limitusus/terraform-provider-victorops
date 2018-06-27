// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// OnCallIntervalDuration on call interval duration
// swagger:model onCallIntervalDuration
type OnCallIntervalDuration struct {

	// hours
	Hours float64 `json:"hours,omitempty"`

	// minutes
	Minutes float64 `json:"minutes,omitempty"`
}

// Validate validates this on call interval duration
func (m *OnCallIntervalDuration) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OnCallIntervalDuration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OnCallIntervalDuration) UnmarshalBinary(b []byte) error {
	var res OnCallIntervalDuration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}