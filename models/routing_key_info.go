// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// RoutingKeyInfo Routing keys contain the following fields
// swagger:model RoutingKeyInfo
type RoutingKeyInfo struct {

	// is default
	IsDefault bool `json:"isDefault,omitempty"`

	// routing key
	RoutingKey string `json:"routingKey,omitempty"`

	// targets
	Targets RoutingKeyInfoTargets `json:"targets"`
}

// Validate validates this routing key info
func (m *RoutingKeyInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *RoutingKeyInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RoutingKeyInfo) UnmarshalBinary(b []byte) error {
	var res RoutingKeyInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
