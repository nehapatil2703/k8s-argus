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

// GetDeviceDatasourceInstanceGroupOverviewGraphDataReader is a Reader for the GetDeviceDatasourceInstanceGroupOverviewGraphData structure.
type GetDeviceDatasourceInstanceGroupOverviewGraphDataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceDatasourceInstanceGroupOverviewGraphDataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDeviceDatasourceInstanceGroupOverviewGraphDataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetDeviceDatasourceInstanceGroupOverviewGraphDataDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDeviceDatasourceInstanceGroupOverviewGraphDataOK creates a GetDeviceDatasourceInstanceGroupOverviewGraphDataOK with default headers values
func NewGetDeviceDatasourceInstanceGroupOverviewGraphDataOK() *GetDeviceDatasourceInstanceGroupOverviewGraphDataOK {
	return &GetDeviceDatasourceInstanceGroupOverviewGraphDataOK{}
}

/*GetDeviceDatasourceInstanceGroupOverviewGraphDataOK handles this case with default header values.

successful operation
*/
type GetDeviceDatasourceInstanceGroupOverviewGraphDataOK struct {
	Payload *models.GraphPlot
}

func (o *GetDeviceDatasourceInstanceGroupOverviewGraphDataOK) Error() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{deviceDsId}/groups/{dsigId}/graphs/{ographId}/data][%d] getDeviceDatasourceInstanceGroupOverviewGraphDataOK  %+v", 200, o.Payload)
}

func (o *GetDeviceDatasourceInstanceGroupOverviewGraphDataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.GraphPlot)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceDatasourceInstanceGroupOverviewGraphDataDefault creates a GetDeviceDatasourceInstanceGroupOverviewGraphDataDefault with default headers values
func NewGetDeviceDatasourceInstanceGroupOverviewGraphDataDefault(code int) *GetDeviceDatasourceInstanceGroupOverviewGraphDataDefault {
	return &GetDeviceDatasourceInstanceGroupOverviewGraphDataDefault{
		_statusCode: code,
	}
}

/*GetDeviceDatasourceInstanceGroupOverviewGraphDataDefault handles this case with default header values.

Error
*/
type GetDeviceDatasourceInstanceGroupOverviewGraphDataDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get device datasource instance group overview graph data default response
func (o *GetDeviceDatasourceInstanceGroupOverviewGraphDataDefault) Code() int {
	return o._statusCode
}

func (o *GetDeviceDatasourceInstanceGroupOverviewGraphDataDefault) Error() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{deviceDsId}/groups/{dsigId}/graphs/{ographId}/data][%d] getDeviceDatasourceInstanceGroupOverviewGraphData default  %+v", o._statusCode, o.Payload)
}

func (o *GetDeviceDatasourceInstanceGroupOverviewGraphDataDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
