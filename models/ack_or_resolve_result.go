// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AckOrResolveResult The results of ack/resolve for a individual incident
// swagger:model AckOrResolveResult
type AckOrResolveResult struct {

	// cmd accepted
	CmdAccepted bool `json:"cmdAccepted,omitempty"`

	// entity Id
	EntityID string `json:"entityId,omitempty"`

	// incident number
	IncidentNumber string `json:"incidentNumber,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this ack or resolve result
func (m *AckOrResolveResult) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AckOrResolveResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AckOrResolveResult) UnmarshalBinary(b []byte) error {
	var res AckOrResolveResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}