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

// DeleteDeviceGroupByIDReader is a Reader for the DeleteDeviceGroupByID structure.
type DeleteDeviceGroupByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDeviceGroupByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteDeviceGroupByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteDeviceGroupByIDDefault(response.Code())
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

// NewDeleteDeviceGroupByIDOK creates a DeleteDeviceGroupByIDOK with default headers values
func NewDeleteDeviceGroupByIDOK() *DeleteDeviceGroupByIDOK {
	return &DeleteDeviceGroupByIDOK{}
}

/*DeleteDeviceGroupByIDOK handles this case with default header values.

successful operation
*/
type DeleteDeviceGroupByIDOK struct {
	Payload interface{}
}

func (o *DeleteDeviceGroupByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /device/groups/{id}][%d] deleteDeviceGroupByIdOK  %+v", 200, o.Payload)
}

func (o *DeleteDeviceGroupByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDeviceGroupByIDDefault creates a DeleteDeviceGroupByIDDefault with default headers values
func NewDeleteDeviceGroupByIDDefault(code int) *DeleteDeviceGroupByIDDefault {
	return &DeleteDeviceGroupByIDDefault{
		_statusCode: code,
	}
}

/*DeleteDeviceGroupByIDDefault handles this case with default header values.

Error
*/
type DeleteDeviceGroupByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete device group by Id default response
func (o *DeleteDeviceGroupByIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteDeviceGroupByIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /device/groups/{id}][%d] deleteDeviceGroupById default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteDeviceGroupByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
