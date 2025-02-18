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

// GetAPITokenListByAdminIDReader is a Reader for the GetAPITokenListByAdminID structure.
type GetAPITokenListByAdminIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPITokenListByAdminIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAPITokenListByAdminIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetAPITokenListByAdminIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAPITokenListByAdminIDOK creates a GetAPITokenListByAdminIDOK with default headers values
func NewGetAPITokenListByAdminIDOK() *GetAPITokenListByAdminIDOK {
	return &GetAPITokenListByAdminIDOK{}
}

/*GetAPITokenListByAdminIDOK handles this case with default header values.

successful operation
*/
type GetAPITokenListByAdminIDOK struct {
	Payload *models.APITokenPaginationResponse
}

func (o *GetAPITokenListByAdminIDOK) Error() string {
	return fmt.Sprintf("[GET /setting/admins/{adminId}/apitokens][%d] getApiTokenListByAdminIdOK  %+v", 200, o.Payload)
}

func (o *GetAPITokenListByAdminIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.APITokenPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPITokenListByAdminIDDefault creates a GetAPITokenListByAdminIDDefault with default headers values
func NewGetAPITokenListByAdminIDDefault(code int) *GetAPITokenListByAdminIDDefault {
	return &GetAPITokenListByAdminIDDefault{
		_statusCode: code,
	}
}

/*GetAPITokenListByAdminIDDefault handles this case with default header values.

Error
*/
type GetAPITokenListByAdminIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get Api token list by admin Id default response
func (o *GetAPITokenListByAdminIDDefault) Code() int {
	return o._statusCode
}

func (o *GetAPITokenListByAdminIDDefault) Error() string {
	return fmt.Sprintf("[GET /setting/admins/{adminId}/apitokens][%d] getApiTokenListByAdminId default  %+v", o._statusCode, o.Payload)
}

func (o *GetAPITokenListByAdminIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
