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

// AddCollectorGroupReader is a Reader for the AddCollectorGroup structure.
type AddCollectorGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddCollectorGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddCollectorGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewAddCollectorGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddCollectorGroupOK creates a AddCollectorGroupOK with default headers values
func NewAddCollectorGroupOK() *AddCollectorGroupOK {
	return &AddCollectorGroupOK{}
}

/*AddCollectorGroupOK handles this case with default header values.

successful operation
*/
type AddCollectorGroupOK struct {
	Payload *models.CollectorGroup
}

func (o *AddCollectorGroupOK) Error() string {
	return fmt.Sprintf("[POST /setting/collector/groups][%d] addCollectorGroupOK  %+v", 200, o.Payload)
}

func (o *AddCollectorGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.CollectorGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddCollectorGroupDefault creates a AddCollectorGroupDefault with default headers values
func NewAddCollectorGroupDefault(code int) *AddCollectorGroupDefault {
	return &AddCollectorGroupDefault{
		_statusCode: code,
	}
}

/*AddCollectorGroupDefault handles this case with default header values.

Error
*/
type AddCollectorGroupDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the add collector group default response
func (o *AddCollectorGroupDefault) Code() int {
	return o._statusCode
}

func (o *AddCollectorGroupDefault) Error() string {
	return fmt.Sprintf("[POST /setting/collector/groups][%d] addCollectorGroup default  %+v", o._statusCode, o.Payload)
}

func (o *AddCollectorGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
