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

// WebsiteNOCItem website n o c item
// swagger:model WebsiteNOCItem
type WebsiteNOCItem struct {

	// group by
	GroupBy string `json:"groupBy,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// website group name
	// Required: true
	WebsiteGroupName *string `json:"websiteGroupName"`

	// website name
	// Required: true
	WebsiteName *string `json:"websiteName"`
}

// Type gets the type of this subtype
func (m *WebsiteNOCItem) Type() string {
	return "website"
}

// SetType sets the type of this subtype
func (m *WebsiteNOCItem) SetType(val string) {
}

// GroupBy gets the group by of this subtype

// Name gets the name of this subtype

// WebsiteGroupName gets the website group name of this subtype

// WebsiteName gets the website name of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *WebsiteNOCItem) UnmarshalJSON(raw []byte) error {
	var data struct {

		// group by
		GroupBy string `json:"groupBy,omitempty"`

		// name
		// Required: true
		Name *string `json:"name"`

		// website group name
		// Required: true
		WebsiteGroupName *string `json:"websiteGroupName"`

		// website name
		// Required: true
		WebsiteName *string `json:"websiteName"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Type string `json:"type"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result WebsiteNOCItem

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.GroupBy = data.GroupBy

	result.Name = data.Name

	result.WebsiteGroupName = data.WebsiteGroupName

	result.WebsiteName = data.WebsiteName

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m WebsiteNOCItem) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// group by
		GroupBy string `json:"groupBy,omitempty"`

		// name
		// Required: true
		Name *string `json:"name"`

		// website group name
		// Required: true
		WebsiteGroupName *string `json:"websiteGroupName"`

		// website name
		// Required: true
		WebsiteName *string `json:"websiteName"`
	}{

		GroupBy: m.GroupBy,

		Name: m.Name,

		WebsiteGroupName: m.WebsiteGroupName,

		WebsiteName: m.WebsiteName,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Type string `json:"type"`
	}{

		Type: m.Type(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this website n o c item
func (m *WebsiteNOCItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWebsiteGroupName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWebsiteName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebsiteNOCItem) validateName(formats strfmt.Registry) error {
	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *WebsiteNOCItem) validateWebsiteGroupName(formats strfmt.Registry) error {
	if err := validate.Required("websiteGroupName", "body", m.WebsiteGroupName); err != nil {
		return err
	}

	return nil
}

func (m *WebsiteNOCItem) validateWebsiteName(formats strfmt.Registry) error {
	if err := validate.Required("websiteName", "body", m.WebsiteName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebsiteNOCItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebsiteNOCItem) UnmarshalBinary(b []byte) error {
	var res WebsiteNOCItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
