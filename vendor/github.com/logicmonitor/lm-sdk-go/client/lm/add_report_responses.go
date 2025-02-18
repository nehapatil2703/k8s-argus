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

// AddReportReader is a Reader for the AddReport structure.
type AddReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddReportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewAddReportDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddReportOK creates a AddReportOK with default headers values
func NewAddReportOK() *AddReportOK {
	return &AddReportOK{}
}

/*AddReportOK handles this case with default header values.

successful operation
*/
type AddReportOK struct {
	Payload models.ReportBase
}

func (o *AddReportOK) Error() string {
	return fmt.Sprintf("[POST /report/reports][%d] addReportOK  %+v", 200, o.Payload)
}

func (o *AddReportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload as interface type
	payload, err := models.UnmarshalReportBase(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewAddReportDefault creates a AddReportDefault with default headers values
func NewAddReportDefault(code int) *AddReportDefault {
	return &AddReportDefault{
		_statusCode: code,
	}
}

/*AddReportDefault handles this case with default header values.

Error
*/
type AddReportDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the add report default response
func (o *AddReportDefault) Code() int {
	return o._statusCode
}

func (o *AddReportDefault) Error() string {
	return fmt.Sprintf("[POST /report/reports][%d] addReport default  %+v", o._statusCode, o.Payload)
}

func (o *AddReportDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
