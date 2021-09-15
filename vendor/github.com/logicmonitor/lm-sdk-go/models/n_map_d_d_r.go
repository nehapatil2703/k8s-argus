// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NMapDDR n map d d r
// swagger:model NMapDDR
type NMapDDR struct {

	// Information about how discovered devices are included in or excluded from monitoring. Assignment objects can be created for different types of discovered devices
	Assignment []*Assignment `json:"assignment,omitempty"`

	// change name
	ChangeName string `json:"changeName,omitempty"`
}

// Validate validates this n map d d r
func (m *NMapDDR) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssignment(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NMapDDR) validateAssignment(formats strfmt.Registry) error {
	if swag.IsZero(m.Assignment) { // not required
		return nil
	}

	for i := 0; i < len(m.Assignment); i++ {
		if swag.IsZero(m.Assignment[i]) { // not required
			continue
		}

		if m.Assignment[i] != nil {
			if err := m.Assignment[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("assignment" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NMapDDR) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NMapDDR) UnmarshalBinary(b []byte) error {
	var res NMapDDR
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
