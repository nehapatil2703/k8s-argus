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

// CloudWatchAutoDiscoveryMethod cloud watch auto discovery method
// swagger:model CloudWatchAutoDiscoveryMethod
type CloudWatchAutoDiscoveryMethod struct {

	// cluster dimension
	// Required: true
	ClusterDimension *string `json:"clusterDimension"`

	// namespace
	// Required: true
	Namespace *string `json:"namespace"`

	// node dimension
	// Required: true
	NodeDimension *string `json:"nodeDimension"`
}

// Name gets the name of this subtype
func (m *CloudWatchAutoDiscoveryMethod) Name() string {
	return "ad_cloudwatch"
}

// SetName sets the name of this subtype
func (m *CloudWatchAutoDiscoveryMethod) SetName(val string) {
}

// ClusterDimension gets the cluster dimension of this subtype

// Namespace gets the namespace of this subtype

// NodeDimension gets the node dimension of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *CloudWatchAutoDiscoveryMethod) UnmarshalJSON(raw []byte) error {
	var data struct {

		// cluster dimension
		// Required: true
		ClusterDimension *string `json:"clusterDimension"`

		// namespace
		// Required: true
		Namespace *string `json:"namespace"`

		// node dimension
		// Required: true
		NodeDimension *string `json:"nodeDimension"`
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

	var result CloudWatchAutoDiscoveryMethod

	if base.Name != result.Name() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid name value: %q", base.Name)
	}

	result.ClusterDimension = data.ClusterDimension

	result.Namespace = data.Namespace

	result.NodeDimension = data.NodeDimension

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m CloudWatchAutoDiscoveryMethod) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// cluster dimension
		// Required: true
		ClusterDimension *string `json:"clusterDimension"`

		// namespace
		// Required: true
		Namespace *string `json:"namespace"`

		// node dimension
		// Required: true
		NodeDimension *string `json:"nodeDimension"`
	}{

		ClusterDimension: m.ClusterDimension,

		Namespace: m.Namespace,

		NodeDimension: m.NodeDimension,
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

// Validate validates this cloud watch auto discovery method
func (m *CloudWatchAutoDiscoveryMethod) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterDimension(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeDimension(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudWatchAutoDiscoveryMethod) validateClusterDimension(formats strfmt.Registry) error {
	if err := validate.Required("clusterDimension", "body", m.ClusterDimension); err != nil {
		return err
	}

	return nil
}

func (m *CloudWatchAutoDiscoveryMethod) validateNamespace(formats strfmt.Registry) error {
	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

func (m *CloudWatchAutoDiscoveryMethod) validateNodeDimension(formats strfmt.Registry) error {
	if err := validate.Required("nodeDimension", "body", m.NodeDimension); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudWatchAutoDiscoveryMethod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudWatchAutoDiscoveryMethod) UnmarshalBinary(b []byte) error {
	var res CloudWatchAutoDiscoveryMethod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
