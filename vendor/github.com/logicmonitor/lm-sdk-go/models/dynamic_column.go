// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DynamicColumn dynamic column
// swagger:model DynamicColumn
type DynamicColumn struct {

	// is hidden
	IsHidden bool `json:"isHidden,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this dynamic column
func (m *DynamicColumn) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DynamicColumn) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DynamicColumn) UnmarshalBinary(b []byte) error {
	var res DynamicColumn
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
