// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AzureAccountTestResult azure account test result
//
// swagger:model AzureAccountTestResult
type AzureAccountTestResult struct {

	// detail link
	// Read Only: true
	DetailLink interface{} `json:"detailLink,omitempty"`

	// no permission services
	// Read Only: true
	NoPermissionServices interface{} `json:"noPermissionServices,omitempty"`
}

// Validate validates this azure account test result
func (m *AzureAccountTestResult) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this azure account test result based on context it is used
func (m *AzureAccountTestResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AzureAccountTestResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureAccountTestResult) UnmarshalBinary(b []byte) error {
	var res AzureAccountTestResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
