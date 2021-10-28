// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/logicmonitor/lm-sdk-go/models"
)

// GetCollectorInstallerReader is a Reader for the GetCollectorInstaller structure.
type GetCollectorInstallerReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *GetCollectorInstallerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCollectorInstallerOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetCollectorInstallerTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetCollectorInstallerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCollectorInstallerOK creates a GetCollectorInstallerOK with default headers values
func NewGetCollectorInstallerOK(writer io.Writer) *GetCollectorInstallerOK {
	return &GetCollectorInstallerOK{

		Payload: writer,
	}
}

/* GetCollectorInstallerOK describes a response with status code 200, with default header values.

successful operation
*/
type GetCollectorInstallerOK struct {
	Payload io.Writer
}

func (o *GetCollectorInstallerOK) Error() string {
	return fmt.Sprintf("[GET /setting/collector/collectors/{collectorId}/installers/{osAndArch}][%d] getCollectorInstallerOK  %+v", 200, o.Payload)
}
func (o *GetCollectorInstallerOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *GetCollectorInstallerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCollectorInstallerTooManyRequests creates a GetCollectorInstallerTooManyRequests with default headers values
func NewGetCollectorInstallerTooManyRequests() *GetCollectorInstallerTooManyRequests {
	return &GetCollectorInstallerTooManyRequests{}
}

/* GetCollectorInstallerTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetCollectorInstallerTooManyRequests struct {

	/* Request limit per X-Rate-Limit-Window
	 */
	XRateLimitLimit int64

	/* The number of requests left for the time window
	 */
	XRateLimitRemaining int64

	/* The rolling time window length with the unit of second
	 */
	XRateLimitWindow int64
}

func (o *GetCollectorInstallerTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /setting/collector/collectors/{collectorId}/installers/{osAndArch}][%d] getCollectorInstallerTooManyRequests ", 429)
}

func (o *GetCollectorInstallerTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-rate-limit-limit
	hdrXRateLimitLimit := response.GetHeader("x-rate-limit-limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("x-rate-limit-limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header x-rate-limit-remaining
	hdrXRateLimitRemaining := response.GetHeader("x-rate-limit-remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("x-rate-limit-remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	// hydrates response header x-rate-limit-window
	hdrXRateLimitWindow := response.GetHeader("x-rate-limit-window")

	if hdrXRateLimitWindow != "" {
		valxRateLimitWindow, err := swag.ConvertInt64(hdrXRateLimitWindow)
		if err != nil {
			return errors.InvalidType("x-rate-limit-window", "header", "int64", hdrXRateLimitWindow)
		}
		o.XRateLimitWindow = valxRateLimitWindow
	}

	return nil
}

// NewGetCollectorInstallerDefault creates a GetCollectorInstallerDefault with default headers values
func NewGetCollectorInstallerDefault(code int) *GetCollectorInstallerDefault {
	return &GetCollectorInstallerDefault{
		_statusCode: code,
	}
}

/* GetCollectorInstallerDefault describes a response with status code -1, with default header values.

Error
*/
type GetCollectorInstallerDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get collector installer default response
func (o *GetCollectorInstallerDefault) Code() int {
	return o._statusCode
}

func (o *GetCollectorInstallerDefault) Error() string {
	return fmt.Sprintf("[GET /setting/collector/collectors/{collectorId}/installers/{osAndArch}][%d] getCollectorInstaller default  %+v", o._statusCode, o.Payload)
}
func (o *GetCollectorInstallerDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetCollectorInstallerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
