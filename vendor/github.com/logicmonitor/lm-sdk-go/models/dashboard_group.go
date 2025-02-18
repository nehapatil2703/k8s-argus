// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DashboardGroup dashboard group
// swagger:model DashboardGroup
type DashboardGroup struct {

	// The dashboards that belong to the group
	// Read Only: true
	Dashboards []*DashboardData `json:"dashboards,omitempty"`

	// This is a description of the dashboard group
	Description string `json:"description,omitempty"`

	// The full path of the dashboard group
	// Read Only: true
	FullPath string `json:"fullPath,omitempty"`

	// The Id of the dashboard group
	// Read Only: true
	ID int32 `json:"id,omitempty"`

	// The name of the dashboard group
	// Required: true
	Name *string `json:"name"`

	// The number of dashboards that belong to the Dashboard Group and any sub-groups
	// Read Only: true
	NumOfDashboards int64 `json:"numOfDashboards,omitempty"`

	// The number of dashboards that belong directly to the Dashboard Group
	// Read Only: true
	NumOfDirectDashboards int64 `json:"numOfDirectDashboards,omitempty"`

	// The number of groups directly under the Dashboard Group
	// Read Only: true
	NumOfDirectSubGroups int64 `json:"numOfDirectSubGroups,omitempty"`

	// The Id of the parent dashboard group
	ParentID int32 `json:"parentId,omitempty"`

	// The template which is used for import dashboard group
	Template interface{} `json:"template,omitempty"`

	// The permission of the user that made the API call
	// Read Only: true
	UserPermission string `json:"userPermission,omitempty"`

	// The tokens assigned at the group level
	WidgetTokens []*WidgetToken `json:"widgetTokens,omitempty"`
}

// Validate validates this dashboard group
func (m *DashboardGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDashboards(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWidgetTokens(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DashboardGroup) validateDashboards(formats strfmt.Registry) error {
	if swag.IsZero(m.Dashboards) { // not required
		return nil
	}

	for i := 0; i < len(m.Dashboards); i++ {
		if swag.IsZero(m.Dashboards[i]) { // not required
			continue
		}

		if m.Dashboards[i] != nil {
			if err := m.Dashboards[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dashboards" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DashboardGroup) validateName(formats strfmt.Registry) error {
	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DashboardGroup) validateWidgetTokens(formats strfmt.Registry) error {
	if swag.IsZero(m.WidgetTokens) { // not required
		return nil
	}

	for i := 0; i < len(m.WidgetTokens); i++ {
		if swag.IsZero(m.WidgetTokens[i]) { // not required
			continue
		}

		if m.WidgetTokens[i] != nil {
			if err := m.WidgetTokens[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("widgetTokens" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DashboardGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DashboardGroup) UnmarshalBinary(b []byte) error {
	var res DashboardGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
