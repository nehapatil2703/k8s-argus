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

// AddReportGroupReader is a Reader for the AddReportGroup structure.
type AddReportGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddReportGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddReportGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewAddReportGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddReportGroupOK creates a AddReportGroupOK with default headers values
func NewAddReportGroupOK() *AddReportGroupOK {
	return &AddReportGroupOK{}
}

/*AddReportGroupOK handles this case with default header values.

successful operation
*/
type AddReportGroupOK struct {
	Payload *models.ReportGroup
}

func (o *AddReportGroupOK) Error() string {
	return fmt.Sprintf("[POST /report/groups][%d] addReportGroupOK  %+v", 200, o.Payload)
}

func (o *AddReportGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ReportGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddReportGroupDefault creates a AddReportGroupDefault with default headers values
func NewAddReportGroupDefault(code int) *AddReportGroupDefault {
	return &AddReportGroupDefault{
		_statusCode: code,
	}
}

/*AddReportGroupDefault handles this case with default header values.

Error
*/
type AddReportGroupDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the add report group default response
func (o *AddReportGroupDefault) Code() int {
	return o._statusCode
}

func (o *AddReportGroupDefault) Error() string {
	return fmt.Sprintf("[POST /report/groups][%d] addReportGroup default  %+v", o._statusCode, o.Payload)
}

func (o *AddReportGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
