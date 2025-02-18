// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	models "github.com/logicmonitor/lm-sdk-go/models"
)

// PatchDeviceGroupDatasourceAlertSettingReader is a Reader for the PatchDeviceGroupDatasourceAlertSetting structure.
type PatchDeviceGroupDatasourceAlertSettingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchDeviceGroupDatasourceAlertSettingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchDeviceGroupDatasourceAlertSettingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPatchDeviceGroupDatasourceAlertSettingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchDeviceGroupDatasourceAlertSettingOK creates a PatchDeviceGroupDatasourceAlertSettingOK with default headers values
func NewPatchDeviceGroupDatasourceAlertSettingOK() *PatchDeviceGroupDatasourceAlertSettingOK {
	return &PatchDeviceGroupDatasourceAlertSettingOK{}
}

/*PatchDeviceGroupDatasourceAlertSettingOK handles this case with default header values.

successful operation
*/
type PatchDeviceGroupDatasourceAlertSettingOK struct {
	Payload *models.DeviceGroupDataSourceAlertConfig
}

func (o *PatchDeviceGroupDatasourceAlertSettingOK) Error() string {
	return fmt.Sprintf("[PATCH /device/groups/{deviceGroupId}/datasources/{dsId}/alertsettings][%d] patchDeviceGroupDatasourceAlertSettingOK  %+v", 200, o.Payload)
}

func (o *PatchDeviceGroupDatasourceAlertSettingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.DeviceGroupDataSourceAlertConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchDeviceGroupDatasourceAlertSettingDefault creates a PatchDeviceGroupDatasourceAlertSettingDefault with default headers values
func NewPatchDeviceGroupDatasourceAlertSettingDefault(code int) *PatchDeviceGroupDatasourceAlertSettingDefault {
	return &PatchDeviceGroupDatasourceAlertSettingDefault{
		_statusCode: code,
	}
}

/*PatchDeviceGroupDatasourceAlertSettingDefault handles this case with default header values.

Error
*/
type PatchDeviceGroupDatasourceAlertSettingDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch device group datasource alert setting default response
func (o *PatchDeviceGroupDatasourceAlertSettingDefault) Code() int {
	return o._statusCode
}

func (o *PatchDeviceGroupDatasourceAlertSettingDefault) Error() string {
	return fmt.Sprintf("[PATCH /device/groups/{deviceGroupId}/datasources/{dsId}/alertsettings][%d] patchDeviceGroupDatasourceAlertSetting default  %+v", o._statusCode, o.Payload)
}

func (o *PatchDeviceGroupDatasourceAlertSettingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
