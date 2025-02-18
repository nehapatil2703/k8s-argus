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

// GoogleMapWidgetData google map widget data
// swagger:model GoogleMapWidgetData
type GoogleMapWidgetData struct {
	titleField string

	// items
	// Read Only: true
	Items []*MapItemInfo `json:"items,omitempty"`
}

// Title gets the title of this subtype
func (m *GoogleMapWidgetData) Title() string {
	return m.titleField
}

// SetTitle sets the title of this subtype
func (m *GoogleMapWidgetData) SetTitle(val string) {
	m.titleField = val
}

// Type gets the type of this subtype
func (m *GoogleMapWidgetData) Type() string {
	return "gmap"
}

// SetType sets the type of this subtype
func (m *GoogleMapWidgetData) SetType(val string) {
}

// Items gets the items of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *GoogleMapWidgetData) UnmarshalJSON(raw []byte) error {
	var data struct {

		// items
		// Read Only: true
		Items []*MapItemInfo `json:"items,omitempty"`
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

	var result GoogleMapWidgetData

	result.titleField = base.Title

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.Items = data.Items

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m GoogleMapWidgetData) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// items
		// Read Only: true
		Items []*MapItemInfo `json:"items,omitempty"`
	}{

		Items: m.Items,
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

// Validate validates this google map widget data
func (m *GoogleMapWidgetData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GoogleMapWidgetData) validateItems(formats strfmt.Registry) error {
	if swag.IsZero(m.Items) { // not required
		return nil
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GoogleMapWidgetData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GoogleMapWidgetData) UnmarshalBinary(b []byte) error {
	var res GoogleMapWidgetData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
