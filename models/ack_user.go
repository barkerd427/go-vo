// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AckUser A JSON object with first/last name and email of the user that acknowledged the incident.
// swagger:model AckUser
type AckUser struct {

	// email
	Email string `json:"email,omitempty"`

	// first
	First string `json:"first,omitempty"`

	// last
	Last string `json:"last,omitempty"`
}

// Validate validates this ack user
func (m *AckUser) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AckUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AckUser) UnmarshalBinary(b []byte) error {
	var res AckUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
