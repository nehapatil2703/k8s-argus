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

// GetSDTHistoryByDeviceIDReader is a Reader for the GetSDTHistoryByDeviceID structure.
type GetSDTHistoryByDeviceIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSDTHistoryByDeviceIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSDTHistoryByDeviceIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetSDTHistoryByDeviceIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSDTHistoryByDeviceIDOK creates a GetSDTHistoryByDeviceIDOK with default headers values
func NewGetSDTHistoryByDeviceIDOK() *GetSDTHistoryByDeviceIDOK {
	return &GetSDTHistoryByDeviceIDOK{}
}

/*GetSDTHistoryByDeviceIDOK handles this case with default header values.

successful operation
*/
type GetSDTHistoryByDeviceIDOK struct {
	Payload *models.DeviceSDTHistoryPaginationResponse
}

func (o *GetSDTHistoryByDeviceIDOK) Error() string {
	return fmt.Sprintf("[GET /device/devices/{id}/historysdts][%d] getSdtHistoryByDeviceIdOK  %+v", 200, o.Payload)
}

func (o *GetSDTHistoryByDeviceIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.DeviceSDTHistoryPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSDTHistoryByDeviceIDDefault creates a GetSDTHistoryByDeviceIDDefault with default headers values
func NewGetSDTHistoryByDeviceIDDefault(code int) *GetSDTHistoryByDeviceIDDefault {
	return &GetSDTHistoryByDeviceIDDefault{
		_statusCode: code,
	}
}

/*GetSDTHistoryByDeviceIDDefault handles this case with default header values.

Error
*/
type GetSDTHistoryByDeviceIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get SDT history by device Id default response
func (o *GetSDTHistoryByDeviceIDDefault) Code() int {
	return o._statusCode
}

func (o *GetSDTHistoryByDeviceIDDefault) Error() string {
	return fmt.Sprintf("[GET /device/devices/{id}/historysdts][%d] getSDTHistoryByDeviceId default  %+v", o._statusCode, o.Payload)
}

func (o *GetSDTHistoryByDeviceIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
