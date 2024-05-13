// Code generated by go-swagger; DO NOT EDIT.

package jobservice

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
)

// NewStopRunningJobParams creates a new StopRunningJobParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStopRunningJobParams() *StopRunningJobParams {
	return &StopRunningJobParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStopRunningJobParamsWithTimeout creates a new StopRunningJobParams object
// with the ability to set a timeout on a request.
func NewStopRunningJobParamsWithTimeout(timeout time.Duration) *StopRunningJobParams {
	return &StopRunningJobParams{
		timeout: timeout,
	}
}

// NewStopRunningJobParamsWithContext creates a new StopRunningJobParams object
// with the ability to set a context for a request.
func NewStopRunningJobParamsWithContext(ctx context.Context) *StopRunningJobParams {
	return &StopRunningJobParams{
		Context: ctx,
	}
}

// NewStopRunningJobParamsWithHTTPClient creates a new StopRunningJobParams object
// with the ability to set a custom HTTPClient for a request.
func NewStopRunningJobParamsWithHTTPClient(client *http.Client) *StopRunningJobParams {
	return &StopRunningJobParams{
		HTTPClient: client,
	}
}

/*
StopRunningJobParams contains all the parameters to send to the API endpoint

	for the stop running job operation.

	Typically these are written to a http.Request.
*/
type StopRunningJobParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string `js:"xRequestID"`

	/* JobID.

	   The id of the job.
	*/
	JobID string `js:"jobID"`

	timeout    time.Duration
	Context    context.Context `js:"context"`
	HTTPClient *http.Client    `js:"httpClient"`
}

// WithDefaults hydrates default values in the stop running job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopRunningJobParams) WithDefaults() *StopRunningJobParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stop running job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopRunningJobParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stop running job params
func (o *StopRunningJobParams) WithTimeout(timeout time.Duration) *StopRunningJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop running job params
func (o *StopRunningJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop running job params
func (o *StopRunningJobParams) WithContext(ctx context.Context) *StopRunningJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop running job params
func (o *StopRunningJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop running job params
func (o *StopRunningJobParams) WithHTTPClient(client *http.Client) *StopRunningJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop running job params
func (o *StopRunningJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the stop running job params
func (o *StopRunningJobParams) WithXRequestID(xRequestID *string) *StopRunningJobParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the stop running job params
func (o *StopRunningJobParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithJobID adds the jobID to the stop running job params
func (o *StopRunningJobParams) WithJobID(jobID string) *StopRunningJobParams {
	o.SetJobID(jobID)
	return o
}

// SetJobID adds the jobId to the stop running job params
func (o *StopRunningJobParams) SetJobID(jobID string) {
	o.JobID = jobID
}

// WriteToRequest writes these params to a swagger request
func (o *StopRunningJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param job_id
	if err := r.SetPathParam("job_id", o.JobID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
