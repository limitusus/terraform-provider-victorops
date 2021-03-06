// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ActiveIncidentInfo Incidents contain the following fields (all should be considered optional)
// swagger:model ActiveIncidentInfo
type ActiveIncidentInfo struct {

	// The number of alerts received for this incident
	AlertCount float64 `json:"alertCount,omitempty"`

	// The current phase of the incident "resolved", "triggered" or "acknowledged".
	CurrentPhase string `json:"currentPhase,omitempty"`

	// The unique identification of the entity being monitored that caused the incident
	EntityID string `json:"entityId,omitempty"`

	// The host on which the incident occurred
	Host string `json:"host,omitempty"`

	// The VictorOps incident number
	IncidentNumber string `json:"incidentNumber,omitempty"`

	// The unique id of the last alert for the incident
	LastAlertID string `json:"lastAlertId,omitempty"`

	// The time of the last alert received for the incident
	LastAlertTime string `json:"lastAlertTime,omitempty"`

	// The escalation policies that were triggered for the incident
	PagedPolicies []*EscalationPolicyInfo `json:"pagedPolicies"`

	// The teams that were paged for the incident
	PagedTeams []string `json:"pagedTeams"`

	// The users that were paged for the incident.
	PagedUsers []string `json:"pagedUsers"`

	// The service name causing the incident (if any)
	Service string `json:"service,omitempty"`

	// The time the incident started
	StartTime string `json:"startTime,omitempty"`

	// Transitions of the incident state over time
	Transitions []*IncidentTransition `json:"transitions"`
}

// Validate validates this active incident info
func (m *ActiveIncidentInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePagedPolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransitions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActiveIncidentInfo) validatePagedPolicies(formats strfmt.Registry) error {

	if swag.IsZero(m.PagedPolicies) { // not required
		return nil
	}

	for i := 0; i < len(m.PagedPolicies); i++ {
		if swag.IsZero(m.PagedPolicies[i]) { // not required
			continue
		}

		if m.PagedPolicies[i] != nil {
			if err := m.PagedPolicies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pagedPolicies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ActiveIncidentInfo) validateTransitions(formats strfmt.Registry) error {

	if swag.IsZero(m.Transitions) { // not required
		return nil
	}

	for i := 0; i < len(m.Transitions); i++ {
		if swag.IsZero(m.Transitions[i]) { // not required
			continue
		}

		if m.Transitions[i] != nil {
			if err := m.Transitions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("transitions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ActiveIncidentInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActiveIncidentInfo) UnmarshalBinary(b []byte) error {
	var res ActiveIncidentInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
