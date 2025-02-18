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

// NewUpdateDeviceGroupDatasourceAlertSettingParams creates a new UpdateDeviceGroupDatasourceAlertSettingParams object
// with the default values initialized.
func NewUpdateDeviceGroupDatasourceAlertSettingParams() *UpdateDeviceGroupDatasourceAlertSettingParams {
	var ()
	return &UpdateDeviceGroupDatasourceAlertSettingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDeviceGroupDatasourceAlertSettingParamsWithTimeout creates a new UpdateDeviceGroupDatasourceAlertSettingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateDeviceGroupDatasourceAlertSettingParamsWithTimeout(timeout time.Duration) *UpdateDeviceGroupDatasourceAlertSettingParams {
	var ()
	return &UpdateDeviceGroupDatasourceAlertSettingParams{

		timeout: timeout,
	}
}

// NewUpdateDeviceGroupDatasourceAlertSettingParamsWithContext creates a new UpdateDeviceGroupDatasourceAlertSettingParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateDeviceGroupDatasourceAlertSettingParamsWithContext(ctx context.Context) *UpdateDeviceGroupDatasourceAlertSettingParams {
	var ()
	return &UpdateDeviceGroupDatasourceAlertSettingParams{

		Context: ctx,
	}
}

// NewUpdateDeviceGroupDatasourceAlertSettingParamsWithHTTPClient creates a new UpdateDeviceGroupDatasourceAlertSettingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateDeviceGroupDatasourceAlertSettingParamsWithHTTPClient(client *http.Client) *UpdateDeviceGroupDatasourceAlertSettingParams {
	var ()
	return &UpdateDeviceGroupDatasourceAlertSettingParams{
		HTTPClient: client,
	}
}

/*UpdateDeviceGroupDatasourceAlertSettingParams contains all the parameters to send to the API endpoint
for the update device group datasource alert setting operation typically these are written to a http.Request
*/
type UpdateDeviceGroupDatasourceAlertSettingParams struct {

	/*Body*/
	Body *models.DeviceGroupDataSourceAlertConfig
	/*DeviceGroupID*/
	DeviceGroupID int32
	/*DsID*/
	DsID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update device group datasource alert setting params
func (o *UpdateDeviceGroupDatasourceAlertSettingParams) WithTimeout(timeout time.Duration) *UpdateDeviceGroupDatasourceAlertSettingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update device group datasource alert setting params
func (o *UpdateDeviceGroupDatasourceAlertSettingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update device group datasource alert setting params
func (o *UpdateDeviceGroupDatasourceAlertSettingParams) WithContext(ctx context.Context) *UpdateDeviceGroupDatasourceAlertSettingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update device group datasource alert setting params
func (o *UpdateDeviceGroupDatasourceAlertSettingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update device group datasource alert setting params
func (o *UpdateDeviceGroupDatasourceAlertSettingParams) WithHTTPClient(client *http.Client) *UpdateDeviceGroupDatasourceAlertSettingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update device group datasource alert setting params
func (o *UpdateDeviceGroupDatasourceAlertSettingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update device group datasource alert setting params
func (o *UpdateDeviceGroupDatasourceAlertSettingParams) WithBody(body *models.DeviceGroupDataSourceAlertConfig) *UpdateDeviceGroupDatasourceAlertSettingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update device group datasource alert setting params
func (o *UpdateDeviceGroupDatasourceAlertSettingParams) SetBody(body *models.DeviceGroupDataSourceAlertConfig) {
	o.Body = body
}

// WithDeviceGroupID adds the deviceGroupID to the update device group datasource alert setting params
func (o *UpdateDeviceGroupDatasourceAlertSettingParams) WithDeviceGroupID(deviceGroupID int32) *UpdateDeviceGroupDatasourceAlertSettingParams {
	o.SetDeviceGroupID(deviceGroupID)
	return o
}

// SetDeviceGroupID adds the deviceGroupId to the update device group datasource alert setting params
func (o *UpdateDeviceGroupDatasourceAlertSettingParams) SetDeviceGroupID(deviceGroupID int32) {
	o.DeviceGroupID = deviceGroupID
}

// WithDsID adds the dsID to the update device group datasource alert setting params
func (o *UpdateDeviceGroupDatasourceAlertSettingParams) WithDsID(dsID int32) *UpdateDeviceGroupDatasourceAlertSettingParams {
	o.SetDsID(dsID)
	return o
}

// SetDsID adds the dsId to the update device group datasource alert setting params
func (o *UpdateDeviceGroupDatasourceAlertSettingParams) SetDsID(dsID int32) {
	o.DsID = dsID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDeviceGroupDatasourceAlertSettingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
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

	// path param dsId
	if err := r.SetPathParam("dsId", swag.FormatInt32(o.DsID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
