// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AlertBatchResponse A collection of requested alerts
// swagger:model AlertBatchResponse
type AlertBatchResponse struct {

	// alerts
	Alerts AlertBatchResponseAlerts `json:"alerts"`

	// request Id
	RequestID string `json:"requestId,omitempty"`

	// request size
	RequestSize int64 `json:"requestSize,omitempty"`

	// response size
	ResponseSize int64 `json:"responseSize,omitempty"`
}

// Validate validates this alert batch response
func (m *AlertBatchResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AlertBatchResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertBatchResponse) UnmarshalBinary(b []byte) error {
	var res AlertBatchResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
