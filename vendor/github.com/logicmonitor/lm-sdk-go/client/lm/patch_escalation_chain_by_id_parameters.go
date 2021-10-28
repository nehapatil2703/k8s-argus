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

	"github.com/logicmonitor/lm-sdk-go/models"
)

// NewPatchEscalationChainByIDParams creates a new PatchEscalationChainByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchEscalationChainByIDParams() *PatchEscalationChainByIDParams {
	return &PatchEscalationChainByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchEscalationChainByIDParamsWithTimeout creates a new PatchEscalationChainByIDParams object
// with the ability to set a timeout on a request.
func NewPatchEscalationChainByIDParamsWithTimeout(timeout time.Duration) *PatchEscalationChainByIDParams {
	return &PatchEscalationChainByIDParams{
		timeout: timeout,
	}
}

// NewPatchEscalationChainByIDParamsWithContext creates a new PatchEscalationChainByIDParams object
// with the ability to set a context for a request.
func NewPatchEscalationChainByIDParamsWithContext(ctx context.Context) *PatchEscalationChainByIDParams {
	return &PatchEscalationChainByIDParams{
		Context: ctx,
	}
}

// NewPatchEscalationChainByIDParamsWithHTTPClient creates a new PatchEscalationChainByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchEscalationChainByIDParamsWithHTTPClient(client *http.Client) *PatchEscalationChainByIDParams {
	return &PatchEscalationChainByIDParams{
		HTTPClient: client,
	}
}

/* PatchEscalationChainByIDParams contains all the parameters to send to the API endpoint
   for the patch escalation chain by Id operation.

   Typically these are written to a http.Request.
*/
type PatchEscalationChainByIDParams struct {

	// PatchFields.
	PatchFields *string

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v1.0.0-argus1"
	UserAgent *string

	// Body.
	Body *models.EscalatingChain

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch escalation chain by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchEscalationChainByIDParams) WithDefaults() *PatchEscalationChainByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch escalation chain by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchEscalationChainByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v1.0.0-argus1")
	)

	val := PatchEscalationChainByIDParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the patch escalation chain by Id params
func (o *PatchEscalationChainByIDParams) WithTimeout(timeout time.Duration) *PatchEscalationChainByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch escalation chain by Id params
func (o *PatchEscalationChainByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch escalation chain by Id params
func (o *PatchEscalationChainByIDParams) WithContext(ctx context.Context) *PatchEscalationChainByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch escalation chain by Id params
func (o *PatchEscalationChainByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch escalation chain by Id params
func (o *PatchEscalationChainByIDParams) WithHTTPClient(client *http.Client) *PatchEscalationChainByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch escalation chain by Id params
func (o *PatchEscalationChainByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPatchFields adds the patchFields to the patch escalation chain by Id params
func (o *PatchEscalationChainByIDParams) WithPatchFields(patchFields *string) *PatchEscalationChainByIDParams {
	o.SetPatchFields(patchFields)
	return o
}

// SetPatchFields adds the patchFields to the patch escalation chain by Id params
func (o *PatchEscalationChainByIDParams) SetPatchFields(patchFields *string) {
	o.PatchFields = patchFields
}

// WithUserAgent adds the userAgent to the patch escalation chain by Id params
func (o *PatchEscalationChainByIDParams) WithUserAgent(userAgent *string) *PatchEscalationChainByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the patch escalation chain by Id params
func (o *PatchEscalationChainByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithBody adds the body to the patch escalation chain by Id params
func (o *PatchEscalationChainByIDParams) WithBody(body *models.EscalatingChain) *PatchEscalationChainByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch escalation chain by Id params
func (o *PatchEscalationChainByIDParams) SetBody(body *models.EscalatingChain) {
	o.Body = body
}

// WithID adds the id to the patch escalation chain by Id params
func (o *PatchEscalationChainByIDParams) WithID(id int32) *PatchEscalationChainByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch escalation chain by Id params
func (o *PatchEscalationChainByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchEscalationChainByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PatchFields != nil {

		// query param PatchFields
		var qrPatchFields string

		if o.PatchFields != nil {
			qrPatchFields = *o.PatchFields
		}
		qPatchFields := qrPatchFields
		if qPatchFields != "" {

			if err := r.SetQueryParam("PatchFields", qPatchFields); err != nil {
				return err
			}
		}
	}

	if o.UserAgent != nil {

		// header param User-Agent
		if err := r.SetHeaderParam("User-Agent", *o.UserAgent); err != nil {
			return err
		}
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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
