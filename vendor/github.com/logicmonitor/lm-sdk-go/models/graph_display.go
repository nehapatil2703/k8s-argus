// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GraphDisplay graph display
// swagger:model GraphDisplay
type GraphDisplay struct {

	// color
	Color string `json:"color,omitempty"`

	// legend
	Legend string `json:"legend,omitempty"`

	// option
	Option string `json:"option,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this graph display
func (m *GraphDisplay) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GraphDisplay) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GraphDisplay) UnmarshalBinary(b []byte) error {
	var res GraphDisplay
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
