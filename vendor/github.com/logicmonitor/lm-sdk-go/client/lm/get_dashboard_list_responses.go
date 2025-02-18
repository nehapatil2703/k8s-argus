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

// GetDashboardListReader is a Reader for the GetDashboardList structure.
type GetDashboardListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDashboardListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDashboardListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetDashboardListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDashboardListOK creates a GetDashboardListOK with default headers values
func NewGetDashboardListOK() *GetDashboardListOK {
	return &GetDashboardListOK{}
}

/*GetDashboardListOK handles this case with default header values.

successful operation
*/
type GetDashboardListOK struct {
	Payload *models.DashboardPaginationResponse
}

func (o *GetDashboardListOK) Error() string {
	return fmt.Sprintf("[GET /dashboard/dashboards][%d] getDashboardListOK  %+v", 200, o.Payload)
}

func (o *GetDashboardListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.DashboardPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDashboardListDefault creates a GetDashboardListDefault with default headers values
func NewGetDashboardListDefault(code int) *GetDashboardListDefault {
	return &GetDashboardListDefault{
		_statusCode: code,
	}
}

/*GetDashboardListDefault handles this case with default header values.

Error
*/
type GetDashboardListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get dashboard list default response
func (o *GetDashboardListDefault) Code() int {
	return o._statusCode
}

func (o *GetDashboardListDefault) Error() string {
	return fmt.Sprintf("[GET /dashboard/dashboards][%d] getDashboardList default  %+v", o._statusCode, o.Payload)
}

func (o *GetDashboardListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
