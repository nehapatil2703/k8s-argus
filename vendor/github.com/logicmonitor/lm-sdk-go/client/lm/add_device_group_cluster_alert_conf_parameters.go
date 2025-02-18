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
	models "github.com/logicmonitor/lm-sdk-go/models"
	"golang.org/x/net/context"
)

// NewAddDeviceGroupClusterAlertConfParams creates a new AddDeviceGroupClusterAlertConfParams object
// with the default values initialized.
func NewAddDeviceGroupClusterAlertConfParams() *AddDeviceGroupClusterAlertConfParams {
	var ()
	return &AddDeviceGroupClusterAlertConfParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddDeviceGroupClusterAlertConfParamsWithTimeout creates a new AddDeviceGroupClusterAlertConfParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddDeviceGroupClusterAlertConfParamsWithTimeout(timeout time.Duration) *AddDeviceGroupClusterAlertConfParams {
	var ()
	return &AddDeviceGroupClusterAlertConfParams{

		timeout: timeout,
	}
}

// NewAddDeviceGroupClusterAlertConfParamsWithContext creates a new AddDeviceGroupClusterAlertConfParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddDeviceGroupClusterAlertConfParamsWithContext(ctx context.Context) *AddDeviceGroupClusterAlertConfParams {
	var ()
	return &AddDeviceGroupClusterAlertConfParams{

		Context: ctx,
	}
}

// NewAddDeviceGroupClusterAlertConfParamsWithHTTPClient creates a new AddDeviceGroupClusterAlertConfParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddDeviceGroupClusterAlertConfParamsWithHTTPClient(client *http.Client) *AddDeviceGroupClusterAlertConfParams {
	var ()
	return &AddDeviceGroupClusterAlertConfParams{
		HTTPClient: client,
	}
}

/*AddDeviceGroupClusterAlertConfParams contains all the parameters to send to the API endpoint
for the add device group cluster alert conf operation typically these are written to a http.Request
*/
type AddDeviceGroupClusterAlertConfParams struct {

	/*Body*/
	Body *models.DeviceClusterAlertConfig
	/*DeviceGroupID*/
	DeviceGroupID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add device group cluster alert conf params
func (o *AddDeviceGroupClusterAlertConfParams) WithTimeout(timeout time.Duration) *AddDeviceGroupClusterAlertConfParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add device group cluster alert conf params
func (o *AddDeviceGroupClusterAlertConfParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add device group cluster alert conf params
func (o *AddDeviceGroupClusterAlertConfParams) WithContext(ctx context.Context) *AddDeviceGroupClusterAlertConfParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add device group cluster alert conf params
func (o *AddDeviceGroupClusterAlertConfParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add device group cluster alert conf params
func (o *AddDeviceGroupClusterAlertConfParams) WithHTTPClient(client *http.Client) *AddDeviceGroupClusterAlertConfParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add device group cluster alert conf params
func (o *AddDeviceGroupClusterAlertConfParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add device group cluster alert conf params
func (o *AddDeviceGroupClusterAlertConfParams) WithBody(body *models.DeviceClusterAlertConfig) *AddDeviceGroupClusterAlertConfParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add device group cluster alert conf params
func (o *AddDeviceGroupClusterAlertConfParams) SetBody(body *models.DeviceClusterAlertConfig) {
	o.Body = body
}

// WithDeviceGroupID adds the deviceGroupID to the add device group cluster alert conf params
func (o *AddDeviceGroupClusterAlertConfParams) WithDeviceGroupID(deviceGroupID int32) *AddDeviceGroupClusterAlertConfParams {
	o.SetDeviceGroupID(deviceGroupID)
	return o
}

// SetDeviceGroupID adds the deviceGroupId to the add device group cluster alert conf params
func (o *AddDeviceGroupClusterAlertConfParams) SetDeviceGroupID(deviceGroupID int32) {
	o.DeviceGroupID = deviceGroupID
}

// WriteToRequest writes these params to a swagger request
func (o *AddDeviceGroupClusterAlertConfParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param deviceGroupId
	if err := r.SetPathParam("deviceGroupId", swag.FormatInt32(o.DeviceGroupID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
