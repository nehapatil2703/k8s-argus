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

// AckAlertByIDReader is a Reader for the AckAlertByID structure.
type AckAlertByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AckAlertByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAckAlertByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewAckAlertByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAckAlertByIDOK creates a AckAlertByIDOK with default headers values
func NewAckAlertByIDOK() *AckAlertByIDOK {
	return &AckAlertByIDOK{}
}

/*AckAlertByIDOK handles this case with default header values.

successful operation
*/
type AckAlertByIDOK struct {
	Payload interface{}
}

func (o *AckAlertByIDOK) Error() string {
	return fmt.Sprintf("[POST /alert/alerts/{id}/ack][%d] ackAlertByIdOK  %+v", 200, o.Payload)
}

func (o *AckAlertByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAckAlertByIDDefault creates a AckAlertByIDDefault with default headers values
func NewAckAlertByIDDefault(code int) *AckAlertByIDDefault {
	return &AckAlertByIDDefault{
		_statusCode: code,
	}
}

/*AckAlertByIDDefault handles this case with default header values.

Error
*/
type AckAlertByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the ack alert by Id default response
func (o *AckAlertByIDDefault) Code() int {
	return o._statusCode
}

func (o *AckAlertByIDDefault) Error() string {
	return fmt.Sprintf("[POST /alert/alerts/{id}/ack][%d] ackAlertById default  %+v", o._statusCode, o.Payload)
}

func (o *AckAlertByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
