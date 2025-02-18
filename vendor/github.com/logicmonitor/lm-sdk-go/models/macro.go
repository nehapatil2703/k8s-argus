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

// Macro macro
// swagger:model Macro
type Macro struct {

	// description
	Description string `json:"description,omitempty"`

	// value
	// Required: true
	Value *string `json:"value"`

	// variable
	// Required: true
	Variable *string `json:"variable"`
}

// Validate validates this macro
func (m *Macro) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVariable(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Macro) validateValue(formats strfmt.Registry) error {
	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

func (m *Macro) validateVariable(formats strfmt.Registry) error {
	if err := validate.Required("variable", "body", m.Variable); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Macro) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Macro) UnmarshalBinary(b []byte) error {
	var res Macro
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
