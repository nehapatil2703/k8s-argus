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

// AwsServiceLimitsFromTrustedAdvisorCollectorAttribute aws service limits from trusted advisor collector attribute
// swagger:model AwsServiceLimitsFromTrustedAdvisorCollectorAttribute
type AwsServiceLimitsFromTrustedAdvisorCollectorAttribute struct {

	// aws service name
	AwsServiceName string `json:"awsServiceName,omitempty"`

	// period
	Period int32 `json:"period,omitempty"`
}

// Name gets the name of this subtype
func (m *AwsServiceLimitsFromTrustedAdvisorCollectorAttribute) Name() string {
	return "awsservicelimitsfromtrustedadvisor"
}

// SetName sets the name of this subtype
func (m *AwsServiceLimitsFromTrustedAdvisorCollectorAttribute) SetName(val string) {
}

// AwsServiceName gets the aws service name of this subtype

// Period gets the period of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *AwsServiceLimitsFromTrustedAdvisorCollectorAttribute) UnmarshalJSON(raw []byte) error {
	var data struct {

		// aws service name
		AwsServiceName string `json:"awsServiceName,omitempty"`

		// period
		Period int32 `json:"period,omitempty"`
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

	var result AwsServiceLimitsFromTrustedAdvisorCollectorAttribute

	if base.Name != result.Name() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid name value: %q", base.Name)
	}

	result.AwsServiceName = data.AwsServiceName

	result.Period = data.Period

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m AwsServiceLimitsFromTrustedAdvisorCollectorAttribute) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// aws service name
		AwsServiceName string `json:"awsServiceName,omitempty"`

		// period
		Period int32 `json:"period,omitempty"`
	}{

		AwsServiceName: m.AwsServiceName,

		Period: m.Period,
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

// Validate validates this aws service limits from trusted advisor collector attribute
func (m *AwsServiceLimitsFromTrustedAdvisorCollectorAttribute) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AwsServiceLimitsFromTrustedAdvisorCollectorAttribute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsServiceLimitsFromTrustedAdvisorCollectorAttribute) UnmarshalBinary(b []byte) error {
	var res AwsServiceLimitsFromTrustedAdvisorCollectorAttribute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
