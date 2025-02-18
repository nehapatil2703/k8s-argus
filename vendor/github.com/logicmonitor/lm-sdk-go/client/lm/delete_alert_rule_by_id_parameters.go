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

// NewDeleteAlertRuleByIDParams creates a new DeleteAlertRuleByIDParams object
// with the default values initialized.
func NewDeleteAlertRuleByIDParams() *DeleteAlertRuleByIDParams {
	var ()
	return &DeleteAlertRuleByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAlertRuleByIDParamsWithTimeout creates a new DeleteAlertRuleByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAlertRuleByIDParamsWithTimeout(timeout time.Duration) *DeleteAlertRuleByIDParams {
	var ()
	return &DeleteAlertRuleByIDParams{

		timeout: timeout,
	}
}

// NewDeleteAlertRuleByIDParamsWithContext creates a new DeleteAlertRuleByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAlertRuleByIDParamsWithContext(ctx context.Context) *DeleteAlertRuleByIDParams {
	var ()
	return &DeleteAlertRuleByIDParams{

		Context: ctx,
	}
}

// NewDeleteAlertRuleByIDParamsWithHTTPClient creates a new DeleteAlertRuleByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAlertRuleByIDParamsWithHTTPClient(client *http.Client) *DeleteAlertRuleByIDParams {
	var ()
	return &DeleteAlertRuleByIDParams{
		HTTPClient: client,
	}
}

/*DeleteAlertRuleByIDParams contains all the parameters to send to the API endpoint
for the delete alert rule by Id operation typically these are written to a http.Request
*/
type DeleteAlertRuleByIDParams struct {

	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete alert rule by Id params
func (o *DeleteAlertRuleByIDParams) WithTimeout(timeout time.Duration) *DeleteAlertRuleByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete alert rule by Id params
func (o *DeleteAlertRuleByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete alert rule by Id params
func (o *DeleteAlertRuleByIDParams) WithContext(ctx context.Context) *DeleteAlertRuleByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete alert rule by Id params
func (o *DeleteAlertRuleByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete alert rule by Id params
func (o *DeleteAlertRuleByIDParams) WithHTTPClient(client *http.Client) *DeleteAlertRuleByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete alert rule by Id params
func (o *DeleteAlertRuleByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete alert rule by Id params
func (o *DeleteAlertRuleByIDParams) WithID(id int32) *DeleteAlertRuleByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete alert rule by Id params
func (o *DeleteAlertRuleByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAlertRuleByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
