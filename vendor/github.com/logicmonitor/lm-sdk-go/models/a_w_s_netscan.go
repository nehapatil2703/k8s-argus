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

// AWSNetscan a w s netscan
// swagger:model AWSNetscan
type AWSNetscan struct {
	collectorField *int32

	collectorDescriptionField string

	collectorGroupField int32

	collectorGroupNameField string

	creatorField string

	descriptionField string

	duplicateField *ExcludeDuplicateIps

	groupField string

	idField int32

	nameField *string

	nextStartField string

	nextStartEpochField int64

	nsgIdField int32

	scheduleField *NetScanSchedule

	versionField int32

	// access Id
	// Read Only: true
	AccessID string `json:"accessId,omitempty"`

	// account Id
	// Read Only: true
	AccountID string `json:"accountId,omitempty"`

	// assumed role arn
	// Read Only: true
	AssumedRoleArn string `json:"assumedRoleArn,omitempty"`

	// aws a z
	// Read Only: true
	AwsAZ string `json:"awsAZ,omitempty"`

	// aws service
	// Read Only: true
	AwsService string `json:"awsService,omitempty"`

	// external Id
	// Read Only: true
	ExternalID string `json:"externalId,omitempty"`

	// group Id
	// Read Only: true
	GroupID int32 `json:"groupId,omitempty"`

	// root name
	// Read Only: true
	RootName string `json:"rootName,omitempty"`

	// secret key
	// Read Only: true
	SecretKey string `json:"secretKey,omitempty"`
}

// Collector gets the collector of this subtype
func (m *AWSNetscan) Collector() *int32 {
	return m.collectorField
}

// SetCollector sets the collector of this subtype
func (m *AWSNetscan) SetCollector(val *int32) {
	m.collectorField = val
}

// CollectorDescription gets the collector description of this subtype
func (m *AWSNetscan) CollectorDescription() string {
	return m.collectorDescriptionField
}

// SetCollectorDescription sets the collector description of this subtype
func (m *AWSNetscan) SetCollectorDescription(val string) {
	m.collectorDescriptionField = val
}

// CollectorGroup gets the collector group of this subtype
func (m *AWSNetscan) CollectorGroup() int32 {
	return m.collectorGroupField
}

// SetCollectorGroup sets the collector group of this subtype
func (m *AWSNetscan) SetCollectorGroup(val int32) {
	m.collectorGroupField = val
}

// CollectorGroupName gets the collector group name of this subtype
func (m *AWSNetscan) CollectorGroupName() string {
	return m.collectorGroupNameField
}

// SetCollectorGroupName sets the collector group name of this subtype
func (m *AWSNetscan) SetCollectorGroupName(val string) {
	m.collectorGroupNameField = val
}

// Creator gets the creator of this subtype
func (m *AWSNetscan) Creator() string {
	return m.creatorField
}

// SetCreator sets the creator of this subtype
func (m *AWSNetscan) SetCreator(val string) {
	m.creatorField = val
}

// Description gets the description of this subtype
func (m *AWSNetscan) Description() string {
	return m.descriptionField
}

// SetDescription sets the description of this subtype
func (m *AWSNetscan) SetDescription(val string) {
	m.descriptionField = val
}

// Duplicate gets the duplicate of this subtype
func (m *AWSNetscan) Duplicate() *ExcludeDuplicateIps {
	return m.duplicateField
}

// SetDuplicate sets the duplicate of this subtype
func (m *AWSNetscan) SetDuplicate(val *ExcludeDuplicateIps) {
	m.duplicateField = val
}

// Group gets the group of this subtype
func (m *AWSNetscan) Group() string {
	return m.groupField
}

// SetGroup sets the group of this subtype
func (m *AWSNetscan) SetGroup(val string) {
	m.groupField = val
}

// ID gets the id of this subtype
func (m *AWSNetscan) ID() int32 {
	return m.idField
}

// SetID sets the id of this subtype
func (m *AWSNetscan) SetID(val int32) {
	m.idField = val
}

// Method gets the method of this subtype
func (m *AWSNetscan) Method() string {
	return "aws"
}

// SetMethod sets the method of this subtype
func (m *AWSNetscan) SetMethod(val string) {
}

// Name gets the name of this subtype
func (m *AWSNetscan) Name() *string {
	return m.nameField
}

// SetName sets the name of this subtype
func (m *AWSNetscan) SetName(val *string) {
	m.nameField = val
}

// NextStart gets the next start of this subtype
func (m *AWSNetscan) NextStart() string {
	return m.nextStartField
}

// SetNextStart sets the next start of this subtype
func (m *AWSNetscan) SetNextStart(val string) {
	m.nextStartField = val
}

// NextStartEpoch gets the next start epoch of this subtype
func (m *AWSNetscan) NextStartEpoch() int64 {
	return m.nextStartEpochField
}

// SetNextStartEpoch sets the next start epoch of this subtype
func (m *AWSNetscan) SetNextStartEpoch(val int64) {
	m.nextStartEpochField = val
}

// NsgID gets the nsg Id of this subtype
func (m *AWSNetscan) NsgID() int32 {
	return m.nsgIdField
}

// SetNsgID sets the nsg Id of this subtype
func (m *AWSNetscan) SetNsgID(val int32) {
	m.nsgIdField = val
}

// Schedule gets the schedule of this subtype
func (m *AWSNetscan) Schedule() *NetScanSchedule {
	return m.scheduleField
}

// SetSchedule sets the schedule of this subtype
func (m *AWSNetscan) SetSchedule(val *NetScanSchedule) {
	m.scheduleField = val
}

// Version gets the version of this subtype
func (m *AWSNetscan) Version() int32 {
	return m.versionField
}

// SetVersion sets the version of this subtype
func (m *AWSNetscan) SetVersion(val int32) {
	m.versionField = val
}

// AccessID gets the access Id of this subtype

// AccountID gets the account Id of this subtype

// AssumedRoleArn gets the assumed role arn of this subtype

// AwsAZ gets the aws a z of this subtype

// AwsService gets the aws service of this subtype

// ExternalID gets the external Id of this subtype

// GroupID gets the group Id of this subtype

// RootName gets the root name of this subtype

// SecretKey gets the secret key of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *AWSNetscan) UnmarshalJSON(raw []byte) error {
	var data struct {

		// access Id
		// Read Only: true
		AccessID string `json:"accessId,omitempty"`

		// account Id
		// Read Only: true
		AccountID string `json:"accountId,omitempty"`

		// assumed role arn
		// Read Only: true
		AssumedRoleArn string `json:"assumedRoleArn,omitempty"`

		// aws a z
		// Read Only: true
		AwsAZ string `json:"awsAZ,omitempty"`

		// aws service
		// Read Only: true
		AwsService string `json:"awsService,omitempty"`

		// external Id
		// Read Only: true
		ExternalID string `json:"externalId,omitempty"`

		// group Id
		// Read Only: true
		GroupID int32 `json:"groupId,omitempty"`

		// root name
		// Read Only: true
		RootName string `json:"rootName,omitempty"`

		// secret key
		// Read Only: true
		SecretKey string `json:"secretKey,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Collector *int32 `json:"collector"`

		CollectorDescription string `json:"collectorDescription,omitempty"`

		CollectorGroup int32 `json:"collectorGroup,omitempty"`

		CollectorGroupName string `json:"collectorGroupName,omitempty"`

		Creator string `json:"creator,omitempty"`

		Description string `json:"description,omitempty"`

		Duplicate *ExcludeDuplicateIps `json:"duplicate"`

		Group string `json:"group,omitempty"`

		ID int32 `json:"id,omitempty"`

		Method string `json:"method"`

		Name *string `json:"name"`

		NextStart string `json:"nextStart,omitempty"`

		NextStartEpoch int64 `json:"nextStartEpoch,omitempty"`

		NsgID int32 `json:"nsgId,omitempty"`

		Schedule *NetScanSchedule `json:"schedule,omitempty"`

		Version int32 `json:"version,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result AWSNetscan

	result.collectorField = base.Collector

	result.collectorDescriptionField = base.CollectorDescription

	result.collectorGroupField = base.CollectorGroup

	result.collectorGroupNameField = base.CollectorGroupName

	result.creatorField = base.Creator

	result.descriptionField = base.Description

	result.duplicateField = base.Duplicate

	result.groupField = base.Group

	result.idField = base.ID

	if base.Method != result.Method() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid method value: %q", base.Method)
	}

	result.nameField = base.Name

	result.nextStartField = base.NextStart

	result.nextStartEpochField = base.NextStartEpoch

	result.nsgIdField = base.NsgID

	result.scheduleField = base.Schedule

	result.versionField = base.Version

	result.AccessID = data.AccessID

	result.AccountID = data.AccountID

	result.AssumedRoleArn = data.AssumedRoleArn

	result.AwsAZ = data.AwsAZ

	result.AwsService = data.AwsService

	result.ExternalID = data.ExternalID

	result.GroupID = data.GroupID

	result.RootName = data.RootName

	result.SecretKey = data.SecretKey

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m AWSNetscan) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// access Id
		// Read Only: true
		AccessID string `json:"accessId,omitempty"`

		// account Id
		// Read Only: true
		AccountID string `json:"accountId,omitempty"`

		// assumed role arn
		// Read Only: true
		AssumedRoleArn string `json:"assumedRoleArn,omitempty"`

		// aws a z
		// Read Only: true
		AwsAZ string `json:"awsAZ,omitempty"`

		// aws service
		// Read Only: true
		AwsService string `json:"awsService,omitempty"`

		// external Id
		// Read Only: true
		ExternalID string `json:"externalId,omitempty"`

		// group Id
		// Read Only: true
		GroupID int32 `json:"groupId,omitempty"`

		// root name
		// Read Only: true
		RootName string `json:"rootName,omitempty"`

		// secret key
		// Read Only: true
		SecretKey string `json:"secretKey,omitempty"`
	}{

		AccessID: m.AccessID,

		AccountID: m.AccountID,

		AssumedRoleArn: m.AssumedRoleArn,

		AwsAZ: m.AwsAZ,

		AwsService: m.AwsService,

		ExternalID: m.ExternalID,

		GroupID: m.GroupID,

		RootName: m.RootName,

		SecretKey: m.SecretKey,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Collector *int32 `json:"collector"`

		CollectorDescription string `json:"collectorDescription,omitempty"`

		CollectorGroup int32 `json:"collectorGroup,omitempty"`

		CollectorGroupName string `json:"collectorGroupName,omitempty"`

		Creator string `json:"creator,omitempty"`

		Description string `json:"description,omitempty"`

		Duplicate *ExcludeDuplicateIps `json:"duplicate"`

		Group string `json:"group,omitempty"`

		ID int32 `json:"id,omitempty"`

		Method string `json:"method"`

		Name *string `json:"name"`

		NextStart string `json:"nextStart,omitempty"`

		NextStartEpoch int64 `json:"nextStartEpoch,omitempty"`

		NsgID int32 `json:"nsgId,omitempty"`

		Schedule *NetScanSchedule `json:"schedule,omitempty"`

		Version int32 `json:"version,omitempty"`
	}{

		Collector: m.Collector(),

		CollectorDescription: m.CollectorDescription(),

		CollectorGroup: m.CollectorGroup(),

		CollectorGroupName: m.CollectorGroupName(),

		Creator: m.Creator(),

		Description: m.Description(),

		Duplicate: m.Duplicate(),

		Group: m.Group(),

		ID: m.ID(),

		Method: m.Method(),

		Name: m.Name(),

		NextStart: m.NextStart(),

		NextStartEpoch: m.NextStartEpoch(),

		NsgID: m.NsgID(),

		Schedule: m.Schedule(),

		Version: m.Version(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this a w s netscan
func (m *AWSNetscan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCollector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDuplicate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AWSNetscan) validateCollector(formats strfmt.Registry) error {
	if err := validate.Required("collector", "body", m.Collector()); err != nil {
		return err
	}

	return nil
}

func (m *AWSNetscan) validateDuplicate(formats strfmt.Registry) error {
	if err := validate.Required("duplicate", "body", m.Duplicate()); err != nil {
		return err
	}

	if m.Duplicate() != nil {
		if err := m.Duplicate().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("duplicate")
			}
			return err
		}
	}

	return nil
}

func (m *AWSNetscan) validateName(formats strfmt.Registry) error {
	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

func (m *AWSNetscan) validateSchedule(formats strfmt.Registry) error {
	if swag.IsZero(m.Schedule()) { // not required
		return nil
	}

	if m.Schedule() != nil {
		if err := m.Schedule().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AWSNetscan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AWSNetscan) UnmarshalBinary(b []byte) error {
	var res AWSNetscan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
