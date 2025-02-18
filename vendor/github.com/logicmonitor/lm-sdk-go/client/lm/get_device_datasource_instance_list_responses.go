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

// GetDeviceDatasourceInstanceListReader is a Reader for the GetDeviceDatasourceInstanceList structure.
type GetDeviceDatasourceInstanceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceDatasourceInstanceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDeviceDatasourceInstanceListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetDeviceDatasourceInstanceListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDeviceDatasourceInstanceListOK creates a GetDeviceDatasourceInstanceListOK with default headers values
func NewGetDeviceDatasourceInstanceListOK() *GetDeviceDatasourceInstanceListOK {
	return &GetDeviceDatasourceInstanceListOK{}
}

/*GetDeviceDatasourceInstanceListOK handles this case with default header values.

successful operation
*/
type GetDeviceDatasourceInstanceListOK struct {
	Payload *models.DeviceDatasourceInstancePaginationResponse
}

func (o *GetDeviceDatasourceInstanceListOK) Error() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{hdsId}/instances][%d] getDeviceDatasourceInstanceListOK  %+v", 200, o.Payload)
}

func (o *GetDeviceDatasourceInstanceListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.DeviceDatasourceInstancePaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceDatasourceInstanceListDefault creates a GetDeviceDatasourceInstanceListDefault with default headers values
func NewGetDeviceDatasourceInstanceListDefault(code int) *GetDeviceDatasourceInstanceListDefault {
	return &GetDeviceDatasourceInstanceListDefault{
		_statusCode: code,
	}
}

/*GetDeviceDatasourceInstanceListDefault handles this case with default header values.

Error
*/
type GetDeviceDatasourceInstanceListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get device datasource instance list default response
func (o *GetDeviceDatasourceInstanceListDefault) Code() int {
	return o._statusCode
}

func (o *GetDeviceDatasourceInstanceListDefault) Error() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{hdsId}/instances][%d] getDeviceDatasourceInstanceList default  %+v", o._statusCode, o.Payload)
}

func (o *GetDeviceDatasourceInstanceListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
