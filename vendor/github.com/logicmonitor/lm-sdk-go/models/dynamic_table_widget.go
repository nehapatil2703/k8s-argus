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
	"github.com/go-openapi/validate"
)

// DynamicTableWidget dynamic table widget
// swagger:model DynamicTableWidget
type DynamicTableWidget struct {
	dashboardIdField *int32

	descriptionField string

	idField int32

	intervalField int32

	lastUpdatedByField string

	lastUpdatedOnField int64

	nameField *string

	themeField string

	timescaleField string

	userPermissionField string

	// columns
	// Required: true
	Columns []*DynamicTableWidgetColumn `json:"columns"`

	// The full name of the selected datasource
	// Read Only: true
	DataSourceFullName string `json:"dataSourceFullName,omitempty"`

	// The id of the selected datasource
	// Required: true
	DataSourceID *int32 `json:"dataSourceId"`

	// forecast
	Forecast *TableWidgetForecastConfiguration `json:"forecast,omitempty"`

	// rows
	// Required: true
	Rows []*DynamicTableWidgetRow `json:"rows"`

	// sort order
	SortOrder string `json:"sortOrder,omitempty"`

	// top x
	TopX int32 `json:"topX,omitempty"`
}

// DashboardID gets the dashboard Id of this subtype
func (m *DynamicTableWidget) DashboardID() *int32 {
	return m.dashboardIdField
}

// SetDashboardID sets the dashboard Id of this subtype
func (m *DynamicTableWidget) SetDashboardID(val *int32) {
	m.dashboardIdField = val
}

// Description gets the description of this subtype
func (m *DynamicTableWidget) Description() string {
	return m.descriptionField
}

// SetDescription sets the description of this subtype
func (m *DynamicTableWidget) SetDescription(val string) {
	m.descriptionField = val
}

// ID gets the id of this subtype
func (m *DynamicTableWidget) ID() int32 {
	return m.idField
}

// SetID sets the id of this subtype
func (m *DynamicTableWidget) SetID(val int32) {
	m.idField = val
}

// Interval gets the interval of this subtype
func (m *DynamicTableWidget) Interval() int32 {
	return m.intervalField
}

// SetInterval sets the interval of this subtype
func (m *DynamicTableWidget) SetInterval(val int32) {
	m.intervalField = val
}

// LastUpdatedBy gets the last updated by of this subtype
func (m *DynamicTableWidget) LastUpdatedBy() string {
	return m.lastUpdatedByField
}

// SetLastUpdatedBy sets the last updated by of this subtype
func (m *DynamicTableWidget) SetLastUpdatedBy(val string) {
	m.lastUpdatedByField = val
}

// LastUpdatedOn gets the last updated on of this subtype
func (m *DynamicTableWidget) LastUpdatedOn() int64 {
	return m.lastUpdatedOnField
}

// SetLastUpdatedOn sets the last updated on of this subtype
func (m *DynamicTableWidget) SetLastUpdatedOn(val int64) {
	m.lastUpdatedOnField = val
}

// Name gets the name of this subtype
func (m *DynamicTableWidget) Name() *string {
	return m.nameField
}

// SetName sets the name of this subtype
func (m *DynamicTableWidget) SetName(val *string) {
	m.nameField = val
}

// Theme gets the theme of this subtype
func (m *DynamicTableWidget) Theme() string {
	return m.themeField
}

// SetTheme sets the theme of this subtype
func (m *DynamicTableWidget) SetTheme(val string) {
	m.themeField = val
}

// Timescale gets the timescale of this subtype
func (m *DynamicTableWidget) Timescale() string {
	return m.timescaleField
}

// SetTimescale sets the timescale of this subtype
func (m *DynamicTableWidget) SetTimescale(val string) {
	m.timescaleField = val
}

// Type gets the type of this subtype
func (m *DynamicTableWidget) Type() string {
	return "dynamicTable"
}

// SetType sets the type of this subtype
func (m *DynamicTableWidget) SetType(val string) {
}

// UserPermission gets the user permission of this subtype
func (m *DynamicTableWidget) UserPermission() string {
	return m.userPermissionField
}

// SetUserPermission sets the user permission of this subtype
func (m *DynamicTableWidget) SetUserPermission(val string) {
	m.userPermissionField = val
}

// Columns gets the columns of this subtype

// DataSourceFullName gets the data source full name of this subtype

// DataSourceID gets the data source Id of this subtype

// Forecast gets the forecast of this subtype

// Rows gets the rows of this subtype

// SortOrder gets the sort order of this subtype

// TopX gets the top x of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *DynamicTableWidget) UnmarshalJSON(raw []byte) error {
	var data struct {

		// columns
		// Required: true
		Columns []*DynamicTableWidgetColumn `json:"columns"`

		// The full name of the selected datasource
		// Read Only: true
		DataSourceFullName string `json:"dataSourceFullName,omitempty"`

		// The id of the selected datasource
		// Required: true
		DataSourceID *int32 `json:"dataSourceId"`

		// forecast
		Forecast *TableWidgetForecastConfiguration `json:"forecast,omitempty"`

		// rows
		// Required: true
		Rows []*DynamicTableWidgetRow `json:"rows"`

		// sort order
		SortOrder string `json:"sortOrder,omitempty"`

		// top x
		TopX int32 `json:"topX,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		DashboardID *int32 `json:"dashboardId"`

		Description string `json:"description,omitempty"`

		ID int32 `json:"id,omitempty"`

		Interval int32 `json:"interval,omitempty"`

		LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`

		LastUpdatedOn int64 `json:"lastUpdatedOn,omitempty"`

		Name *string `json:"name"`

		Theme string `json:"theme,omitempty"`

		Timescale string `json:"timescale,omitempty"`

		Type string `json:"type"`

		UserPermission string `json:"userPermission,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result DynamicTableWidget

	result.dashboardIdField = base.DashboardID

	result.descriptionField = base.Description

	result.idField = base.ID

	result.intervalField = base.Interval

	result.lastUpdatedByField = base.LastUpdatedBy

	result.lastUpdatedOnField = base.LastUpdatedOn

	result.nameField = base.Name

	result.themeField = base.Theme

	result.timescaleField = base.Timescale

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.userPermissionField = base.UserPermission

	result.Columns = data.Columns

	result.DataSourceFullName = data.DataSourceFullName

	result.DataSourceID = data.DataSourceID

	result.Forecast = data.Forecast

	result.Rows = data.Rows

	result.SortOrder = data.SortOrder

	result.TopX = data.TopX

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m DynamicTableWidget) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// columns
		// Required: true
		Columns []*DynamicTableWidgetColumn `json:"columns"`

		// The full name of the selected datasource
		// Read Only: true
		DataSourceFullName string `json:"dataSourceFullName,omitempty"`

		// The id of the selected datasource
		// Required: true
		DataSourceID *int32 `json:"dataSourceId"`

		// forecast
		Forecast *TableWidgetForecastConfiguration `json:"forecast,omitempty"`

		// rows
		// Required: true
		Rows []*DynamicTableWidgetRow `json:"rows"`

		// sort order
		SortOrder string `json:"sortOrder,omitempty"`

		// top x
		TopX int32 `json:"topX,omitempty"`
	}{

		Columns: m.Columns,

		DataSourceFullName: m.DataSourceFullName,

		DataSourceID: m.DataSourceID,

		Forecast: m.Forecast,

		Rows: m.Rows,

		SortOrder: m.SortOrder,

		TopX: m.TopX,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		DashboardID *int32 `json:"dashboardId"`

		Description string `json:"description,omitempty"`

		ID int32 `json:"id,omitempty"`

		Interval int32 `json:"interval,omitempty"`

		LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`

		LastUpdatedOn int64 `json:"lastUpdatedOn,omitempty"`

		Name *string `json:"name"`

		Theme string `json:"theme,omitempty"`

		Timescale string `json:"timescale,omitempty"`

		Type string `json:"type"`

		UserPermission string `json:"userPermission,omitempty"`
	}{

		DashboardID: m.DashboardID(),

		Description: m.Description(),

		ID: m.ID(),

		Interval: m.Interval(),

		LastUpdatedBy: m.LastUpdatedBy(),

		LastUpdatedOn: m.LastUpdatedOn(),

		Name: m.Name(),

		Theme: m.Theme(),

		Timescale: m.Timescale(),

		Type: m.Type(),

		UserPermission: m.UserPermission(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this dynamic table widget
func (m *DynamicTableWidget) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDashboardID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateColumns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataSourceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateForecast(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRows(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DynamicTableWidget) validateDashboardID(formats strfmt.Registry) error {
	if err := validate.Required("dashboardId", "body", m.DashboardID()); err != nil {
		return err
	}

	return nil
}

func (m *DynamicTableWidget) validateName(formats strfmt.Registry) error {
	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

func (m *DynamicTableWidget) validateColumns(formats strfmt.Registry) error {
	if err := validate.Required("columns", "body", m.Columns); err != nil {
		return err
	}

	for i := 0; i < len(m.Columns); i++ {
		if swag.IsZero(m.Columns[i]) { // not required
			continue
		}

		if m.Columns[i] != nil {
			if err := m.Columns[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("columns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DynamicTableWidget) validateDataSourceID(formats strfmt.Registry) error {
	if err := validate.Required("dataSourceId", "body", m.DataSourceID); err != nil {
		return err
	}

	return nil
}

func (m *DynamicTableWidget) validateForecast(formats strfmt.Registry) error {
	if swag.IsZero(m.Forecast) { // not required
		return nil
	}

	if m.Forecast != nil {
		if err := m.Forecast.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("forecast")
			}
			return err
		}
	}

	return nil
}

func (m *DynamicTableWidget) validateRows(formats strfmt.Registry) error {
	if err := validate.Required("rows", "body", m.Rows); err != nil {
		return err
	}

	for i := 0; i < len(m.Rows); i++ {
		if swag.IsZero(m.Rows[i]) { // not required
			continue
		}

		if m.Rows[i] != nil {
			if err := m.Rows[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DynamicTableWidget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DynamicTableWidget) UnmarshalBinary(b []byte) error {
	var res DynamicTableWidget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
