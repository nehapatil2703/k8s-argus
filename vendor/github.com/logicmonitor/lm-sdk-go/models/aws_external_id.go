// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AwsExternalID aws external Id
// swagger:model AwsExternalId
type AwsExternalID struct {

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// external Id
	ExternalID string `json:"externalId,omitempty"`
}

// Validate validates this aws external Id
func (m *AwsExternalID) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AwsExternalID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsExternalID) UnmarshalBinary(b []byte) error {
	var res AwsExternalID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
