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

// UpdateDeviceGroupByIDReader is a Reader for the UpdateDeviceGroupByID structure.
type UpdateDeviceGroupByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDeviceGroupByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateDeviceGroupByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateDeviceGroupByIDDefault(response.Code())
		if result.Code() == 429 {
			errResp := &models.ErrorResponse{
				ErrorCode: 429,
				ErrorDetail: map[string]interface{}{
					"x-rate-limit-limit":     response.GetHeader("x-rate-limit-limit"),
					"x-rate-limit-remaining": response.GetHeader("x-rate-limit-remaining"),
					"x-rate-limit-window":    response.GetHeader("x-rate-limit-window"),
				},
				ErrorMessage: "Customized response from argus sdk",
			}
			result.Payload = errResp
			return nil, result
		}
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateDeviceGroupByIDOK creates a UpdateDeviceGroupByIDOK with default headers values
func NewUpdateDeviceGroupByIDOK() *UpdateDeviceGroupByIDOK {
	return &UpdateDeviceGroupByIDOK{}
}

/*UpdateDeviceGroupByIDOK handles this case with default header values.

successful operation
*/
type UpdateDeviceGroupByIDOK struct {
	Payload *models.DeviceGroup
}

func (o *UpdateDeviceGroupByIDOK) Error() string {
	return fmt.Sprintf("[PUT /device/groups/{id}][%d] updateDeviceGroupByIdOK  %+v", 200, o.Payload)
}

func (o *UpdateDeviceGroupByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.DeviceGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDeviceGroupByIDDefault creates a UpdateDeviceGroupByIDDefault with default headers values
func NewUpdateDeviceGroupByIDDefault(code int) *UpdateDeviceGroupByIDDefault {
	return &UpdateDeviceGroupByIDDefault{
		_statusCode: code,
	}
}

/*UpdateDeviceGroupByIDDefault handles this case with default header values.

Error
*/
type UpdateDeviceGroupByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update device group by Id default response
func (o *UpdateDeviceGroupByIDDefault) Code() int {
	return o._statusCode
}

func (o *UpdateDeviceGroupByIDDefault) Error() string {
	return fmt.Sprintf("[PUT /device/groups/{id}][%d] updateDeviceGroupById default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateDeviceGroupByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
