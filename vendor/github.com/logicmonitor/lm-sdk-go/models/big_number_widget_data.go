// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BigNumberWidgetData big number widget data
// swagger:model BigNumberWidgetData
type BigNumberWidgetData struct {
	titleField string

	// data
	// Read Only: true
	Data []*BigNumberData `json:"data,omitempty"`
}

// Title gets the title of this subtype
func (m *BigNumberWidgetData) Title() string {
	return m.titleField
}

// SetTitle sets the title of this subtype
func (m *BigNumberWidgetData) SetTitle(val string) {
	m.titleField = val
}

// Type gets the type of this subtype
func (m *BigNumberWidgetData) Type() string {
	return "bigNumber"
}

// SetType sets the type of this subtype
func (m *BigNumberWidgetData) SetType(val string) {
}

// Data gets the data of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *BigNumberWidgetData) UnmarshalJSON(raw []byte) error {
	var data struct {

		// data
		// Read Only: true
		Data []*BigNumberData `json:"data,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Title string `json:"title,omitempty"`

		Type string `json:"type,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result BigNumberWidgetData

	result.titleField = base.Title

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.Data = data.Data

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m BigNumberWidgetData) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// data
		// Read Only: true
		Data []*BigNumberData `json:"data,omitempty"`
	}{

		Data: m.Data,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Title string `json:"title,omitempty"`

		Type string `json:"type,omitempty"`
	}{

		Title: m.Title(),

		Type: m.Type(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this big number widget data
func (m *BigNumberWidgetData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BigNumberWidgetData) validateData(formats strfmt.Registry) error {
	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BigNumberWidgetData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BigNumberWidgetData) UnmarshalBinary(b []byte) error {
	var res BigNumberWidgetData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
