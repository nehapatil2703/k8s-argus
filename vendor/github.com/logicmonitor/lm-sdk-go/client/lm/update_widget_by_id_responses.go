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

// UpdateWidgetByIDReader is a Reader for the UpdateWidgetByID structure.
type UpdateWidgetByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateWidgetByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateWidgetByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateWidgetByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateWidgetByIDOK creates a UpdateWidgetByIDOK with default headers values
func NewUpdateWidgetByIDOK() *UpdateWidgetByIDOK {
	return &UpdateWidgetByIDOK{}
}

/*UpdateWidgetByIDOK handles this case with default header values.

successful operation
*/
type UpdateWidgetByIDOK struct {
	Payload models.Widget
}

func (o *UpdateWidgetByIDOK) Error() string {
	return fmt.Sprintf("[PUT /dashboard/widgets/{id}][%d] updateWidgetByIdOK  %+v", 200, o.Payload)
}

func (o *UpdateWidgetByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload as interface type
	payload, err := models.UnmarshalWidget(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewUpdateWidgetByIDDefault creates a UpdateWidgetByIDDefault with default headers values
func NewUpdateWidgetByIDDefault(code int) *UpdateWidgetByIDDefault {
	return &UpdateWidgetByIDDefault{
		_statusCode: code,
	}
}

/*UpdateWidgetByIDDefault handles this case with default header values.

Error
*/
type UpdateWidgetByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update widget by Id default response
func (o *UpdateWidgetByIDDefault) Code() int {
	return o._statusCode
}

func (o *UpdateWidgetByIDDefault) Error() string {
	return fmt.Sprintf("[PUT /dashboard/widgets/{id}][%d] updateWidgetById default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateWidgetByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
