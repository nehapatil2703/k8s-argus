// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InternalCollectorAttribute internal collector attribute
// swagger:model InternalCollectorAttribute
type InternalCollectorAttribute struct {
	InternalCollectorAttributeAllOf1
}

// Name gets the name of this subtype
func (m *InternalCollectorAttribute) Name() string {
	return "internal"
}

// SetName sets the name of this subtype
func (m *InternalCollectorAttribute) SetName(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *InternalCollectorAttribute) UnmarshalJSON(raw []byte) error {
	var data struct {
		InternalCollectorAttributeAllOf1
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

	var result InternalCollectorAttribute

	if base.Name != result.Name() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid name value: %q", base.Name)
	}

	result.InternalCollectorAttributeAllOf1 = data.InternalCollectorAttributeAllOf1

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m InternalCollectorAttribute) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		InternalCollectorAttributeAllOf1
	}{

		InternalCollectorAttributeAllOf1: m.InternalCollectorAttributeAllOf1,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Name string `json:"name"`
	}{

		Name: m.Name(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this internal collector attribute
func (m *InternalCollectorAttribute) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with InternalCollectorAttributeAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *InternalCollectorAttribute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InternalCollectorAttribute) UnmarshalBinary(b []byte) error {
	var res InternalCollectorAttribute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// InternalCollectorAttributeAllOf1 internal collector attribute all of1
// swagger:model InternalCollectorAttributeAllOf1
type InternalCollectorAttributeAllOf1 interface{}
