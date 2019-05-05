// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAlertRuleByIDParams creates a new GetAlertRuleByIDParams object
// with the default values initialized.
func NewGetAlertRuleByIDParams() *GetAlertRuleByIDParams {
	var ()
	return &GetAlertRuleByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAlertRuleByIDParamsWithTimeout creates a new GetAlertRuleByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAlertRuleByIDParamsWithTimeout(timeout time.Duration) *GetAlertRuleByIDParams {
	var ()
	return &GetAlertRuleByIDParams{

		timeout: timeout,
	}
}

// NewGetAlertRuleByIDParamsWithContext creates a new GetAlertRuleByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAlertRuleByIDParamsWithContext(ctx context.Context) *GetAlertRuleByIDParams {
	var ()
	return &GetAlertRuleByIDParams{

		Context: ctx,
	}
}

// NewGetAlertRuleByIDParamsWithHTTPClient creates a new GetAlertRuleByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAlertRuleByIDParamsWithHTTPClient(client *http.Client) *GetAlertRuleByIDParams {
	var ()
	return &GetAlertRuleByIDParams{
		HTTPClient: client,
	}
}

/*GetAlertRuleByIDParams contains all the parameters to send to the API endpoint
for the get alert rule by Id operation typically these are written to a http.Request
*/
type GetAlertRuleByIDParams struct {

	/*Fields*/
	Fields *string
	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get alert rule by Id params
func (o *GetAlertRuleByIDParams) WithTimeout(timeout time.Duration) *GetAlertRuleByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get alert rule by Id params
func (o *GetAlertRuleByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get alert rule by Id params
func (o *GetAlertRuleByIDParams) WithContext(ctx context.Context) *GetAlertRuleByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get alert rule by Id params
func (o *GetAlertRuleByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get alert rule by Id params
func (o *GetAlertRuleByIDParams) WithHTTPClient(client *http.Client) *GetAlertRuleByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get alert rule by Id params
func (o *GetAlertRuleByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get alert rule by Id params
func (o *GetAlertRuleByIDParams) WithFields(fields *string) *GetAlertRuleByIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get alert rule by Id params
func (o *GetAlertRuleByIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the get alert rule by Id params
func (o *GetAlertRuleByIDParams) WithID(id int32) *GetAlertRuleByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get alert rule by Id params
func (o *GetAlertRuleByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetAlertRuleByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
