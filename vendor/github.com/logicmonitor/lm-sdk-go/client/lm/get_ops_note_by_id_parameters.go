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
	"golang.org/x/net/context"
)

// NewGetOpsNoteByIDParams creates a new GetOpsNoteByIDParams object
// with the default values initialized.
func NewGetOpsNoteByIDParams() *GetOpsNoteByIDParams {
	var ()
	return &GetOpsNoteByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOpsNoteByIDParamsWithTimeout creates a new GetOpsNoteByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOpsNoteByIDParamsWithTimeout(timeout time.Duration) *GetOpsNoteByIDParams {
	var ()
	return &GetOpsNoteByIDParams{

		timeout: timeout,
	}
}

// NewGetOpsNoteByIDParamsWithContext creates a new GetOpsNoteByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOpsNoteByIDParamsWithContext(ctx context.Context) *GetOpsNoteByIDParams {
	var ()
	return &GetOpsNoteByIDParams{

		Context: ctx,
	}
}

// NewGetOpsNoteByIDParamsWithHTTPClient creates a new GetOpsNoteByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOpsNoteByIDParamsWithHTTPClient(client *http.Client) *GetOpsNoteByIDParams {
	var ()
	return &GetOpsNoteByIDParams{
		HTTPClient: client,
	}
}

/*GetOpsNoteByIDParams contains all the parameters to send to the API endpoint
for the get ops note by Id operation typically these are written to a http.Request
*/
type GetOpsNoteByIDParams struct {

	/*Fields*/
	Fields *string
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get ops note by Id params
func (o *GetOpsNoteByIDParams) WithTimeout(timeout time.Duration) *GetOpsNoteByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ops note by Id params
func (o *GetOpsNoteByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ops note by Id params
func (o *GetOpsNoteByIDParams) WithContext(ctx context.Context) *GetOpsNoteByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ops note by Id params
func (o *GetOpsNoteByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ops note by Id params
func (o *GetOpsNoteByIDParams) WithHTTPClient(client *http.Client) *GetOpsNoteByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ops note by Id params
func (o *GetOpsNoteByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get ops note by Id params
func (o *GetOpsNoteByIDParams) WithFields(fields *string) *GetOpsNoteByIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get ops note by Id params
func (o *GetOpsNoteByIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the get ops note by Id params
func (o *GetOpsNoteByIDParams) WithID(id string) *GetOpsNoteByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get ops note by Id params
func (o *GetOpsNoteByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetOpsNoteByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
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
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
