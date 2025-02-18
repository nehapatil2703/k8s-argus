// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AuditLog audit log
// swagger:model AuditLog
type AuditLog struct {

	// The description of the action recorded in the access log entry
	// Read Only: true
	Description string `json:"description,omitempty"`

	// The time, in epoch seconds, that the action recorded in the access log entry occurred
	// Read Only: true
	HappenedOn int64 `json:"happenedOn,omitempty"`

	// The date and time that the action recorded in the access log entry occured
	// Read Only: true
	HappenedOnLocal string `json:"happenedOnLocal,omitempty"`

	// The Id of the access log entry
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The IP address that the action was performed from
	// Read Only: true
	IP string `json:"ip,omitempty"`

	// The Id of the session during which the action was performed
	// Read Only: true
	SessionID string `json:"sessionId,omitempty"`

	// The username associated with the user that performed the action recorded in the access log entry
	// Read Only: true
	Username string `json:"username,omitempty"`
}

// Validate validates this audit log
func (m *AuditLog) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuditLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditLog) UnmarshalBinary(b []byte) error {
	var res AuditLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
