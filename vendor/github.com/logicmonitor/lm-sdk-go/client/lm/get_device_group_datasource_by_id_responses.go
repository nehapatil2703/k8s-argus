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

// GetDeviceGroupDatasourceByIDReader is a Reader for the GetDeviceGroupDatasourceByID structure.
type GetDeviceGroupDatasourceByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceGroupDatasourceByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDeviceGroupDatasourceByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetDeviceGroupDatasourceByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDeviceGroupDatasourceByIDOK creates a GetDeviceGroupDatasourceByIDOK with default headers values
func NewGetDeviceGroupDatasourceByIDOK() *GetDeviceGroupDatasourceByIDOK {
	return &GetDeviceGroupDatasourceByIDOK{}
}

/*GetDeviceGroupDatasourceByIDOK handles this case with default header values.

successful operation
*/
type GetDeviceGroupDatasourceByIDOK struct {
	Payload *models.DeviceGroupDataSource
}

func (o *GetDeviceGroupDatasourceByIDOK) Error() string {
	return fmt.Sprintf("[GET /device/groups/{deviceGroupId}/datasources/{id}][%d] getDeviceGroupDatasourceByIdOK  %+v", 200, o.Payload)
}

func (o *GetDeviceGroupDatasourceByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.DeviceGroupDataSource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceGroupDatasourceByIDDefault creates a GetDeviceGroupDatasourceByIDDefault with default headers values
func NewGetDeviceGroupDatasourceByIDDefault(code int) *GetDeviceGroupDatasourceByIDDefault {
	return &GetDeviceGroupDatasourceByIDDefault{
		_statusCode: code,
	}
}

/*GetDeviceGroupDatasourceByIDDefault handles this case with default header values.

Error
*/
type GetDeviceGroupDatasourceByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get device group datasource by Id default response
func (o *GetDeviceGroupDatasourceByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetDeviceGroupDatasourceByIDDefault) Error() string {
	return fmt.Sprintf("[GET /device/groups/{deviceGroupId}/datasources/{id}][%d] getDeviceGroupDatasourceById default  %+v", o._statusCode, o.Payload)
}

func (o *GetDeviceGroupDatasourceByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
