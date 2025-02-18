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

// DeleteAdminByIDReader is a Reader for the DeleteAdminByID structure.
type DeleteAdminByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAdminByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteAdminByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteAdminByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAdminByIDOK creates a DeleteAdminByIDOK with default headers values
func NewDeleteAdminByIDOK() *DeleteAdminByIDOK {
	return &DeleteAdminByIDOK{}
}

/*DeleteAdminByIDOK handles this case with default header values.

successful operation
*/
type DeleteAdminByIDOK struct {
	Payload interface{}
}

func (o *DeleteAdminByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /setting/admins/{id}][%d] deleteAdminByIdOK  %+v", 200, o.Payload)
}

func (o *DeleteAdminByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAdminByIDDefault creates a DeleteAdminByIDDefault with default headers values
func NewDeleteAdminByIDDefault(code int) *DeleteAdminByIDDefault {
	return &DeleteAdminByIDDefault{
		_statusCode: code,
	}
}

/*DeleteAdminByIDDefault handles this case with default header values.

Error
*/
type DeleteAdminByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete admin by Id default response
func (o *DeleteAdminByIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAdminByIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /setting/admins/{id}][%d] deleteAdminById default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteAdminByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
