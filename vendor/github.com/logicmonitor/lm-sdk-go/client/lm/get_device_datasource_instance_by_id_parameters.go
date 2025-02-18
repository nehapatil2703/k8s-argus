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

// NewGetDeviceDatasourceInstanceByIDParams creates a new GetDeviceDatasourceInstanceByIDParams object
// with the default values initialized.
func NewGetDeviceDatasourceInstanceByIDParams() *GetDeviceDatasourceInstanceByIDParams {
	var ()
	return &GetDeviceDatasourceInstanceByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceDatasourceInstanceByIDParamsWithTimeout creates a new GetDeviceDatasourceInstanceByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeviceDatasourceInstanceByIDParamsWithTimeout(timeout time.Duration) *GetDeviceDatasourceInstanceByIDParams {
	var ()
	return &GetDeviceDatasourceInstanceByIDParams{

		timeout: timeout,
	}
}

// NewGetDeviceDatasourceInstanceByIDParamsWithContext creates a new GetDeviceDatasourceInstanceByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeviceDatasourceInstanceByIDParamsWithContext(ctx context.Context) *GetDeviceDatasourceInstanceByIDParams {
	var ()
	return &GetDeviceDatasourceInstanceByIDParams{

		Context: ctx,
	}
}

// NewGetDeviceDatasourceInstanceByIDParamsWithHTTPClient creates a new GetDeviceDatasourceInstanceByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeviceDatasourceInstanceByIDParamsWithHTTPClient(client *http.Client) *GetDeviceDatasourceInstanceByIDParams {
	var ()
	return &GetDeviceDatasourceInstanceByIDParams{
		HTTPClient: client,
	}
}

/*GetDeviceDatasourceInstanceByIDParams contains all the parameters to send to the API endpoint
for the get device datasource instance by Id operation typically these are written to a http.Request
*/
type GetDeviceDatasourceInstanceByIDParams struct {

	/*DeviceID*/
	DeviceID int32
	/*Fields*/
	Fields *string
	/*HdsID
	  The device-datasource ID

	*/
	HdsID int32
	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get device datasource instance by Id params
func (o *GetDeviceDatasourceInstanceByIDParams) WithTimeout(timeout time.Duration) *GetDeviceDatasourceInstanceByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device datasource instance by Id params
func (o *GetDeviceDatasourceInstanceByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device datasource instance by Id params
func (o *GetDeviceDatasourceInstanceByIDParams) WithContext(ctx context.Context) *GetDeviceDatasourceInstanceByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device datasource instance by Id params
func (o *GetDeviceDatasourceInstanceByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device datasource instance by Id params
func (o *GetDeviceDatasourceInstanceByIDParams) WithHTTPClient(client *http.Client) *GetDeviceDatasourceInstanceByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device datasource instance by Id params
func (o *GetDeviceDatasourceInstanceByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the get device datasource instance by Id params
func (o *GetDeviceDatasourceInstanceByIDParams) WithDeviceID(deviceID int32) *GetDeviceDatasourceInstanceByIDParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get device datasource instance by Id params
func (o *GetDeviceDatasourceInstanceByIDParams) SetDeviceID(deviceID int32) {
	o.DeviceID = deviceID
}

// WithFields adds the fields to the get device datasource instance by Id params
func (o *GetDeviceDatasourceInstanceByIDParams) WithFields(fields *string) *GetDeviceDatasourceInstanceByIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get device datasource instance by Id params
func (o *GetDeviceDatasourceInstanceByIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithHdsID adds the hdsID to the get device datasource instance by Id params
func (o *GetDeviceDatasourceInstanceByIDParams) WithHdsID(hdsID int32) *GetDeviceDatasourceInstanceByIDParams {
	o.SetHdsID(hdsID)
	return o
}

// SetHdsID adds the hdsId to the get device datasource instance by Id params
func (o *GetDeviceDatasourceInstanceByIDParams) SetHdsID(hdsID int32) {
	o.HdsID = hdsID
}

// WithID adds the id to the get device datasource instance by Id params
func (o *GetDeviceDatasourceInstanceByIDParams) WithID(id int32) *GetDeviceDatasourceInstanceByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get device datasource instance by Id params
func (o *GetDeviceDatasourceInstanceByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceDatasourceInstanceByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
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

	// path param hdsId
	if err := r.SetPathParam("hdsId", swag.FormatInt32(o.HdsID)); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
