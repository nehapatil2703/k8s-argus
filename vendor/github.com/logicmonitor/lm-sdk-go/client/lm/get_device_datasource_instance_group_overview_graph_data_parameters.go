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

// NewGetDeviceDatasourceInstanceGroupOverviewGraphDataParams creates a new GetDeviceDatasourceInstanceGroupOverviewGraphDataParams object
// with the default values initialized.
func NewGetDeviceDatasourceInstanceGroupOverviewGraphDataParams() *GetDeviceDatasourceInstanceGroupOverviewGraphDataParams {
	var ()
	return &GetDeviceDatasourceInstanceGroupOverviewGraphDataParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceDatasourceInstanceGroupOverviewGraphDataParamsWithTimeout creates a new GetDeviceDatasourceInstanceGroupOverviewGraphDataParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeviceDatasourceInstanceGroupOverviewGraphDataParamsWithTimeout(timeout time.Duration) *GetDeviceDatasourceInstanceGroupOverviewGraphDataParams {
	var ()
	return &GetDeviceDatasourceInstanceGroupOverviewGraphDataParams{

		timeout: timeout,
	}
}

// NewGetDeviceDatasourceInstanceGroupOverviewGraphDataParamsWithContext creates a new GetDeviceDatasourceInstanceGroupOverviewGraphDataParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeviceDatasourceInstanceGroupOverviewGraphDataParamsWithContext(ctx context.Context) *GetDeviceDatasourceInstanceGroupOverviewGraphDataParams {
	var ()
	return &GetDeviceDatasourceInstanceGroupOverviewGraphDataParams{

		Context: ctx,
	}
}

// NewGetDeviceDatasourceInstanceGroupOverviewGraphDataParamsWithHTTPClient creates a new GetDeviceDatasourceInstanceGroupOverviewGraphDataParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeviceDatasourceInstanceGroupOverviewGraphDataParamsWithHTTPClient(client *http.Client) *GetDeviceDatasourceInstanceGroupOverviewGraphDataParams {
	var ()
	return &GetDeviceDatasourceInstanceGroupOverviewGraphDataParams{
		HTTPClient: client,
	}
}

/*GetDeviceDatasourceInstanceGroupOverviewGraphDataParams contains all the parameters to send to the API endpoint
for the get device datasource instance group overview graph data operation typically these are written to a http.Request
*/
type GetDeviceDatasourceInstanceGroupOverviewGraphDataParams struct {

	/*DeviceDsID
	  The device-datasource ID you'd like to add an instance group for

	*/
	DeviceDsID int32
	/*DeviceID*/
	DeviceID int32
	/*DsigID*/
	DsigID int32
	/*End*/
	End *int64
	/*Format*/
	Format *string
	/*OgraphID*/
	OgraphID int32
	/*Start*/
	Start *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get device datasource instance group overview graph data params
func (o *GetDeviceDatasourceInstanceGroupOverviewGraphDataParams) WithTimeout(timeout time.Duration) *GetDeviceDatasourceInstanceGroupOverviewGraphDataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device datasource instance group overview graph data params
func (o *GetDeviceDatasourceInstanceGroupOverviewGraphDataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device datasource instance group overview graph data params
func (o *GetDeviceDatasourceInstanceGroupOverviewGraphDataParams) WithContext(ctx context.Context) *GetDeviceDatasourceInstanceGroupOverviewGraphDataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device datasource instance group overview graph data params
func (o *GetDeviceDatasourceInstanceGroupOverviewGraphDataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device datasource instance group overview graph data params
func (o *GetDeviceDatasourceInstanceGroupOverviewGraphDataParams) WithHTTPClient(client *http.Client) *GetDeviceDatasourceInstanceGroupOverviewGraphDataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device datasource instance group overview graph data params
func (o *GetDeviceDatasourceInstanceGroupOverviewGraphDataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceDsID adds the deviceDsID to the get device datasource instance group overview graph data params
func (o *GetDeviceDatasourceInstanceGroupOverviewGraphDataParams) WithDeviceDsID(deviceDsID int32) *GetDeviceDatasourceInstanceGroupOverviewGraphDataParams {
	o.SetDeviceDsID(deviceDsID)
	return o
}

// SetDeviceDsID adds the deviceDsId to the get device datasource instance group overview graph data params
func (o *GetDeviceDatasourceInstanceGroupOverviewGraphDataParams) SetDeviceDsID(deviceDsID int32) {
	o.DeviceDsID = deviceDsID
}

// WithDeviceID adds the deviceID to the get device datasource instance group overview graph data params
func (o *GetDeviceDatasourceInstanceGroupOverviewGraphDataParams) WithDeviceID(deviceID int32) *GetDeviceDatasourceInstanceGroupOverviewGraphDataParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get device datasource instance group overview graph data params
func (o *GetDeviceDatasourceInstanceGroupOverviewGraphDataParams) SetDeviceID(deviceID int32) {
	o.DeviceID = deviceID
}

// WithDsigID adds the dsigID to the get device datasource instance group overview graph data params
func (o *GetDeviceDatasourceInstanceGroupOverviewGraphDataParams) WithDsigID(dsigID int32) *GetDeviceDatasourceInstanceGroupOverviewGraphDataParams {
	o.SetDsigID(dsigID)
	return o
}

// SetDsigID adds the dsigId to the get device datasource instance group overview graph data params
func (o *GetDeviceDatasourceInstanceGroupOverviewGraphDataParams) SetDsigID(dsigID int32) {
	o.DsigID = dsigID
}

// WithEnd adds the end to the get device datasource instance group overview graph data params
func (o *GetDeviceDatasourceInstanceGroupOverviewGraphDataParams) WithEnd(end *int64) *GetDeviceDatasourceInstanceGroupOverviewGraphDataParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the get device datasource instance group overview graph data params
func (o *GetDeviceDatasourceInstanceGroupOverviewGraphDataParams) SetEnd(end *int64) {
	o.End = end
}

// WithFormat adds the format to the get device datasource instance group overview graph data params
func (o *GetDeviceDatasourceInstanceGroupOverviewGraphDataParams) WithFormat(format *string) *GetDeviceDatasourceInstanceGroupOverviewGraphDataParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the get device datasource instance group overview graph data params
func (o *GetDeviceDatasourceInstanceGroupOverviewGraphDataParams) SetFormat(format *string) {
	o.Format = format
}

// WithOgraphID adds the ographID to the get device datasource instance group overview graph data params
func (o *GetDeviceDatasourceInstanceGroupOverviewGraphDataParams) WithOgraphID(ographID int32) *GetDeviceDatasourceInstanceGroupOverviewGraphDataParams {
	o.SetOgraphID(ographID)
	return o
}

// SetOgraphID adds the ographId to the get device datasource instance group overview graph data params
func (o *GetDeviceDatasourceInstanceGroupOverviewGraphDataParams) SetOgraphID(ographID int32) {
	o.OgraphID = ographID
}

// WithStart adds the start to the get device datasource instance group overview graph data params
func (o *GetDeviceDatasourceInstanceGroupOverviewGraphDataParams) WithStart(start *int64) *GetDeviceDatasourceInstanceGroupOverviewGraphDataParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get device datasource instance group overview graph data params
func (o *GetDeviceDatasourceInstanceGroupOverviewGraphDataParams) SetStart(start *int64) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceDatasourceInstanceGroupOverviewGraphDataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deviceDsId
	if err := r.SetPathParam("deviceDsId", swag.FormatInt32(o.DeviceDsID)); err != nil {
		return err
	}

	// path param deviceId
	if err := r.SetPathParam("deviceId", swag.FormatInt32(o.DeviceID)); err != nil {
		return err
	}

	// path param dsigId
	if err := r.SetPathParam("dsigId", swag.FormatInt32(o.DsigID)); err != nil {
		return err
	}

	if o.End != nil {

		// query param end
		var qrEnd int64
		if o.End != nil {
			qrEnd = *o.End
		}
		qEnd := swag.FormatInt64(qrEnd)
		if qEnd != "" {
			if err := r.SetQueryParam("end", qEnd); err != nil {
				return err
			}
		}

	}

	if o.Format != nil {

		// query param format
		var qrFormat string
		if o.Format != nil {
			qrFormat = *o.Format
		}
		qFormat := qrFormat
		if qFormat != "" {
			if err := r.SetQueryParam("format", qFormat); err != nil {
				return err
			}
		}

	}

	// path param ographId
	if err := r.SetPathParam("ographId", swag.FormatInt32(o.OgraphID)); err != nil {
		return err
	}

	if o.Start != nil {

		// query param start
		var qrStart int64
		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := swag.FormatInt64(qrStart)
		if qStart != "" {
			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
