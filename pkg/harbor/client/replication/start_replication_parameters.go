// Code generated by go-swagger; DO NOT EDIT.

package replication

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

	"github.com/heww/xk6-harbor/pkg/harbor/models"
)

// NewStartReplicationParams creates a new StartReplicationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStartReplicationParams() *StartReplicationParams {
	return &StartReplicationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStartReplicationParamsWithTimeout creates a new StartReplicationParams object
// with the ability to set a timeout on a request.
func NewStartReplicationParamsWithTimeout(timeout time.Duration) *StartReplicationParams {
	return &StartReplicationParams{
		timeout: timeout,
	}
}

// NewStartReplicationParamsWithContext creates a new StartReplicationParams object
// with the ability to set a context for a request.
func NewStartReplicationParamsWithContext(ctx context.Context) *StartReplicationParams {
	return &StartReplicationParams{
		Context: ctx,
	}
}

// NewStartReplicationParamsWithHTTPClient creates a new StartReplicationParams object
// with the ability to set a custom HTTPClient for a request.
func NewStartReplicationParamsWithHTTPClient(client *http.Client) *StartReplicationParams {
	return &StartReplicationParams{
		HTTPClient: client,
	}
}

/* StartReplicationParams contains all the parameters to send to the API endpoint
   for the start replication operation.

   Typically these are written to a http.Request.
*/
type StartReplicationParams struct {

	/* Execution.

	   The ID of policy that the execution belongs to
	*/
	Execution *models.StartReplicationExecution `js:"execution"`

	timeout    time.Duration
	Context    context.Context `js:"context"`
	HTTPClient *http.Client    `js:"httpClient"`
}

// WithDefaults hydrates default values in the start replication params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartReplicationParams) WithDefaults() *StartReplicationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the start replication params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartReplicationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the start replication params
func (o *StartReplicationParams) WithTimeout(timeout time.Duration) *StartReplicationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start replication params
func (o *StartReplicationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start replication params
func (o *StartReplicationParams) WithContext(ctx context.Context) *StartReplicationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start replication params
func (o *StartReplicationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start replication params
func (o *StartReplicationParams) WithHTTPClient(client *http.Client) *StartReplicationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start replication params
func (o *StartReplicationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExecution adds the execution to the start replication params
func (o *StartReplicationParams) WithExecution(execution *models.StartReplicationExecution) *StartReplicationParams {
	o.SetExecution(execution)
	return o
}

// SetExecution adds the execution to the start replication params
func (o *StartReplicationParams) SetExecution(execution *models.StartReplicationExecution) {
	o.Execution = execution
}

// WriteToRequest writes these params to a swagger request
func (o *StartReplicationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Execution != nil {
		if err := r.SetBodyParam(o.Execution); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
