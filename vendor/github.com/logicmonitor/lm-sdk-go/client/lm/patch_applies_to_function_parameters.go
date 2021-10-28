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

// NewPatchAppliesToFunctionParams creates a new PatchAppliesToFunctionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchAppliesToFunctionParams() *PatchAppliesToFunctionParams {
	return &PatchAppliesToFunctionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAppliesToFunctionParamsWithTimeout creates a new PatchAppliesToFunctionParams object
// with the ability to set a timeout on a request.
func NewPatchAppliesToFunctionParamsWithTimeout(timeout time.Duration) *PatchAppliesToFunctionParams {
	return &PatchAppliesToFunctionParams{
		timeout: timeout,
	}
}

// NewPatchAppliesToFunctionParamsWithContext creates a new PatchAppliesToFunctionParams object
// with the ability to set a context for a request.
func NewPatchAppliesToFunctionParamsWithContext(ctx context.Context) *PatchAppliesToFunctionParams {
	return &PatchAppliesToFunctionParams{
		Context: ctx,
	}
}

// NewPatchAppliesToFunctionParamsWithHTTPClient creates a new PatchAppliesToFunctionParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchAppliesToFunctionParamsWithHTTPClient(client *http.Client) *PatchAppliesToFunctionParams {
	return &PatchAppliesToFunctionParams{
		HTTPClient: client,
	}
}

/* PatchAppliesToFunctionParams contains all the parameters to send to the API endpoint
   for the patch applies to function operation.

   Typically these are written to a http.Request.
*/
type PatchAppliesToFunctionParams struct {

	// PatchFields.
	PatchFields *string

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v1.0.0-argus1"
	UserAgent *string

	// Body.
	Body *models.AppliesToFunction

	// ID.
	//
	// Format: int32
	ID int32

	// IgnoreReference.
	IgnoreReference *bool

	// Reason.
	Reason *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch applies to function params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAppliesToFunctionParams) WithDefaults() *PatchAppliesToFunctionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch applies to function params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAppliesToFunctionParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v1.0.0-argus1")

		ignoreReferenceDefault = bool(false)
	)

	val := PatchAppliesToFunctionParams{
		UserAgent:       &userAgentDefault,
		IgnoreReference: &ignoreReferenceDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the patch applies to function params
func (o *PatchAppliesToFunctionParams) WithTimeout(timeout time.Duration) *PatchAppliesToFunctionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch applies to function params
func (o *PatchAppliesToFunctionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch applies to function params
func (o *PatchAppliesToFunctionParams) WithContext(ctx context.Context) *PatchAppliesToFunctionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch applies to function params
func (o *PatchAppliesToFunctionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch applies to function params
func (o *PatchAppliesToFunctionParams) WithHTTPClient(client *http.Client) *PatchAppliesToFunctionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch applies to function params
func (o *PatchAppliesToFunctionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPatchFields adds the patchFields to the patch applies to function params
func (o *PatchAppliesToFunctionParams) WithPatchFields(patchFields *string) *PatchAppliesToFunctionParams {
	o.SetPatchFields(patchFields)
	return o
}

// SetPatchFields adds the patchFields to the patch applies to function params
func (o *PatchAppliesToFunctionParams) SetPatchFields(patchFields *string) {
	o.PatchFields = patchFields
}

// WithUserAgent adds the userAgent to the patch applies to function params
func (o *PatchAppliesToFunctionParams) WithUserAgent(userAgent *string) *PatchAppliesToFunctionParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the patch applies to function params
func (o *PatchAppliesToFunctionParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithBody adds the body to the patch applies to function params
func (o *PatchAppliesToFunctionParams) WithBody(body *models.AppliesToFunction) *PatchAppliesToFunctionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch applies to function params
func (o *PatchAppliesToFunctionParams) SetBody(body *models.AppliesToFunction) {
	o.Body = body
}

// WithID adds the id to the patch applies to function params
func (o *PatchAppliesToFunctionParams) WithID(id int32) *PatchAppliesToFunctionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch applies to function params
func (o *PatchAppliesToFunctionParams) SetID(id int32) {
	o.ID = id
}

// WithIgnoreReference adds the ignoreReference to the patch applies to function params
func (o *PatchAppliesToFunctionParams) WithIgnoreReference(ignoreReference *bool) *PatchAppliesToFunctionParams {
	o.SetIgnoreReference(ignoreReference)
	return o
}

// SetIgnoreReference adds the ignoreReference to the patch applies to function params
func (o *PatchAppliesToFunctionParams) SetIgnoreReference(ignoreReference *bool) {
	o.IgnoreReference = ignoreReference
}

// WithReason adds the reason to the patch applies to function params
func (o *PatchAppliesToFunctionParams) WithReason(reason *string) *PatchAppliesToFunctionParams {
	o.SetReason(reason)
	return o
}

// SetReason adds the reason to the patch applies to function params
func (o *PatchAppliesToFunctionParams) SetReason(reason *string) {
	o.Reason = reason
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAppliesToFunctionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.IgnoreReference != nil {

		// query param ignoreReference
		var qrIgnoreReference bool

		if o.IgnoreReference != nil {
			qrIgnoreReference = *o.IgnoreReference
		}
		qIgnoreReference := swag.FormatBool(qrIgnoreReference)
		if qIgnoreReference != "" {

			if err := r.SetQueryParam("ignoreReference", qIgnoreReference); err != nil {
				return err
			}
		}
	}

	if o.Reason != nil {

		// query param reason
		var qrReason string

		if o.Reason != nil {
			qrReason = *o.Reason
		}
		qReason := qrReason
		if qReason != "" {

			if err := r.SetQueryParam("reason", qReason); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
