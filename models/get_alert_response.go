// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetAlertResponse Alert details. All fields should be considered optional.
// swagger:model GetAlertResponse
type GetAlertResponse struct {

	// The user that acknowledged the incident.
	AckAuthor string `json:"ackAuthor,omitempty"`

	// A user entered comment for the acknowledgment.
	AckMsg string `json:"ackMsg,omitempty"`

	// Used within VictorOps to display a human-readable name for the entity.
	EntityDisplayName string `json:"entityDisplayName,omitempty"`

	// Identifies the entity (host, service, etc.) this alert was about.
	//
	EntityID string `json:"entityId,omitempty"`

	// The type of alert; INFO, WARNING, ACKNOWLEDGEMENT, CRITICAL, RECOVERY
	//
	MessageType string `json:"messageType,omitempty"`

	// The name of the monitoring system software (eg. nagios, icinga, sensu, etc.)
	MonitoringTool string `json:"monitoringTool,omitempty"`

	// The full, raw alert details JSON string (i.e. parse the string into a JSON object)
	//
	Raw string `json:"raw,omitempty"`

	// Any additional status information from the alert item.
	StateMessage string `json:"stateMessage,omitempty"`

	// The time this entity entered its current state (seconds since epoch).
	StateStartTime float64 `json:"stateStartTime,omitempty"`

	// Timestamp of the alert in seconds since epoch.
	Timestamp float64 `json:"timestamp,omitempty"`
}

// Validate validates this get alert response
func (m *GetAlertResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetAlertResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetAlertResponse) UnmarshalBinary(b []byte) error {
	var res GetAlertResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
