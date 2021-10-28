// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ReportBase report base
//
// swagger:discriminator ReportBase type
type ReportBase interface {
	runtime.Validatable
	runtime.ContextValidatable

	// The id of the custom report template, if the report is a custom report. An id of 0 indicates that the report is not a custom report
	// Read Only: true
	CustomReportTypeID() int32
	SetCustomReportTypeID(int32)

	// The name of the custom report template
	// Read Only: true
	CustomReportTypeName() string
	SetCustomReportTypeName(string)

	// Whether or not the report is configured to be delivered via email. Acceptable values are: none, email
	// Example: email
	Delivery() string
	SetDelivery(string)

	// The description of the report
	// Example: This is a daily alerts report
	Description() string
	SetDescription(string)

	// Whether or not other users are allowed to view the report as the user who last modified the report
	// Read Only: true
	EnableViewAsOtherUser() *bool
	SetEnableViewAsOtherUser(*bool)

	// The format of the report. Allowable values are: HTML, PDF, CSV, WORD
	// Example: HTML
	Format() string
	SetFormat(string)

	// The Id of the group the report is in, where Id=0 is the root report group
	// Example: 5
	GroupID() int32
	SetGroupID(int32)

	// The id of the report
	// Read Only: true
	ID() int32
	SetID(int32)

	// The time, in epoch format, that the report was last generated
	// Read Only: true
	LastGenerateOn() int64
	SetLastGenerateOn(int64)

	// The number of pages in the report, the last time it was generated
	// Read Only: true
	LastGeneratePages() int32
	SetLastGeneratePages(int32)

	// The size of the report, in Bytes, the last time it was generated
	// Read Only: true
	LastGenerateSize() int64
	SetLastGenerateSize(int64)

	// The Id of the user that last modified the report
	// Read Only: true
	LastmodifyUserID() int32
	SetLastmodifyUserID(int32)

	// The username of the user that last modified the report
	// Read Only: true
	LastmodifyUserName() string
	SetLastmodifyUserName(string)

	// The name of the report
	// Example: Daily Alerts Report
	// Required: true
	Name() *string
	SetName(*string)

	// If the report is configured to be delivered via email, this object provides the recipients that the report will be delivered to
	Recipients() []*ReportRecipient
	SetRecipients([]*ReportRecipient)

	// The number of links associated with the report, where each link corresponds to a generated report
	// Read Only: true
	ReportLinkNum() int32
	SetReportLinkNum(int32)

	// A cron schedule that indicates when the report will be delivered via email
	// Example: 0 7 * * 1
	Schedule() string
	SetSchedule(string)

	// The sepecific timezone for the scheduled report
	// Example: America/Los_Angeles
	ScheduleTimezone() string
	SetScheduleTimezone(string)

	// The report type. Acceptable values are: Alert,Alert SLA,Alert threshold,Alert trends,Host CPU,Host group inventory,Host inventory,Host metric trends,Interfaces Bandwidth,Netflow device metric,Service Level Agreement,Website Service Overview,Word template,Audit Log,Alert Forecasting,Dashboard,Website SLA,User,Role
	// Example: Alert
	// Required: true
	Type() string
	SetType(string)

	// The permissions associated with the user who made the API call
	// Read Only: true
	UserPermission() string
	SetUserPermission(string)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type reportBase struct {
	customReportTypeIdField int32

	customReportTypeNameField string

	deliveryField string

	descriptionField string

	enableViewAsOtherUserField *bool

	formatField string

	groupIdField int32

	idField int32

	lastGenerateOnField int64

	lastGeneratePagesField int32

	lastGenerateSizeField int64

	lastmodifyUserIdField int32

	lastmodifyUserNameField string

	nameField *string

	recipientsField []*ReportRecipient

	reportLinkNumField int32

	scheduleField string

	scheduleTimezoneField string

	typeField string

	userPermissionField string
}

// CustomReportTypeID gets the custom report type Id of this polymorphic type
func (m *reportBase) CustomReportTypeID() int32 {
	return m.customReportTypeIdField
}

// SetCustomReportTypeID sets the custom report type Id of this polymorphic type
func (m *reportBase) SetCustomReportTypeID(val int32) {
	m.customReportTypeIdField = val
}

// CustomReportTypeName gets the custom report type name of this polymorphic type
func (m *reportBase) CustomReportTypeName() string {
	return m.customReportTypeNameField
}

// SetCustomReportTypeName sets the custom report type name of this polymorphic type
func (m *reportBase) SetCustomReportTypeName(val string) {
	m.customReportTypeNameField = val
}

// Delivery gets the delivery of this polymorphic type
func (m *reportBase) Delivery() string {
	return m.deliveryField
}

// SetDelivery sets the delivery of this polymorphic type
func (m *reportBase) SetDelivery(val string) {
	m.deliveryField = val
}

// Description gets the description of this polymorphic type
func (m *reportBase) Description() string {
	return m.descriptionField
}

// SetDescription sets the description of this polymorphic type
func (m *reportBase) SetDescription(val string) {
	m.descriptionField = val
}

// EnableViewAsOtherUser gets the enable view as other user of this polymorphic type
func (m *reportBase) EnableViewAsOtherUser() *bool {
	return m.enableViewAsOtherUserField
}

// SetEnableViewAsOtherUser sets the enable view as other user of this polymorphic type
func (m *reportBase) SetEnableViewAsOtherUser(val *bool) {
	m.enableViewAsOtherUserField = val
}

// Format gets the format of this polymorphic type
func (m *reportBase) Format() string {
	return m.formatField
}

// SetFormat sets the format of this polymorphic type
func (m *reportBase) SetFormat(val string) {
	m.formatField = val
}

// GroupID gets the group Id of this polymorphic type
func (m *reportBase) GroupID() int32 {
	return m.groupIdField
}

// SetGroupID sets the group Id of this polymorphic type
func (m *reportBase) SetGroupID(val int32) {
	m.groupIdField = val
}

// ID gets the id of this polymorphic type
func (m *reportBase) ID() int32 {
	return m.idField
}

// SetID sets the id of this polymorphic type
func (m *reportBase) SetID(val int32) {
	m.idField = val
}

// LastGenerateOn gets the last generate on of this polymorphic type
func (m *reportBase) LastGenerateOn() int64 {
	return m.lastGenerateOnField
}

// SetLastGenerateOn sets the last generate on of this polymorphic type
func (m *reportBase) SetLastGenerateOn(val int64) {
	m.lastGenerateOnField = val
}

// LastGeneratePages gets the last generate pages of this polymorphic type
func (m *reportBase) LastGeneratePages() int32 {
	return m.lastGeneratePagesField
}

// SetLastGeneratePages sets the last generate pages of this polymorphic type
func (m *reportBase) SetLastGeneratePages(val int32) {
	m.lastGeneratePagesField = val
}

// LastGenerateSize gets the last generate size of this polymorphic type
func (m *reportBase) LastGenerateSize() int64 {
	return m.lastGenerateSizeField
}

// SetLastGenerateSize sets the last generate size of this polymorphic type
func (m *reportBase) SetLastGenerateSize(val int64) {
	m.lastGenerateSizeField = val
}

// LastmodifyUserID gets the lastmodify user Id of this polymorphic type
func (m *reportBase) LastmodifyUserID() int32 {
	return m.lastmodifyUserIdField
}

// SetLastmodifyUserID sets the lastmodify user Id of this polymorphic type
func (m *reportBase) SetLastmodifyUserID(val int32) {
	m.lastmodifyUserIdField = val
}

// LastmodifyUserName gets the lastmodify user name of this polymorphic type
func (m *reportBase) LastmodifyUserName() string {
	return m.lastmodifyUserNameField
}

// SetLastmodifyUserName sets the lastmodify user name of this polymorphic type
func (m *reportBase) SetLastmodifyUserName(val string) {
	m.lastmodifyUserNameField = val
}

// Name gets the name of this polymorphic type
func (m *reportBase) Name() *string {
	return m.nameField
}

// SetName sets the name of this polymorphic type
func (m *reportBase) SetName(val *string) {
	m.nameField = val
}

// Recipients gets the recipients of this polymorphic type
func (m *reportBase) Recipients() []*ReportRecipient {
	return m.recipientsField
}

// SetRecipients sets the recipients of this polymorphic type
func (m *reportBase) SetRecipients(val []*ReportRecipient) {
	m.recipientsField = val
}

// ReportLinkNum gets the report link num of this polymorphic type
func (m *reportBase) ReportLinkNum() int32 {
	return m.reportLinkNumField
}

// SetReportLinkNum sets the report link num of this polymorphic type
func (m *reportBase) SetReportLinkNum(val int32) {
	m.reportLinkNumField = val
}

// Schedule gets the schedule of this polymorphic type
func (m *reportBase) Schedule() string {
	return m.scheduleField
}

// SetSchedule sets the schedule of this polymorphic type
func (m *reportBase) SetSchedule(val string) {
	m.scheduleField = val
}

// ScheduleTimezone gets the schedule timezone of this polymorphic type
func (m *reportBase) ScheduleTimezone() string {
	return m.scheduleTimezoneField
}

// SetScheduleTimezone sets the schedule timezone of this polymorphic type
func (m *reportBase) SetScheduleTimezone(val string) {
	m.scheduleTimezoneField = val
}

// Type gets the type of this polymorphic type
func (m *reportBase) Type() string {
	return "ReportBase"
}

// SetType sets the type of this polymorphic type
func (m *reportBase) SetType(val string) {
}

// UserPermission gets the user permission of this polymorphic type
func (m *reportBase) UserPermission() string {
	return m.userPermissionField
}

// SetUserPermission sets the user permission of this polymorphic type
func (m *reportBase) SetUserPermission(val string) {
	m.userPermissionField = val
}

// UnmarshalReportBaseSlice unmarshals polymorphic slices of ReportBase
func UnmarshalReportBaseSlice(reader io.Reader, consumer runtime.Consumer) ([]ReportBase, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []ReportBase
	for _, element := range elements {
		obj, err := unmarshalReportBase(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalReportBase unmarshals polymorphic ReportBase
func UnmarshalReportBase(reader io.Reader, consumer runtime.Consumer) (ReportBase, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalReportBase(data, consumer)
}

func unmarshalReportBase(data []byte, consumer runtime.Consumer) (ReportBase, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the type property.
	var getType struct {
		Type string `json:"type"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("type", "body", getType.Type); err != nil {
		return nil, err
	}

	// The value of type is used to determine which type to create and unmarshal the data into
	switch getType.Type {
	case "Alert":
		var result AlertReport
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "Alert Forecasting":
		var result AlertForecastingReport
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "Alert SLA":
		var result AlertSLAReport
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "Alert threshold":
		var result AlertThresholdReport
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "Alert trends":
		var result AlertTrendsReport
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "Audit Log":
		var result AuditLogReport
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "Dashboard":
		var result DashboardReport
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "Host CPU":
		var result HostCPUReport
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "Host group inventory":
		var result HostGroupInventoryReport
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "Host inventory":
		var result HostInventoryReport
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "Host metric trends":
		var result HostMetricsReport
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "Interfaces Bandwidth":
		var result InterfBandwidthReport
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "Netflow device metric":
		var result NetflowReport
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "ReportBase":
		var result reportBase
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "Role":
		var result RoleReport
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "Service Level Agreement":
		var result SLAReport
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "User":
		var result UserReport
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "Website SLA":
		var result WebsiteSLAReport
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "Website Service Overview":
		var result WebsiteOverviewReport
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "Word template":
		var result CustomReport
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid type value: %q", getType.Type)
}

// Validate validates this report base
func (m *reportBase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipients(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *reportBase) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

func (m *reportBase) validateRecipients(formats strfmt.Registry) error {
	if swag.IsZero(m.Recipients()) { // not required
		return nil
	}

	for i := 0; i < len(m.Recipients()); i++ {
		if swag.IsZero(m.recipientsField[i]) { // not required
			continue
		}

		if m.recipientsField[i] != nil {
			if err := m.recipientsField[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recipients" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this report base based on the context it is used
func (m *reportBase) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCustomReportTypeID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCustomReportTypeName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnableViewAsOtherUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastGenerateOn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastGeneratePages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastGenerateSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastmodifyUserID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastmodifyUserName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecipients(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReportLinkNum(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserPermission(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *reportBase) contextValidateCustomReportTypeID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "customReportTypeId", "body", int32(m.CustomReportTypeID())); err != nil {
		return err
	}

	return nil
}

func (m *reportBase) contextValidateCustomReportTypeName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "customReportTypeName", "body", string(m.CustomReportTypeName())); err != nil {
		return err
	}

	return nil
}

func (m *reportBase) contextValidateEnableViewAsOtherUser(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "enableViewAsOtherUser", "body", m.EnableViewAsOtherUser()); err != nil {
		return err
	}

	return nil
}

func (m *reportBase) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int32(m.ID())); err != nil {
		return err
	}

	return nil
}

func (m *reportBase) contextValidateLastGenerateOn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastGenerateOn", "body", int64(m.LastGenerateOn())); err != nil {
		return err
	}

	return nil
}

func (m *reportBase) contextValidateLastGeneratePages(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastGeneratePages", "body", int32(m.LastGeneratePages())); err != nil {
		return err
	}

	return nil
}

func (m *reportBase) contextValidateLastGenerateSize(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastGenerateSize", "body", int64(m.LastGenerateSize())); err != nil {
		return err
	}

	return nil
}

func (m *reportBase) contextValidateLastmodifyUserID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastmodifyUserId", "body", int32(m.LastmodifyUserID())); err != nil {
		return err
	}

	return nil
}

func (m *reportBase) contextValidateLastmodifyUserName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastmodifyUserName", "body", string(m.LastmodifyUserName())); err != nil {
		return err
	}

	return nil
}

func (m *reportBase) contextValidateRecipients(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Recipients()); i++ {

		if m.recipientsField[i] != nil {
			if err := m.recipientsField[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recipients" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *reportBase) contextValidateReportLinkNum(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "reportLinkNum", "body", int32(m.ReportLinkNum())); err != nil {
		return err
	}

	return nil
}

func (m *reportBase) contextValidateUserPermission(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "userPermission", "body", string(m.UserPermission())); err != nil {
		return err
	}

	return nil
}
