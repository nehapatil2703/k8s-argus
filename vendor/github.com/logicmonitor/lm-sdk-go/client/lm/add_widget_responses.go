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

// AddWidgetReader is a Reader for the AddWidget structure.
type AddWidgetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddWidgetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddWidgetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewAddWidgetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddWidgetOK creates a AddWidgetOK with default headers values
func NewAddWidgetOK() *AddWidgetOK {
	return &AddWidgetOK{}
}

/*AddWidgetOK handles this case with default header values.

successful operation
*/
type AddWidgetOK struct {
	Payload models.Widget
}

func (o *AddWidgetOK) Error() string {
	return fmt.Sprintf("[POST /dashboard/widgets][%d] addWidgetOK  %+v", 200, o.Payload)
}

func (o *AddWidgetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload as interface type
	payload, err := models.UnmarshalWidget(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewAddWidgetDefault creates a AddWidgetDefault with default headers values
func NewAddWidgetDefault(code int) *AddWidgetDefault {
	return &AddWidgetDefault{
		_statusCode: code,
	}
}

/*AddWidgetDefault handles this case with default header values.

Error
*/
type AddWidgetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the add widget default response
func (o *AddWidgetDefault) Code() int {
	return o._statusCode
}

func (o *AddWidgetDefault) Error() string {
	return fmt.Sprintf("[POST /dashboard/widgets][%d] addWidget default  %+v", o._statusCode, o.Payload)
}

func (o *AddWidgetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
