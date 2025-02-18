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

// AzureBillingDiscoveryMethod azure billing discovery method
// swagger:model AzureBillingDiscoveryMethod
type AzureBillingDiscoveryMethod struct {

	// azure billing type
	// Required: true
	AzureBillingType *string `json:"azureBillingType"`
}

// Name gets the name of this subtype
func (m *AzureBillingDiscoveryMethod) Name() string {
	return "ad_azurebilling"
}

// SetName sets the name of this subtype
func (m *AzureBillingDiscoveryMethod) SetName(val string) {
}

// AzureBillingType gets the azure billing type of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *AzureBillingDiscoveryMethod) UnmarshalJSON(raw []byte) error {
	var data struct {

		// azure billing type
		// Required: true
		AzureBillingType *string `json:"azureBillingType"`
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

	var result AzureBillingDiscoveryMethod

	if base.Name != result.Name() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid name value: %q", base.Name)
	}

	result.AzureBillingType = data.AzureBillingType

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m AzureBillingDiscoveryMethod) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// azure billing type
		// Required: true
		AzureBillingType *string `json:"azureBillingType"`
	}{

		AzureBillingType: m.AzureBillingType,
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

// Validate validates this azure billing discovery method
func (m *AzureBillingDiscoveryMethod) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAzureBillingType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureBillingDiscoveryMethod) validateAzureBillingType(formats strfmt.Registry) error {
	if err := validate.Required("azureBillingType", "body", m.AzureBillingType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AzureBillingDiscoveryMethod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureBillingDiscoveryMethod) UnmarshalBinary(b []byte) error {
	var res AzureBillingDiscoveryMethod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
