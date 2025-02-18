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

// DeviceBatchJobSDT device batch job SDT
// swagger:model DeviceBatchJobSDT
type DeviceBatchJobSDT struct {
	adminField string

	commentField string

	durationField int32

	endDateTimeField int64

	endDateTimeOnLocalField string

	endHourField int32

	endMinuteField int32

	hourField int32

	idField string

	isEffectiveField *bool

	minuteField int32

	monthDayField int32

	sdtTypeField string

	startDateTimeField int64

	startDateTimeOnLocalField string

	timezoneField string

	weekDayField string

	weekOfMonthField string

	// The name of the batchjob that the SDT will apply to
	BatchJobName string `json:"batchJobName,omitempty"`

	// The id of the device batchjob that the SDT will be associated with
	DeviceBatchJobID int32 `json:"deviceBatchJobId,omitempty"`

	// The display name of the device associated with the batchjob that the SDT will apply to
	DeviceDisplayName string `json:"deviceDisplayName,omitempty"`

	// The id of the device associated with the batchjob that the SDT will apply to
	DeviceID int32 `json:"deviceId,omitempty"`
}

// Admin gets the admin of this subtype
func (m *DeviceBatchJobSDT) Admin() string {
	return m.adminField
}

// SetAdmin sets the admin of this subtype
func (m *DeviceBatchJobSDT) SetAdmin(val string) {
	m.adminField = val
}

// Comment gets the comment of this subtype
func (m *DeviceBatchJobSDT) Comment() string {
	return m.commentField
}

// SetComment sets the comment of this subtype
func (m *DeviceBatchJobSDT) SetComment(val string) {
	m.commentField = val
}

// Duration gets the duration of this subtype
func (m *DeviceBatchJobSDT) Duration() int32 {
	return m.durationField
}

// SetDuration sets the duration of this subtype
func (m *DeviceBatchJobSDT) SetDuration(val int32) {
	m.durationField = val
}

// EndDateTime gets the end date time of this subtype
func (m *DeviceBatchJobSDT) EndDateTime() int64 {
	return m.endDateTimeField
}

// SetEndDateTime sets the end date time of this subtype
func (m *DeviceBatchJobSDT) SetEndDateTime(val int64) {
	m.endDateTimeField = val
}

// EndDateTimeOnLocal gets the end date time on local of this subtype
func (m *DeviceBatchJobSDT) EndDateTimeOnLocal() string {
	return m.endDateTimeOnLocalField
}

// SetEndDateTimeOnLocal sets the end date time on local of this subtype
func (m *DeviceBatchJobSDT) SetEndDateTimeOnLocal(val string) {
	m.endDateTimeOnLocalField = val
}

// EndHour gets the end hour of this subtype
func (m *DeviceBatchJobSDT) EndHour() int32 {
	return m.endHourField
}

// SetEndHour sets the end hour of this subtype
func (m *DeviceBatchJobSDT) SetEndHour(val int32) {
	m.endHourField = val
}

// EndMinute gets the end minute of this subtype
func (m *DeviceBatchJobSDT) EndMinute() int32 {
	return m.endMinuteField
}

// SetEndMinute sets the end minute of this subtype
func (m *DeviceBatchJobSDT) SetEndMinute(val int32) {
	m.endMinuteField = val
}

// Hour gets the hour of this subtype
func (m *DeviceBatchJobSDT) Hour() int32 {
	return m.hourField
}

// SetHour sets the hour of this subtype
func (m *DeviceBatchJobSDT) SetHour(val int32) {
	m.hourField = val
}

// ID gets the id of this subtype
func (m *DeviceBatchJobSDT) ID() string {
	return m.idField
}

// SetID sets the id of this subtype
func (m *DeviceBatchJobSDT) SetID(val string) {
	m.idField = val
}

// IsEffective gets the is effective of this subtype
func (m *DeviceBatchJobSDT) IsEffective() *bool {
	return m.isEffectiveField
}

// SetIsEffective sets the is effective of this subtype
func (m *DeviceBatchJobSDT) SetIsEffective(val *bool) {
	m.isEffectiveField = val
}

// Minute gets the minute of this subtype
func (m *DeviceBatchJobSDT) Minute() int32 {
	return m.minuteField
}

// SetMinute sets the minute of this subtype
func (m *DeviceBatchJobSDT) SetMinute(val int32) {
	m.minuteField = val
}

// MonthDay gets the month day of this subtype
func (m *DeviceBatchJobSDT) MonthDay() int32 {
	return m.monthDayField
}

// SetMonthDay sets the month day of this subtype
func (m *DeviceBatchJobSDT) SetMonthDay(val int32) {
	m.monthDayField = val
}

// SDTType gets the sdt type of this subtype
func (m *DeviceBatchJobSDT) SDTType() string {
	return m.sdtTypeField
}

// SetSDTType sets the sdt type of this subtype
func (m *DeviceBatchJobSDT) SetSDTType(val string) {
	m.sdtTypeField = val
}

// StartDateTime gets the start date time of this subtype
func (m *DeviceBatchJobSDT) StartDateTime() int64 {
	return m.startDateTimeField
}

// SetStartDateTime sets the start date time of this subtype
func (m *DeviceBatchJobSDT) SetStartDateTime(val int64) {
	m.startDateTimeField = val
}

// StartDateTimeOnLocal gets the start date time on local of this subtype
func (m *DeviceBatchJobSDT) StartDateTimeOnLocal() string {
	return m.startDateTimeOnLocalField
}

// SetStartDateTimeOnLocal sets the start date time on local of this subtype
func (m *DeviceBatchJobSDT) SetStartDateTimeOnLocal(val string) {
	m.startDateTimeOnLocalField = val
}

// Timezone gets the timezone of this subtype
func (m *DeviceBatchJobSDT) Timezone() string {
	return m.timezoneField
}

// SetTimezone sets the timezone of this subtype
func (m *DeviceBatchJobSDT) SetTimezone(val string) {
	m.timezoneField = val
}

// Type gets the type of this subtype
func (m *DeviceBatchJobSDT) Type() string {
	return "DeviceBatchJobSDT"
}

// SetType sets the type of this subtype
func (m *DeviceBatchJobSDT) SetType(val string) {
}

// WeekDay gets the week day of this subtype
func (m *DeviceBatchJobSDT) WeekDay() string {
	return m.weekDayField
}

// SetWeekDay sets the week day of this subtype
func (m *DeviceBatchJobSDT) SetWeekDay(val string) {
	m.weekDayField = val
}

// WeekOfMonth gets the week of month of this subtype
func (m *DeviceBatchJobSDT) WeekOfMonth() string {
	return m.weekOfMonthField
}

// SetWeekOfMonth sets the week of month of this subtype
func (m *DeviceBatchJobSDT) SetWeekOfMonth(val string) {
	m.weekOfMonthField = val
}

// BatchJobName gets the batch job name of this subtype

// DeviceBatchJobID gets the device batch job Id of this subtype

// DeviceDisplayName gets the device display name of this subtype

// DeviceID gets the device Id of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *DeviceBatchJobSDT) UnmarshalJSON(raw []byte) error {
	var data struct {

		// The name of the batchjob that the SDT will apply to
		BatchJobName string `json:"batchJobName,omitempty"`

		// The id of the device batchjob that the SDT will be associated with
		DeviceBatchJobID int32 `json:"deviceBatchJobId,omitempty"`

		// The display name of the device associated with the batchjob that the SDT will apply to
		DeviceDisplayName string `json:"deviceDisplayName,omitempty"`

		// The id of the device associated with the batchjob that the SDT will apply to
		DeviceID int32 `json:"deviceId,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Admin string `json:"admin,omitempty"`

		Comment string `json:"comment,omitempty"`

		Duration int32 `json:"duration,omitempty"`

		EndDateTime int64 `json:"endDateTime,omitempty"`

		EndDateTimeOnLocal string `json:"endDateTimeOnLocal,omitempty"`

		EndHour int32 `json:"endHour,omitempty"`

		EndMinute int32 `json:"endMinute,omitempty"`

		Hour int32 `json:"hour,omitempty"`

		ID string `json:"id,omitempty"`

		IsEffective *bool `json:"isEffective,omitempty"`

		Minute int32 `json:"minute,omitempty"`

		MonthDay int32 `json:"monthDay,omitempty"`

		SDTType string `json:"sdtType,omitempty"`

		StartDateTime int64 `json:"startDateTime,omitempty"`

		StartDateTimeOnLocal string `json:"startDateTimeOnLocal,omitempty"`

		Timezone string `json:"timezone,omitempty"`

		Type string `json:"type"`

		WeekDay string `json:"weekDay,omitempty"`

		WeekOfMonth string `json:"weekOfMonth,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result DeviceBatchJobSDT

	result.adminField = base.Admin

	result.commentField = base.Comment

	result.durationField = base.Duration

	result.endDateTimeField = base.EndDateTime

	result.endDateTimeOnLocalField = base.EndDateTimeOnLocal

	result.endHourField = base.EndHour

	result.endMinuteField = base.EndMinute

	result.hourField = base.Hour

	result.idField = base.ID

	result.isEffectiveField = base.IsEffective

	result.minuteField = base.Minute

	result.monthDayField = base.MonthDay

	result.sdtTypeField = base.SDTType

	result.startDateTimeField = base.StartDateTime

	result.startDateTimeOnLocalField = base.StartDateTimeOnLocal

	result.timezoneField = base.Timezone

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.weekDayField = base.WeekDay

	result.weekOfMonthField = base.WeekOfMonth

	result.BatchJobName = data.BatchJobName

	result.DeviceBatchJobID = data.DeviceBatchJobID

	result.DeviceDisplayName = data.DeviceDisplayName

	result.DeviceID = data.DeviceID

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m DeviceBatchJobSDT) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// The name of the batchjob that the SDT will apply to
		BatchJobName string `json:"batchJobName,omitempty"`

		// The id of the device batchjob that the SDT will be associated with
		DeviceBatchJobID int32 `json:"deviceBatchJobId,omitempty"`

		// The display name of the device associated with the batchjob that the SDT will apply to
		DeviceDisplayName string `json:"deviceDisplayName,omitempty"`

		// The id of the device associated with the batchjob that the SDT will apply to
		DeviceID int32 `json:"deviceId,omitempty"`
	}{

		BatchJobName: m.BatchJobName,

		DeviceBatchJobID: m.DeviceBatchJobID,

		DeviceDisplayName: m.DeviceDisplayName,

		DeviceID: m.DeviceID,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Admin string `json:"admin,omitempty"`

		Comment string `json:"comment,omitempty"`

		Duration int32 `json:"duration,omitempty"`

		EndDateTime int64 `json:"endDateTime,omitempty"`

		EndDateTimeOnLocal string `json:"endDateTimeOnLocal,omitempty"`

		EndHour int32 `json:"endHour,omitempty"`

		EndMinute int32 `json:"endMinute,omitempty"`

		Hour int32 `json:"hour,omitempty"`

		ID string `json:"id,omitempty"`

		IsEffective *bool `json:"isEffective,omitempty"`

		Minute int32 `json:"minute,omitempty"`

		MonthDay int32 `json:"monthDay,omitempty"`

		SDTType string `json:"sdtType,omitempty"`

		StartDateTime int64 `json:"startDateTime,omitempty"`

		StartDateTimeOnLocal string `json:"startDateTimeOnLocal,omitempty"`

		Timezone string `json:"timezone,omitempty"`

		Type string `json:"type"`

		WeekDay string `json:"weekDay,omitempty"`

		WeekOfMonth string `json:"weekOfMonth,omitempty"`
	}{

		Admin: m.Admin(),

		Comment: m.Comment(),

		Duration: m.Duration(),

		EndDateTime: m.EndDateTime(),

		EndDateTimeOnLocal: m.EndDateTimeOnLocal(),

		EndHour: m.EndHour(),

		EndMinute: m.EndMinute(),

		Hour: m.Hour(),

		ID: m.ID(),

		IsEffective: m.IsEffective(),

		Minute: m.Minute(),

		MonthDay: m.MonthDay(),

		SDTType: m.SDTType(),

		StartDateTime: m.StartDateTime(),

		StartDateTimeOnLocal: m.StartDateTimeOnLocal(),

		Timezone: m.Timezone(),

		Type: m.Type(),

		WeekDay: m.WeekDay(),

		WeekOfMonth: m.WeekOfMonth(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this device batch job SDT
func (m *DeviceBatchJobSDT) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DeviceBatchJobSDT) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceBatchJobSDT) UnmarshalBinary(b []byte) error {
	var res DeviceBatchJobSDT
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
