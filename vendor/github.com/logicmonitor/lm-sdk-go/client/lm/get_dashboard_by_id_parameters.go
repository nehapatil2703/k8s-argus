// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetDashboardByIDParams creates a new GetDashboardByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDashboardByIDParams() *GetDashboardByIDParams {
	return &GetDashboardByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDashboardByIDParamsWithTimeout creates a new GetDashboardByIDParams object
// with the ability to set a timeout on a request.
func NewGetDashboardByIDParamsWithTimeout(timeout time.Duration) *GetDashboardByIDParams {
	return &GetDashboardByIDParams{
		timeout: timeout,
	}
}

// NewGetDashboardByIDParamsWithContext creates a new GetDashboardByIDParams object
// with the ability to set a context for a request.
func NewGetDashboardByIDParamsWithContext(ctx context.Context) *GetDashboardByIDParams {
	return &GetDashboardByIDParams{
		Context: ctx,
	}
}

// NewGetDashboardByIDParamsWithHTTPClient creates a new GetDashboardByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDashboardByIDParamsWithHTTPClient(client *http.Client) *GetDashboardByIDParams {
	return &GetDashboardByIDParams{
		HTTPClient: client,
	}
}

/* GetDashboardByIDParams contains all the parameters to send to the API endpoint
   for the get dashboard by Id operation.

   Typically these are written to a http.Request.
*/
type GetDashboardByIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v1.0.0-argus1"
	UserAgent *string

	// Fields.
	Fields *string

	// Format.
	//
	// Default: "json"
	Format *string

	// ID.
	//
	// Format: int32
	ID int32

	// Template.
	Template *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get dashboard by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDashboardByIDParams) WithDefaults() *GetDashboardByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get dashboard by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDashboardByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v1.0.0-argus1")

		formatDefault = string("json")

		templateDefault = bool(false)
	)

	val := GetDashboardByIDParams{
		UserAgent: &userAgentDefault,
		Format:    &formatDefault,
		Template:  &templateDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get dashboard by Id params
func (o *GetDashboardByIDParams) WithTimeout(timeout time.Duration) *GetDashboardByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get dashboard by Id params
func (o *GetDashboardByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get dashboard by Id params
func (o *GetDashboardByIDParams) WithContext(ctx context.Context) *GetDashboardByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get dashboard by Id params
func (o *GetDashboardByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get dashboard by Id params
func (o *GetDashboardByIDParams) WithHTTPClient(client *http.Client) *GetDashboardByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get dashboard by Id params
func (o *GetDashboardByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get dashboard by Id params
func (o *GetDashboardByIDParams) WithUserAgent(userAgent *string) *GetDashboardByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get dashboard by Id params
func (o *GetDashboardByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithFields adds the fields to the get dashboard by Id params
func (o *GetDashboardByIDParams) WithFields(fields *string) *GetDashboardByIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get dashboard by Id params
func (o *GetDashboardByIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFormat adds the format to the get dashboard by Id params
func (o *GetDashboardByIDParams) WithFormat(format *string) *GetDashboardByIDParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the get dashboard by Id params
func (o *GetDashboardByIDParams) SetFormat(format *string) {
	o.Format = format
}

// WithID adds the id to the get dashboard by Id params
func (o *GetDashboardByIDParams) WithID(id int32) *GetDashboardByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get dashboard by Id params
func (o *GetDashboardByIDParams) SetID(id int32) {
	o.ID = id
}

// WithTemplate adds the template to the get dashboard by Id params
func (o *GetDashboardByIDParams) WithTemplate(template *bool) *GetDashboardByIDParams {
	o.SetTemplate(template)
	return o
}

// SetTemplate adds the template to the get dashboard by Id params
func (o *GetDashboardByIDParams) SetTemplate(template *bool) {
	o.Template = template
}

// WriteToRequest writes these params to a swagger request
func (o *GetDashboardByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.UserAgent != nil {

		// header param User-Agent
		if err := r.SetHeaderParam("User-Agent", *o.UserAgent); err != nil {
			return err
		}
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

	if o.Template != nil {

		// query param template
		var qrTemplate bool

		if o.Template != nil {
			qrTemplate = *o.Template
		}
		qTemplate := swag.FormatBool(qrTemplate)
		if qTemplate != "" {

			if err := r.SetQueryParam("template", qTemplate); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
