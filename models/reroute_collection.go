// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RerouteCollection reroute collection
// swagger:model RerouteCollection
type RerouteCollection struct {

	// reroutes
	// Required: true
	Reroutes RerouteCollectionReroutes `json:"reroutes"`

	// user name
	// Required: true
	UserName *string `json:"userName"`
}

// Validate validates this reroute collection
func (m *RerouteCollection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReroutes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUserName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RerouteCollection) validateReroutes(formats strfmt.Registry) error {

	if err := validate.Required("reroutes", "body", m.Reroutes); err != nil {
		return err
	}

	if err := m.Reroutes.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("reroutes")
		}
		return err
	}

	return nil
}

func (m *RerouteCollection) validateUserName(formats strfmt.Registry) error {

	if err := validate.Required("userName", "body", m.UserName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RerouteCollection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RerouteCollection) UnmarshalBinary(b []byte) error {
	var res RerouteCollection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
