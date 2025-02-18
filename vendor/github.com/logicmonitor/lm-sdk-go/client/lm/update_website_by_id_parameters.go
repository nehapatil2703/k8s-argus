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

// NewUpdateWebsiteByIDParams creates a new UpdateWebsiteByIDParams object
// with the default values initialized.
func NewUpdateWebsiteByIDParams() *UpdateWebsiteByIDParams {
	var ()
	return &UpdateWebsiteByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateWebsiteByIDParamsWithTimeout creates a new UpdateWebsiteByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateWebsiteByIDParamsWithTimeout(timeout time.Duration) *UpdateWebsiteByIDParams {
	var ()
	return &UpdateWebsiteByIDParams{

		timeout: timeout,
	}
}

// NewUpdateWebsiteByIDParamsWithContext creates a new UpdateWebsiteByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateWebsiteByIDParamsWithContext(ctx context.Context) *UpdateWebsiteByIDParams {
	var ()
	return &UpdateWebsiteByIDParams{

		Context: ctx,
	}
}

// NewUpdateWebsiteByIDParamsWithHTTPClient creates a new UpdateWebsiteByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateWebsiteByIDParamsWithHTTPClient(client *http.Client) *UpdateWebsiteByIDParams {
	var ()
	return &UpdateWebsiteByIDParams{
		HTTPClient: client,
	}
}

/*UpdateWebsiteByIDParams contains all the parameters to send to the API endpoint
for the update website by Id operation typically these are written to a http.Request
*/
type UpdateWebsiteByIDParams struct {

	/*Body*/
	Body models.Website
	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update website by Id params
func (o *UpdateWebsiteByIDParams) WithTimeout(timeout time.Duration) *UpdateWebsiteByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update website by Id params
func (o *UpdateWebsiteByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update website by Id params
func (o *UpdateWebsiteByIDParams) WithContext(ctx context.Context) *UpdateWebsiteByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update website by Id params
func (o *UpdateWebsiteByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update website by Id params
func (o *UpdateWebsiteByIDParams) WithHTTPClient(client *http.Client) *UpdateWebsiteByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update website by Id params
func (o *UpdateWebsiteByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update website by Id params
func (o *UpdateWebsiteByIDParams) WithBody(body models.Website) *UpdateWebsiteByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update website by Id params
func (o *UpdateWebsiteByIDParams) SetBody(body models.Website) {
	o.Body = body
}

// WithID adds the id to the update website by Id params
func (o *UpdateWebsiteByIDParams) WithID(id int32) *UpdateWebsiteByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update website by Id params
func (o *UpdateWebsiteByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateWebsiteByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
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
