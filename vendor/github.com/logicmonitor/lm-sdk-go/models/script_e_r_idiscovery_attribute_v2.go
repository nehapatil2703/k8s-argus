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

// ScriptERIdiscoveryAttributeV2 script e r idiscovery attribute v2
// swagger:model ScriptERIDiscoveryAttributeV2
type ScriptERIdiscoveryAttributeV2 struct {

	// groovy script
	GroovyScript string `json:"groovyScript,omitempty"`

	// linux cmdline
	LinuxCmdline string `json:"linuxCmdline,omitempty"`

	// linux script
	LinuxScript string `json:"linuxScript,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// type
	Type string `json:"type,omitempty"`

	// win cmdline
	WinCmdline string `json:"winCmdline,omitempty"`

	// win script
	WinScript string `json:"winScript,omitempty"`
}

// Validate validates this script e r idiscovery attribute v2
func (m *ScriptERIdiscoveryAttributeV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScriptERIdiscoveryAttributeV2) validateName(formats strfmt.Registry) error {
	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScriptERIdiscoveryAttributeV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScriptERIdiscoveryAttributeV2) UnmarshalBinary(b []byte) error {
	var res ScriptERIdiscoveryAttributeV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
