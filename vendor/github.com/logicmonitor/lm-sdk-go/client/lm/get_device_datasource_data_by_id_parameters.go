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

// NewGetDeviceDatasourceDataByIDParams creates a new GetDeviceDatasourceDataByIDParams object
// with the default values initialized.
func NewGetDeviceDatasourceDataByIDParams() *GetDeviceDatasourceDataByIDParams {
	var (
		endDefault    = int64(0)
		formatDefault = string("json")
		periodDefault = float64(1)
		startDefault  = int64(0)
	)
	return &GetDeviceDatasourceDataByIDParams{
		End:    &endDefault,
		Format: &formatDefault,
		Period: &periodDefault,
		Start:  &startDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceDatasourceDataByIDParamsWithTimeout creates a new GetDeviceDatasourceDataByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeviceDatasourceDataByIDParamsWithTimeout(timeout time.Duration) *GetDeviceDatasourceDataByIDParams {
	var (
		endDefault    = int64(0)
		formatDefault = string("json")
		periodDefault = float64(1)
		startDefault  = int64(0)
	)
	return &GetDeviceDatasourceDataByIDParams{
		End:    &endDefault,
		Format: &formatDefault,
		Period: &periodDefault,
		Start:  &startDefault,

		timeout: timeout,
	}
}

// NewGetDeviceDatasourceDataByIDParamsWithContext creates a new GetDeviceDatasourceDataByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeviceDatasourceDataByIDParamsWithContext(ctx context.Context) *GetDeviceDatasourceDataByIDParams {
	var (
		endDefault    = int64(0)
		formatDefault = string("json")
		periodDefault = float64(1)
		startDefault  = int64(0)
	)
	return &GetDeviceDatasourceDataByIDParams{
		End:    &endDefault,
		Format: &formatDefault,
		Period: &periodDefault,
		Start:  &startDefault,

		Context: ctx,
	}
}

// NewGetDeviceDatasourceDataByIDParamsWithHTTPClient creates a new GetDeviceDatasourceDataByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeviceDatasourceDataByIDParamsWithHTTPClient(client *http.Client) *GetDeviceDatasourceDataByIDParams {
	var (
		endDefault    = int64(0)
		formatDefault = string("json")
		periodDefault = float64(1)
		startDefault  = int64(0)
	)
	return &GetDeviceDatasourceDataByIDParams{
		End:        &endDefault,
		Format:     &formatDefault,
		Period:     &periodDefault,
		Start:      &startDefault,
		HTTPClient: client,
	}
}

/*GetDeviceDatasourceDataByIDParams contains all the parameters to send to the API endpoint
for the get device datasource data by Id operation typically these are written to a http.Request
*/
type GetDeviceDatasourceDataByIDParams struct {

	/*Datapoints*/
	Datapoints *string
	/*DeviceID*/
	DeviceID int32
	/*End*/
	End *int64
	/*Format*/
	Format *string
	/*ID*/
	ID int32
	/*Period*/
	Period *float64
	/*Start*/
	Start *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get device datasource data by Id params
func (o *GetDeviceDatasourceDataByIDParams) WithTimeout(timeout time.Duration) *GetDeviceDatasourceDataByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device datasource data by Id params
func (o *GetDeviceDatasourceDataByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device datasource data by Id params
func (o *GetDeviceDatasourceDataByIDParams) WithContext(ctx context.Context) *GetDeviceDatasourceDataByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device datasource data by Id params
func (o *GetDeviceDatasourceDataByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device datasource data by Id params
func (o *GetDeviceDatasourceDataByIDParams) WithHTTPClient(client *http.Client) *GetDeviceDatasourceDataByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device datasource data by Id params
func (o *GetDeviceDatasourceDataByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatapoints adds the datapoints to the get device datasource data by Id params
func (o *GetDeviceDatasourceDataByIDParams) WithDatapoints(datapoints *string) *GetDeviceDatasourceDataByIDParams {
	o.SetDatapoints(datapoints)
	return o
}

// SetDatapoints adds the datapoints to the get device datasource data by Id params
func (o *GetDeviceDatasourceDataByIDParams) SetDatapoints(datapoints *string) {
	o.Datapoints = datapoints
}

// WithDeviceID adds the deviceID to the get device datasource data by Id params
func (o *GetDeviceDatasourceDataByIDParams) WithDeviceID(deviceID int32) *GetDeviceDatasourceDataByIDParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get device datasource data by Id params
func (o *GetDeviceDatasourceDataByIDParams) SetDeviceID(deviceID int32) {
	o.DeviceID = deviceID
}

// WithEnd adds the end to the get device datasource data by Id params
func (o *GetDeviceDatasourceDataByIDParams) WithEnd(end *int64) *GetDeviceDatasourceDataByIDParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the get device datasource data by Id params
func (o *GetDeviceDatasourceDataByIDParams) SetEnd(end *int64) {
	o.End = end
}

// WithFormat adds the format to the get device datasource data by Id params
func (o *GetDeviceDatasourceDataByIDParams) WithFormat(format *string) *GetDeviceDatasourceDataByIDParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the get device datasource data by Id params
func (o *GetDeviceDatasourceDataByIDParams) SetFormat(format *string) {
	o.Format = format
}

// WithID adds the id to the get device datasource data by Id params
func (o *GetDeviceDatasourceDataByIDParams) WithID(id int32) *GetDeviceDatasourceDataByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get device datasource data by Id params
func (o *GetDeviceDatasourceDataByIDParams) SetID(id int32) {
	o.ID = id
}

// WithPeriod adds the period to the get device datasource data by Id params
func (o *GetDeviceDatasourceDataByIDParams) WithPeriod(period *float64) *GetDeviceDatasourceDataByIDParams {
	o.SetPeriod(period)
	return o
}

// SetPeriod adds the period to the get device datasource data by Id params
func (o *GetDeviceDatasourceDataByIDParams) SetPeriod(period *float64) {
	o.Period = period
}

// WithStart adds the start to the get device datasource data by Id params
func (o *GetDeviceDatasourceDataByIDParams) WithStart(start *int64) *GetDeviceDatasourceDataByIDParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get device datasource data by Id params
func (o *GetDeviceDatasourceDataByIDParams) SetStart(start *int64) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceDatasourceDataByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Datapoints != nil {

		// query param datapoints
		var qrDatapoints string
		if o.Datapoints != nil {
			qrDatapoints = *o.Datapoints
		}
		qDatapoints := qrDatapoints
		if qDatapoints != "" {
			if err := r.SetQueryParam("datapoints", qDatapoints); err != nil {
				return err
			}
		}

	}

	// path param deviceId
	if err := r.SetPathParam("deviceId", swag.FormatInt32(o.DeviceID)); err != nil {
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

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if o.Period != nil {

		// query param period
		var qrPeriod float64
		if o.Period != nil {
			qrPeriod = *o.Period
		}
		qPeriod := swag.FormatFloat64(qrPeriod)
		if qPeriod != "" {
			if err := r.SetQueryParam("period", qPeriod); err != nil {
				return err
			}
		}

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
