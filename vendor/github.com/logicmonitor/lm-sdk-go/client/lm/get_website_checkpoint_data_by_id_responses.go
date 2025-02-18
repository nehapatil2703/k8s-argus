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

// GetWebsiteCheckpointDataByIDReader is a Reader for the GetWebsiteCheckpointDataByID structure.
type GetWebsiteCheckpointDataByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWebsiteCheckpointDataByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetWebsiteCheckpointDataByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetWebsiteCheckpointDataByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWebsiteCheckpointDataByIDOK creates a GetWebsiteCheckpointDataByIDOK with default headers values
func NewGetWebsiteCheckpointDataByIDOK() *GetWebsiteCheckpointDataByIDOK {
	return &GetWebsiteCheckpointDataByIDOK{}
}

/*GetWebsiteCheckpointDataByIDOK handles this case with default header values.

successful operation
*/
type GetWebsiteCheckpointDataByIDOK struct {
	Payload *models.WebsiteCheckpointRawData
}

func (o *GetWebsiteCheckpointDataByIDOK) Error() string {
	return fmt.Sprintf("[GET /website/websites/{srvId}/checkpoints/{checkId}/data][%d] getWebsiteCheckpointDataByIdOK  %+v", 200, o.Payload)
}

func (o *GetWebsiteCheckpointDataByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.WebsiteCheckpointRawData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebsiteCheckpointDataByIDDefault creates a GetWebsiteCheckpointDataByIDDefault with default headers values
func NewGetWebsiteCheckpointDataByIDDefault(code int) *GetWebsiteCheckpointDataByIDDefault {
	return &GetWebsiteCheckpointDataByIDDefault{
		_statusCode: code,
	}
}

/*GetWebsiteCheckpointDataByIDDefault handles this case with default header values.

Error
*/
type GetWebsiteCheckpointDataByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get website checkpoint data by Id default response
func (o *GetWebsiteCheckpointDataByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetWebsiteCheckpointDataByIDDefault) Error() string {
	return fmt.Sprintf("[GET /website/websites/{srvId}/checkpoints/{checkId}/data][%d] getWebsiteCheckpointDataById default  %+v", o._statusCode, o.Payload)
}

func (o *GetWebsiteCheckpointDataByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
