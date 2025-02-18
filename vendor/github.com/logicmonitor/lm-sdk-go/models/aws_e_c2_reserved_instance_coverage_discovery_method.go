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

// AwsEC2ReservedInstanceCoverageDiscoveryMethod aws e c2 reserved instance coverage discovery method
// swagger:model AwsEC2ReservedInstanceCoverageDiscoveryMethod
type AwsEC2ReservedInstanceCoverageDiscoveryMethod struct {
	AwsEC2ReservedInstanceCoverageDiscoveryMethodAllOf1
}

// Name gets the name of this subtype
func (m *AwsEC2ReservedInstanceCoverageDiscoveryMethod) Name() string {
	return "ad_awsec2reservedinstancecoverage"
}

// SetName sets the name of this subtype
func (m *AwsEC2ReservedInstanceCoverageDiscoveryMethod) SetName(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *AwsEC2ReservedInstanceCoverageDiscoveryMethod) UnmarshalJSON(raw []byte) error {
	var data struct {
		AwsEC2ReservedInstanceCoverageDiscoveryMethodAllOf1
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

	var result AwsEC2ReservedInstanceCoverageDiscoveryMethod

	if base.Name != result.Name() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid name value: %q", base.Name)
	}

	result.AwsEC2ReservedInstanceCoverageDiscoveryMethodAllOf1 = data.AwsEC2ReservedInstanceCoverageDiscoveryMethodAllOf1

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m AwsEC2ReservedInstanceCoverageDiscoveryMethod) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		AwsEC2ReservedInstanceCoverageDiscoveryMethodAllOf1
	}{

		AwsEC2ReservedInstanceCoverageDiscoveryMethodAllOf1: m.AwsEC2ReservedInstanceCoverageDiscoveryMethodAllOf1,
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

// Validate validates this aws e c2 reserved instance coverage discovery method
func (m *AwsEC2ReservedInstanceCoverageDiscoveryMethod) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AwsEC2ReservedInstanceCoverageDiscoveryMethodAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AwsEC2ReservedInstanceCoverageDiscoveryMethod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsEC2ReservedInstanceCoverageDiscoveryMethod) UnmarshalBinary(b []byte) error {
	var res AwsEC2ReservedInstanceCoverageDiscoveryMethod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AwsEC2ReservedInstanceCoverageDiscoveryMethodAllOf1 aws e c2 reserved instance coverage discovery method all of1
// swagger:model AwsEC2ReservedInstanceCoverageDiscoveryMethodAllOf1
type AwsEC2ReservedInstanceCoverageDiscoveryMethodAllOf1 interface{}
