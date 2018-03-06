// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ListRoutingKeysResponse Info about routing keys in the org
// swagger:model ListRoutingKeysResponse
type ListRoutingKeysResponse struct {

	// self Url
	SelfURL string `json:"_selfUrl,omitempty"`

	// routing keys
	RoutingKeys ListRoutingKeysResponseRoutingKeys `json:"routingKeys"`
}

// Validate validates this list routing keys response
func (m *ListRoutingKeysResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ListRoutingKeysResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListRoutingKeysResponse) UnmarshalBinary(b []byte) error {
	var res ListRoutingKeysResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}