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
	"github.com/go-openapi/validate"
)

// XENAutoDiscoveryMethod x e n auto discovery method
// swagger:model XENAutoDiscoveryMethod
type XENAutoDiscoveryMethod struct {

	// entity
	// Required: true
	Entity *string `json:"entity"`
}

// Name gets the name of this subtype
func (m *XENAutoDiscoveryMethod) Name() string {
	return "ad_xen"
}

// SetName sets the name of this subtype
func (m *XENAutoDiscoveryMethod) SetName(val string) {
}

// Entity gets the entity of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *XENAutoDiscoveryMethod) UnmarshalJSON(raw []byte) error {
	var data struct {

		// entity
		// Required: true
		Entity *string `json:"entity"`
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

	var result XENAutoDiscoveryMethod

	if base.Name != result.Name() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid name value: %q", base.Name)
	}

	result.Entity = data.Entity

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m XENAutoDiscoveryMethod) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// entity
		// Required: true
		Entity *string `json:"entity"`
	}{

		Entity: m.Entity,
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

// Validate validates this x e n auto discovery method
func (m *XENAutoDiscoveryMethod) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *XENAutoDiscoveryMethod) validateEntity(formats strfmt.Registry) error {
	if err := validate.Required("entity", "body", m.Entity); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *XENAutoDiscoveryMethod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *XENAutoDiscoveryMethod) UnmarshalBinary(b []byte) error {
	var res XENAutoDiscoveryMethod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
