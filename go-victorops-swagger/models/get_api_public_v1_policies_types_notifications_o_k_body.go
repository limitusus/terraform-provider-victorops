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

// GetAPIPublicV1PoliciesTypesNotificationsOKBody get Api public v1 policies types notifications o k body
// swagger:model getApiPublicV1PoliciesTypesNotificationsOKBody
type GetAPIPublicV1PoliciesTypesNotificationsOKBody struct {

	// self Url
	SelfURL string `json:"_selfUrl,omitempty"`

	// notification types
	NotificationTypes []*NotificationObject `json:"notificationTypes"`
}

// Validate validates this get Api public v1 policies types notifications o k body
func (m *GetAPIPublicV1PoliciesTypesNotificationsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNotificationTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetAPIPublicV1PoliciesTypesNotificationsOKBody) validateNotificationTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.NotificationTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.NotificationTypes); i++ {
		if swag.IsZero(m.NotificationTypes[i]) { // not required
			continue
		}

		if m.NotificationTypes[i] != nil {
			if err := m.NotificationTypes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("notificationTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetAPIPublicV1PoliciesTypesNotificationsOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetAPIPublicV1PoliciesTypesNotificationsOKBody) UnmarshalBinary(b []byte) error {
	var res GetAPIPublicV1PoliciesTypesNotificationsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
