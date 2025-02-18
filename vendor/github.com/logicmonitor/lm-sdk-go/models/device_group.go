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

// DeviceGroup device group
// swagger:model DeviceGroup
type DeviceGroup struct {

	// The Applies to custom query for this group (only for dynamic groups)
	AppliesTo string `json:"appliesTo,omitempty"`

	// The number of instances in each AWS region (only applies to AWS groups)
	// Read Only: true
	AwsRegionsInfo string `json:"awsRegionsInfo,omitempty"`

	// The String result returned by the transaction that tests the AWS credentials associated with the AWS group
	// Read Only: true
	AwsTestResult *AwsAccountTestResult `json:"awsTestResult,omitempty"`

	// The Status code result returned by the transaction that tests the AWS credentials associated with the AWS group
	// Read Only: true
	AwsTestResultCode int32 `json:"awsTestResultCode,omitempty"`

	// The number of instances in each Azure region (only applies to Azure groups)
	// Read Only: true
	AzureRegionsInfo string `json:"azureRegionsInfo,omitempty"`

	// The String result returned by the transaction that tests the Azure credentials associated with the Azure group
	// Read Only: true
	AzureTestResult *AzureAccountTestResult `json:"azureTestResult,omitempty"`

	// The Status code result returned by the transaction that tests the Azure credentials associated with the Azure group
	// Read Only: true
	AzureTestResultCode int32 `json:"azureTestResultCode,omitempty"`

	// The time, in epoch seconds format, that the device group was created
	// Read Only: true
	CreatedOn int64 `json:"createdOn,omitempty"`

	// The properties associated with this device group
	CustomProperties []*NameAndValue `json:"customProperties,omitempty"`

	// The description of the default collector assigned to the device group
	// Read Only: true
	DefaultCollectorDescription string `json:"defaultCollectorDescription,omitempty"`

	// The Id of the default collector assigned to the device group
	DefaultCollectorID int32 `json:"defaultCollectorId,omitempty"`

	// The description of the device group
	Description string `json:"description,omitempty"`

	// Indicates whether alerting is disabled (true) or enabled (false) for this device group
	DisableAlerting bool `json:"disableAlerting,omitempty"`

	// Whether or not alerting is effectively disabled for this device group (alerting may be disabled at a higher level, e.g. parent group)
	// Read Only: true
	EffectiveAlertEnabled *bool `json:"effectiveAlertEnabled,omitempty"`

	// Indicates whether Netflow is enabled (true) or disabled (false) for the device group, the default value is true
	EnableNetflow interface{} `json:"enableNetflow,omitempty"`

	// The extra setting for cloud group
	Extra interface{} `json:"extra,omitempty"`

	// The full path of the device group (i.e. if the group 'Dev' is under a parent group named 'Production', the fullPath would be 'Production/Dev'
	// Read Only: true
	FullPath string `json:"fullPath,omitempty"`

	// gcp regions info
	// Read Only: true
	GcpRegionsInfo string `json:"gcpRegionsInfo,omitempty"`

	// The result returned by the transaction that tests the GCP credentials associated with the GCP group
	// Read Only: true
	GcpTestResult *GcpAccountTestResult `json:"gcpTestResult,omitempty"`

	// The Status code result returned by the transaction that tests the GCP credentials associated with the GCP group
	// Read Only: true
	GcpTestResultCode int32 `json:"gcpTestResultCode,omitempty"`

	// normal | dead
	// The status of this device group, where possible statuses are normal and dead. A group with a status of dead may indicate that one or more devices are dead within the group
	// Read Only: true
	GroupStatus string `json:"groupStatus,omitempty"`

	// The type of device group: normal and dynamic device groups will have groupType=Normal, and AWS groups will have a groupType value of AWS/SERVICE (e.g. AWS/S3)
	GroupType string `json:"groupType,omitempty"`

	// Whether if any Netflow enabled devices in this device group
	// Read Only: true
	HasNetflowEnabledDevices *bool `json:"hasNetflowEnabledDevices,omitempty"`

	// The id of the device group
	// Read Only: true
	ID int32 `json:"id,omitempty"`

	// The name of the device group
	// Required: true
	Name *string `json:"name"`

	// The number of AWS devices that belong to this device group (includes AWS devices in sub groups)
	// Read Only: true
	NumOfAWSDevices int64 `json:"numOfAWSDevices,omitempty"`

	// The number of Azure devices that belong to this device group (includes Azure devices in sub groups)
	// Read Only: true
	NumOfAzureDevices int64 `json:"numOfAzureDevices,omitempty"`

	// The number of AWS and normal devices that belong only to this device group (doesn't include devices in sub-groups)
	// Read Only: true
	NumOfDirectDevices int64 `json:"numOfDirectDevices,omitempty"`

	// The number of sub-groups that belong only to this device group (doesn't include groups under sub-groups)
	// Read Only: true
	NumOfDirectSubGroups int64 `json:"numOfDirectSubGroups,omitempty"`

	// num of gcp devices
	// Read Only: true
	NumOfGcpDevices int64 `json:"numOfGcpDevices,omitempty"`

	// The number of total devices, including both AWS and normal devices, that belong to this device group (includes normal devices in sub groups)
	// Read Only: true
	NumOfHosts int64 `json:"numOfHosts,omitempty"`

	// The id of the parent group for this device group (the root device group has an Id of 1)
	ParentID int32 `json:"parentId,omitempty"`

	// The child device groups within this device group
	// Read Only: true
	SubGroups []*DeviceGroupData `json:"subGroups,omitempty"`

	// The permissions for the device group that are granted to the user that made this API request
	// Read Only: true
	UserPermission string `json:"userPermission,omitempty"`
}

// Validate validates this device group
func (m *DeviceGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAwsTestResult(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzureTestResult(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGcpTestResult(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubGroups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceGroup) validateAwsTestResult(formats strfmt.Registry) error {
	if swag.IsZero(m.AwsTestResult) { // not required
		return nil
	}

	if m.AwsTestResult != nil {
		if err := m.AwsTestResult.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsTestResult")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceGroup) validateAzureTestResult(formats strfmt.Registry) error {
	if swag.IsZero(m.AzureTestResult) { // not required
		return nil
	}

	if m.AzureTestResult != nil {
		if err := m.AzureTestResult.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azureTestResult")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceGroup) validateCustomProperties(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomProperties) { // not required
		return nil
	}

	for i := 0; i < len(m.CustomProperties); i++ {
		if swag.IsZero(m.CustomProperties[i]) { // not required
			continue
		}

		if m.CustomProperties[i] != nil {
			if err := m.CustomProperties[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customProperties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceGroup) validateGcpTestResult(formats strfmt.Registry) error {
	if swag.IsZero(m.GcpTestResult) { // not required
		return nil
	}

	if m.GcpTestResult != nil {
		if err := m.GcpTestResult.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcpTestResult")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceGroup) validateName(formats strfmt.Registry) error {
	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DeviceGroup) validateSubGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.SubGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.SubGroups); i++ {
		if swag.IsZero(m.SubGroups[i]) { // not required
			continue
		}

		if m.SubGroups[i] != nil {
			if err := m.SubGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceGroup) UnmarshalBinary(b []byte) error {
	var res DeviceGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
