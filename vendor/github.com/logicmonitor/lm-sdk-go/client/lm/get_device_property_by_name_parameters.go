// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"golang.org/x/net/context"
)

// NewGetDevicePropertyByNameParams creates a new GetDevicePropertyByNameParams object
// with the default values initialized.
func NewGetDevicePropertyByNameParams() *GetDevicePropertyByNameParams {
	var ()
	return &GetDevicePropertyByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDevicePropertyByNameParamsWithTimeout creates a new GetDevicePropertyByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDevicePropertyByNameParamsWithTimeout(timeout time.Duration) *GetDevicePropertyByNameParams {
	var ()
	return &GetDevicePropertyByNameParams{

		timeout: timeout,
	}
}

// NewGetDevicePropertyByNameParamsWithContext creates a new GetDevicePropertyByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDevicePropertyByNameParamsWithContext(ctx context.Context) *GetDevicePropertyByNameParams {
	var ()
	return &GetDevicePropertyByNameParams{

		Context: ctx,
	}
}

// NewGetDevicePropertyByNameParamsWithHTTPClient creates a new GetDevicePropertyByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDevicePropertyByNameParamsWithHTTPClient(client *http.Client) *GetDevicePropertyByNameParams {
	var ()
	return &GetDevicePropertyByNameParams{
		HTTPClient: client,
	}
}

/*GetDevicePropertyByNameParams contains all the parameters to send to the API endpoint
for the get device property by name operation typically these are written to a http.Request
*/
type GetDevicePropertyByNameParams struct {

	/*DeviceID*/
	DeviceID int32
	/*Fields*/
	Fields *string
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get device property by name params
func (o *GetDevicePropertyByNameParams) WithTimeout(timeout time.Duration) *GetDevicePropertyByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device property by name params
func (o *GetDevicePropertyByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device property by name params
func (o *GetDevicePropertyByNameParams) WithContext(ctx context.Context) *GetDevicePropertyByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device property by name params
func (o *GetDevicePropertyByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device property by name params
func (o *GetDevicePropertyByNameParams) WithHTTPClient(client *http.Client) *GetDevicePropertyByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device property by name params
func (o *GetDevicePropertyByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the get device property by name params
func (o *GetDevicePropertyByNameParams) WithDeviceID(deviceID int32) *GetDevicePropertyByNameParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get device property by name params
func (o *GetDevicePropertyByNameParams) SetDeviceID(deviceID int32) {
	o.DeviceID = deviceID
}

// WithFields adds the fields to the get device property by name params
func (o *GetDevicePropertyByNameParams) WithFields(fields *string) *GetDevicePropertyByNameParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get device property by name params
func (o *GetDevicePropertyByNameParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithName adds the name to the get device property by name params
func (o *GetDevicePropertyByNameParams) WithName(name string) *GetDevicePropertyByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get device property by name params
func (o *GetDevicePropertyByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetDevicePropertyByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deviceId
	if err := r.SetPathParam("deviceId", swag.FormatInt32(o.DeviceID)); err != nil {
		return err
	}

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
