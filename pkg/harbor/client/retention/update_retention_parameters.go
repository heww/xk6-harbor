// Code generated by go-swagger; DO NOT EDIT.

package retention

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

	"github.com/heww/xk6-harbor/pkg/harbor/models"
)

// NewUpdateRetentionParams creates a new UpdateRetentionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateRetentionParams() *UpdateRetentionParams {
	return &UpdateRetentionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRetentionParamsWithTimeout creates a new UpdateRetentionParams object
// with the ability to set a timeout on a request.
func NewUpdateRetentionParamsWithTimeout(timeout time.Duration) *UpdateRetentionParams {
	return &UpdateRetentionParams{
		timeout: timeout,
	}
}

// NewUpdateRetentionParamsWithContext creates a new UpdateRetentionParams object
// with the ability to set a context for a request.
func NewUpdateRetentionParamsWithContext(ctx context.Context) *UpdateRetentionParams {
	return &UpdateRetentionParams{
		Context: ctx,
	}
}

// NewUpdateRetentionParamsWithHTTPClient creates a new UpdateRetentionParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateRetentionParamsWithHTTPClient(client *http.Client) *UpdateRetentionParams {
	return &UpdateRetentionParams{
		HTTPClient: client,
	}
}

/* UpdateRetentionParams contains all the parameters to send to the API endpoint
   for the update retention operation.

   Typically these are written to a http.Request.
*/
type UpdateRetentionParams struct {

	/* ID.

	   Retention ID.

	   Format: int64
	*/
	ID int64 `js:"id"`

	// Policy.
	Policy *models.RetentionPolicy `js:"policy"`

	timeout    time.Duration
	Context    context.Context `js:"context"`
	HTTPClient *http.Client    `js:"httpClient"`
}

// WithDefaults hydrates default values in the update retention params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRetentionParams) WithDefaults() *UpdateRetentionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update retention params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRetentionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update retention params
func (o *UpdateRetentionParams) WithTimeout(timeout time.Duration) *UpdateRetentionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update retention params
func (o *UpdateRetentionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update retention params
func (o *UpdateRetentionParams) WithContext(ctx context.Context) *UpdateRetentionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update retention params
func (o *UpdateRetentionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update retention params
func (o *UpdateRetentionParams) WithHTTPClient(client *http.Client) *UpdateRetentionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update retention params
func (o *UpdateRetentionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update retention params
func (o *UpdateRetentionParams) WithID(id int64) *UpdateRetentionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update retention params
func (o *UpdateRetentionParams) SetID(id int64) {
	o.ID = id
}

// WithPolicy adds the policy to the update retention params
func (o *UpdateRetentionParams) WithPolicy(policy *models.RetentionPolicy) *UpdateRetentionParams {
	o.SetPolicy(policy)
	return o
}

// SetPolicy adds the policy to the update retention params
func (o *UpdateRetentionParams) SetPolicy(policy *models.RetentionPolicy) {
	o.Policy = policy
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRetentionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}
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
