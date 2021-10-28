// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RawDataValues raw data values
//
// swagger:model RawDataValues
type RawDataValues struct {

	// next page params
	// Read Only: true
	NextPageParams string `json:"nextPageParams,omitempty"`

	// time
	// Read Only: true
	Time []int64 `json:"time,omitempty"`

	// values
	// Read Only: true
	Values [][]interface{} `json:"values,omitempty"`
}

// Validate validates this raw data values
func (m *RawDataValues) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this raw data values based on the context it is used
func (m *RawDataValues) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNextPageParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValues(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RawDataValues) contextValidateNextPageParams(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "nextPageParams", "body", string(m.NextPageParams)); err != nil {
		return err
	}

	return nil
}

func (m *RawDataValues) contextValidateTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "time", "body", []int64(m.Time)); err != nil {
		return err
	}

	return nil
}

func (m *RawDataValues) contextValidateValues(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "values", "body", [][]interface{}(m.Values)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RawDataValues) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RawDataValues) UnmarshalBinary(b []byte) error {
	var res RawDataValues
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
