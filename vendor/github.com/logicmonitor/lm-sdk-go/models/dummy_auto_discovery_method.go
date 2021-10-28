// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DummyAutoDiscoveryMethod dummy auto discovery method
//
// swagger:model DummyAutoDiscoveryMethod
type DummyAutoDiscoveryMethod struct {

	// generate count
	// Required: true
	GenerateCount *int32 `json:"generateCount"`

	// get generate count2
	// Required: true
	GetGenerateCount2 *int32 `json:"getGenerateCount2"`

	// max number
	// Required: true
	MaxNumber *int32 `json:"maxNumber"`
}

// Name gets the name of this subtype
func (m *DummyAutoDiscoveryMethod) Name() string {
	return "ad_dummy"
}

// SetName sets the name of this subtype
func (m *DummyAutoDiscoveryMethod) SetName(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *DummyAutoDiscoveryMethod) UnmarshalJSON(raw []byte) error {
	var data struct {

		// generate count
		// Required: true
		GenerateCount *int32 `json:"generateCount"`

		// get generate count2
		// Required: true
		GetGenerateCount2 *int32 `json:"getGenerateCount2"`

		// max number
		// Required: true
		MaxNumber *int32 `json:"maxNumber"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Name string `json:"name"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result DummyAutoDiscoveryMethod

	if base.Name != result.Name() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid name value: %q", base.Name)
	}

	result.GenerateCount = data.GenerateCount
	result.GetGenerateCount2 = data.GetGenerateCount2
	result.MaxNumber = data.MaxNumber

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m DummyAutoDiscoveryMethod) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// generate count
		// Required: true
		GenerateCount *int32 `json:"generateCount"`

		// get generate count2
		// Required: true
		GetGenerateCount2 *int32 `json:"getGenerateCount2"`

		// max number
		// Required: true
		MaxNumber *int32 `json:"maxNumber"`
	}{

		GenerateCount: m.GenerateCount,

		GetGenerateCount2: m.GetGenerateCount2,

		MaxNumber: m.MaxNumber,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Name string `json:"name"`
	}{

		Name: m.Name(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this dummy auto discovery method
func (m *DummyAutoDiscoveryMethod) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGenerateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGetGenerateCount2(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DummyAutoDiscoveryMethod) validateGenerateCount(formats strfmt.Registry) error {

	if err := validate.Required("generateCount", "body", m.GenerateCount); err != nil {
		return err
	}

	return nil
}

func (m *DummyAutoDiscoveryMethod) validateGetGenerateCount2(formats strfmt.Registry) error {

	if err := validate.Required("getGenerateCount2", "body", m.GetGenerateCount2); err != nil {
		return err
	}

	return nil
}

func (m *DummyAutoDiscoveryMethod) validateMaxNumber(formats strfmt.Registry) error {

	if err := validate.Required("maxNumber", "body", m.MaxNumber); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this dummy auto discovery method based on the context it is used
func (m *DummyAutoDiscoveryMethod) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DummyAutoDiscoveryMethod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DummyAutoDiscoveryMethod) UnmarshalBinary(b []byte) error {
	var res DummyAutoDiscoveryMethod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
