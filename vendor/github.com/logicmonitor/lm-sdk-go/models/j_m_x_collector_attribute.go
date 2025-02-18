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

// JMXCollectorAttribute j m x collector attribute
// swagger:model JMXCollectorAttribute
type JMXCollectorAttribute struct {

	// ip
	// Read Only: true
	IP string `json:"ip,omitempty"`

	// query Url
	QueryURL string `json:"queryUrl,omitempty"`
}

// Name gets the name of this subtype
func (m *JMXCollectorAttribute) Name() string {
	return "jmx"
}

// SetName sets the name of this subtype
func (m *JMXCollectorAttribute) SetName(val string) {
}

// IP gets the ip of this subtype

// QueryURL gets the query Url of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *JMXCollectorAttribute) UnmarshalJSON(raw []byte) error {
	var data struct {

		// ip
		// Read Only: true
		IP string `json:"ip,omitempty"`

		// query Url
		QueryURL string `json:"queryUrl,omitempty"`
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

	var result JMXCollectorAttribute

	if base.Name != result.Name() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid name value: %q", base.Name)
	}

	result.IP = data.IP

	result.QueryURL = data.QueryURL

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m JMXCollectorAttribute) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// ip
		// Read Only: true
		IP string `json:"ip,omitempty"`

		// query Url
		QueryURL string `json:"queryUrl,omitempty"`
	}{

		IP: m.IP,

		QueryURL: m.QueryURL,
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

// Validate validates this j m x collector attribute
func (m *JMXCollectorAttribute) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *JMXCollectorAttribute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JMXCollectorAttribute) UnmarshalBinary(b []byte) error {
	var res JMXCollectorAttribute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
