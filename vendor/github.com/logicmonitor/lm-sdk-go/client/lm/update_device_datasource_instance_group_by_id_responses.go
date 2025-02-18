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

// UpdateDeviceDatasourceInstanceGroupByIDReader is a Reader for the UpdateDeviceDatasourceInstanceGroupByID structure.
type UpdateDeviceDatasourceInstanceGroupByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDeviceDatasourceInstanceGroupByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateDeviceDatasourceInstanceGroupByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateDeviceDatasourceInstanceGroupByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateDeviceDatasourceInstanceGroupByIDOK creates a UpdateDeviceDatasourceInstanceGroupByIDOK with default headers values
func NewUpdateDeviceDatasourceInstanceGroupByIDOK() *UpdateDeviceDatasourceInstanceGroupByIDOK {
	return &UpdateDeviceDatasourceInstanceGroupByIDOK{}
}

/*UpdateDeviceDatasourceInstanceGroupByIDOK handles this case with default header values.

successful operation
*/
type UpdateDeviceDatasourceInstanceGroupByIDOK struct {
	Payload *models.DeviceDataSourceInstanceGroup
}

func (o *UpdateDeviceDatasourceInstanceGroupByIDOK) Error() string {
	return fmt.Sprintf("[PUT /device/devices/{deviceId}/devicedatasources/{deviceDsId}/groups/{id}][%d] updateDeviceDatasourceInstanceGroupByIdOK  %+v", 200, o.Payload)
}

func (o *UpdateDeviceDatasourceInstanceGroupByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.DeviceDataSourceInstanceGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDeviceDatasourceInstanceGroupByIDDefault creates a UpdateDeviceDatasourceInstanceGroupByIDDefault with default headers values
func NewUpdateDeviceDatasourceInstanceGroupByIDDefault(code int) *UpdateDeviceDatasourceInstanceGroupByIDDefault {
	return &UpdateDeviceDatasourceInstanceGroupByIDDefault{
		_statusCode: code,
	}
}

/*UpdateDeviceDatasourceInstanceGroupByIDDefault handles this case with default header values.

Error
*/
type UpdateDeviceDatasourceInstanceGroupByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update device datasource instance group by Id default response
func (o *UpdateDeviceDatasourceInstanceGroupByIDDefault) Code() int {
	return o._statusCode
}

func (o *UpdateDeviceDatasourceInstanceGroupByIDDefault) Error() string {
	return fmt.Sprintf("[PUT /device/devices/{deviceId}/devicedatasources/{deviceDsId}/groups/{id}][%d] updateDeviceDatasourceInstanceGroupById default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateDeviceDatasourceInstanceGroupByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
