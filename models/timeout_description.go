// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// TimeoutDescription timeout description
// swagger:model TimeoutDescription
type TimeoutDescription string

const (
	// TimeoutDescriptionIfStillUnackedAfter1Minute captures enum value "If still unacked after 1 minute"
	TimeoutDescriptionIfStillUnackedAfter1Minute TimeoutDescription = "If still unacked after 1 minute"
	// TimeoutDescriptionIfStillUnackedAfter5Minute captures enum value "If still unacked after 5 minute"
	TimeoutDescriptionIfStillUnackedAfter5Minute TimeoutDescription = "If still unacked after 5 minute"
	// TimeoutDescriptionIfStillUnackedAfter10Minute captures enum value "If still unacked after 10 minute"
	TimeoutDescriptionIfStillUnackedAfter10Minute TimeoutDescription = "If still unacked after 10 minute"
	// TimeoutDescriptionIfStillUnackedAfter15Minute captures enum value "If still unacked after 15 minute"
	TimeoutDescriptionIfStillUnackedAfter15Minute TimeoutDescription = "If still unacked after 15 minute"
	// TimeoutDescriptionIfStillUnackedAfter20Minute captures enum value "If still unacked after 20 minute"
	TimeoutDescriptionIfStillUnackedAfter20Minute TimeoutDescription = "If still unacked after 20 minute"
	// TimeoutDescriptionIfStillUnackedAfter25Minute captures enum value "If still unacked after 25 minute"
	TimeoutDescriptionIfStillUnackedAfter25Minute TimeoutDescription = "If still unacked after 25 minute"
	// TimeoutDescriptionIfStillUnackedAfter30Minute captures enum value "If still unacked after 30 minute"
	TimeoutDescriptionIfStillUnackedAfter30Minute TimeoutDescription = "If still unacked after 30 minute"
	// TimeoutDescriptionIfStillUnackedAfter45Minute captures enum value "If still unacked after 45 minute"
	TimeoutDescriptionIfStillUnackedAfter45Minute TimeoutDescription = "If still unacked after 45 minute"
	// TimeoutDescriptionIfStillUnackedAfter60Minute captures enum value "If still unacked after 60 minute"
	TimeoutDescriptionIfStillUnackedAfter60Minute TimeoutDescription = "If still unacked after 60 minute"
)

// for schema
var timeoutDescriptionEnum []interface{}

func init() {
	var res []TimeoutDescription
	if err := json.Unmarshal([]byte(`["If still unacked after 1 minute","If still unacked after 5 minute","If still unacked after 10 minute","If still unacked after 15 minute","If still unacked after 20 minute","If still unacked after 25 minute","If still unacked after 30 minute","If still unacked after 45 minute","If still unacked after 60 minute"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		timeoutDescriptionEnum = append(timeoutDescriptionEnum, v)
	}
}

func (m TimeoutDescription) validateTimeoutDescriptionEnum(path, location string, value TimeoutDescription) error {
	if err := validate.Enum(path, location, value, timeoutDescriptionEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this timeout description
func (m TimeoutDescription) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTimeoutDescriptionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
