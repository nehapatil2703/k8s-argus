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

// TableWidgetInstanceCell table widget instance cell
// swagger:model TableWidgetInstanceCell
type TableWidgetInstanceCell struct {

	// data point Id
	// Read Only: true
	DataPointID int32 `json:"dataPointId,omitempty"`

	// data point name
	// Read Only: true
	DataPointName string `json:"dataPointName,omitempty"`

	// instance Id
	// Required: true
	InstanceID *int32 `json:"instanceId"`

	// instance name
	// Read Only: true
	InstanceName string `json:"instanceName,omitempty"`

	// validation status code
	// Read Only: true
	ValidationStatusCode int32 `json:"validationStatusCode,omitempty"`
}

// Validate validates this table widget instance cell
func (m *TableWidgetInstanceCell) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstanceID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TableWidgetInstanceCell) validateInstanceID(formats strfmt.Registry) error {
	if err := validate.Required("instanceId", "body", m.InstanceID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TableWidgetInstanceCell) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TableWidgetInstanceCell) UnmarshalBinary(b []byte) error {
	var res TableWidgetInstanceCell
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
