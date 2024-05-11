// Code generated by go-swagger; DO NOT EDIT.

package purge

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

// NewStopPurgeParams creates a new StopPurgeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStopPurgeParams() *StopPurgeParams {
	return &StopPurgeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStopPurgeParamsWithTimeout creates a new StopPurgeParams object
// with the ability to set a timeout on a request.
func NewStopPurgeParamsWithTimeout(timeout time.Duration) *StopPurgeParams {
	return &StopPurgeParams{
		timeout: timeout,
	}
}

// NewStopPurgeParamsWithContext creates a new StopPurgeParams object
// with the ability to set a context for a request.
func NewStopPurgeParamsWithContext(ctx context.Context) *StopPurgeParams {
	return &StopPurgeParams{
		Context: ctx,
	}
}

// NewStopPurgeParamsWithHTTPClient creates a new StopPurgeParams object
// with the ability to set a custom HTTPClient for a request.
func NewStopPurgeParamsWithHTTPClient(client *http.Client) *StopPurgeParams {
	return &StopPurgeParams{
		HTTPClient: client,
	}
}

/*
StopPurgeParams contains all the parameters to send to the API endpoint

	for the stop purge operation.

	Typically these are written to a http.Request.
*/
type StopPurgeParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string `js:"xRequestID"`

	/* PurgeID.

	   The ID of the purge log

	   Format: int64
	*/
	PurgeID int64 `js:"purgeID"`

	timeout    time.Duration
	Context    context.Context `js:"context"`
	HTTPClient *http.Client    `js:"httpClient"`
}

// WithDefaults hydrates default values in the stop purge params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopPurgeParams) WithDefaults() *StopPurgeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stop purge params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopPurgeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stop purge params
func (o *StopPurgeParams) WithTimeout(timeout time.Duration) *StopPurgeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop purge params
func (o *StopPurgeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop purge params
func (o *StopPurgeParams) WithContext(ctx context.Context) *StopPurgeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop purge params
func (o *StopPurgeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop purge params
func (o *StopPurgeParams) WithHTTPClient(client *http.Client) *StopPurgeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop purge params
func (o *StopPurgeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the stop purge params
func (o *StopPurgeParams) WithXRequestID(xRequestID *string) *StopPurgeParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the stop purge params
func (o *StopPurgeParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithPurgeID adds the purgeID to the stop purge params
func (o *StopPurgeParams) WithPurgeID(purgeID int64) *StopPurgeParams {
	o.SetPurgeID(purgeID)
	return o
}

// SetPurgeID adds the purgeId to the stop purge params
func (o *StopPurgeParams) SetPurgeID(purgeID int64) {
	o.PurgeID = purgeID
}

// WriteToRequest writes these params to a swagger request
func (o *StopPurgeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}

	// path param purge_id
	if err := r.SetPathParam("purge_id", swag.FormatInt64(o.PurgeID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}