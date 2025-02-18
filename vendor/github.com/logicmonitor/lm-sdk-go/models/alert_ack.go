// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AlertAck alert ack
// swagger:model AlertAck
type AlertAck struct {

	// your comment on the alert
	// Required: true
	AckComment *string `json:"ackComment"`
}

// Validate validates this alert ack
func (m *AlertAck) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAckComment(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertAck) validateAckComment(formats strfmt.Registry) error {
	if err := validate.Required("ackComment", "body", m.AckComment); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AlertAck) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertAck) UnmarshalBinary(b []byte) error {
	var res AlertAck
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
