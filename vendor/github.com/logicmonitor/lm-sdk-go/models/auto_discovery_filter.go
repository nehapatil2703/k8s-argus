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

// AutoDiscoveryFilter auto discovery filter
// swagger:model AutoDiscoveryFilter
type AutoDiscoveryFilter struct {

	// attribute
	// Required: true
	Attribute *string `json:"attribute"`

	// comment
	Comment string `json:"comment,omitempty"`

	// operation
	// Required: true
	Operation *string `json:"operation"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this auto discovery filter
func (m *AutoDiscoveryFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttribute(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutoDiscoveryFilter) validateAttribute(formats strfmt.Registry) error {
	if err := validate.Required("attribute", "body", m.Attribute); err != nil {
		return err
	}

	return nil
}

func (m *AutoDiscoveryFilter) validateOperation(formats strfmt.Registry) error {
	if err := validate.Required("operation", "body", m.Operation); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AutoDiscoveryFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AutoDiscoveryFilter) UnmarshalBinary(b []byte) error {
	var res AutoDiscoveryFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
