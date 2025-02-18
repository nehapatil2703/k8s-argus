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

// Recipient recipient
// swagger:model Recipient
type Recipient struct {

	// the user name if method = admin, or the email address if method = arbitrary
	Addr string `json:"addr,omitempty"`

	// contact
	Contact string `json:"contact,omitempty"`

	// Admin | Arbitrary, where Admin = a user, and Arbitrary = an arbitrary email
	// Required: true
	Method *string `json:"method"`

	// email | sms | voice, where type must be email if method = arbitrary
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this recipient
func (m *Recipient) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Recipient) validateMethod(formats strfmt.Registry) error {
	if err := validate.Required("method", "body", m.Method); err != nil {
		return err
	}

	return nil
}

func (m *Recipient) validateType(formats strfmt.Registry) error {
	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Recipient) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Recipient) UnmarshalBinary(b []byte) error {
	var res Recipient
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
