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

// DeleteAlertRuleByIDReader is a Reader for the DeleteAlertRuleByID structure.
type DeleteAlertRuleByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAlertRuleByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteAlertRuleByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteAlertRuleByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAlertRuleByIDOK creates a DeleteAlertRuleByIDOK with default headers values
func NewDeleteAlertRuleByIDOK() *DeleteAlertRuleByIDOK {
	return &DeleteAlertRuleByIDOK{}
}

/*DeleteAlertRuleByIDOK handles this case with default header values.

successful operation
*/
type DeleteAlertRuleByIDOK struct {
	Payload interface{}
}

func (o *DeleteAlertRuleByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /setting/alert/rules/{id}][%d] deleteAlertRuleByIdOK  %+v", 200, o.Payload)
}

func (o *DeleteAlertRuleByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertRuleByIDDefault creates a DeleteAlertRuleByIDDefault with default headers values
func NewDeleteAlertRuleByIDDefault(code int) *DeleteAlertRuleByIDDefault {
	return &DeleteAlertRuleByIDDefault{
		_statusCode: code,
	}
}

/*DeleteAlertRuleByIDDefault handles this case with default header values.

Error
*/
type DeleteAlertRuleByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete alert rule by Id default response
func (o *DeleteAlertRuleByIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAlertRuleByIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /setting/alert/rules/{id}][%d] deleteAlertRuleById default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteAlertRuleByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
