// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TeamSchedule team schedule
// swagger:model TeamSchedule
type TeamSchedule struct {

	// schedules
	Schedules TeamScheduleSchedules `json:"schedules"`

	// team
	Team *Team `json:"team,omitempty"`
}

// Validate validates this team schedule
func (m *TeamSchedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTeam(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TeamSchedule) validateTeam(formats strfmt.Registry) error {

	if swag.IsZero(m.Team) { // not required
		return nil
	}

	if m.Team != nil {

		if err := m.Team.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("team")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TeamSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TeamSchedule) UnmarshalBinary(b []byte) error {
	var res TeamSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
