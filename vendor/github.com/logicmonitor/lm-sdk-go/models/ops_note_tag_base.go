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

// OpsNoteTagBase ops note tag base
// swagger:model OpsNoteTagBase
type OpsNoteTagBase struct {

	// created on in sec
	// Read Only: true
	CreatedOnInSec int64 `json:"createdOnInSec,omitempty"`

	// id
	// Read Only: true
	ID string `json:"id,omitempty"`

	// release
	// Required: true
	Name *string `json:"name"`

	// update on in sec
	// Read Only: true
	UpdateOnInSec int64 `json:"updateOnInSec,omitempty"`
}

// Validate validates this ops note tag base
func (m *OpsNoteTagBase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpsNoteTagBase) validateName(formats strfmt.Registry) error {
	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpsNoteTagBase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpsNoteTagBase) UnmarshalBinary(b []byte) error {
	var res OpsNoteTagBase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
