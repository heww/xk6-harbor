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

// NewCreateReplicationPolicyParams creates a new CreateReplicationPolicyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateReplicationPolicyParams() *CreateReplicationPolicyParams {
	return &CreateReplicationPolicyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateReplicationPolicyParamsWithTimeout creates a new CreateReplicationPolicyParams object
// with the ability to set a timeout on a request.
func NewCreateReplicationPolicyParamsWithTimeout(timeout time.Duration) *CreateReplicationPolicyParams {
	return &CreateReplicationPolicyParams{
		timeout: timeout,
	}
}

// NewCreateReplicationPolicyParamsWithContext creates a new CreateReplicationPolicyParams object
// with the ability to set a context for a request.
func NewCreateReplicationPolicyParamsWithContext(ctx context.Context) *CreateReplicationPolicyParams {
	return &CreateReplicationPolicyParams{
		Context: ctx,
	}
}

// NewCreateReplicationPolicyParamsWithHTTPClient creates a new CreateReplicationPolicyParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateReplicationPolicyParamsWithHTTPClient(client *http.Client) *CreateReplicationPolicyParams {
	return &CreateReplicationPolicyParams{
		HTTPClient: client,
	}
}

/* CreateReplicationPolicyParams contains all the parameters to send to the API endpoint
   for the create replication policy operation.

   Typically these are written to a http.Request.
*/
type CreateReplicationPolicyParams struct {

	/* Policy.

	   The replication policy
	*/
	Policy *models.ReplicationPolicy `js:"policy"`

	timeout    time.Duration
	Context    context.Context `js:"context"`
	HTTPClient *http.Client    `js:"httpClient"`
}

// WithDefaults hydrates default values in the create replication policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateReplicationPolicyParams) WithDefaults() *CreateReplicationPolicyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create replication policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateReplicationPolicyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create replication policy params
func (o *CreateReplicationPolicyParams) WithTimeout(timeout time.Duration) *CreateReplicationPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create replication policy params
func (o *CreateReplicationPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create replication policy params
func (o *CreateReplicationPolicyParams) WithContext(ctx context.Context) *CreateReplicationPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create replication policy params
func (o *CreateReplicationPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create replication policy params
func (o *CreateReplicationPolicyParams) WithHTTPClient(client *http.Client) *CreateReplicationPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create replication policy params
func (o *CreateReplicationPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPolicy adds the policy to the create replication policy params
func (o *CreateReplicationPolicyParams) WithPolicy(policy *models.ReplicationPolicy) *CreateReplicationPolicyParams {
	o.SetPolicy(policy)
	return o
}

// SetPolicy adds the policy to the create replication policy params
func (o *CreateReplicationPolicyParams) SetPolicy(policy *models.ReplicationPolicy) {
	o.Policy = policy
}

// WriteToRequest writes these params to a swagger request
func (o *CreateReplicationPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Policy != nil {
		if err := r.SetBodyParam(o.Policy); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
