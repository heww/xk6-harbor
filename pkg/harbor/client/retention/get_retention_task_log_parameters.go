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
)

// NewGetRetentionTaskLogParams creates a new GetRetentionTaskLogParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRetentionTaskLogParams() *GetRetentionTaskLogParams {
	return &GetRetentionTaskLogParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRetentionTaskLogParamsWithTimeout creates a new GetRetentionTaskLogParams object
// with the ability to set a timeout on a request.
func NewGetRetentionTaskLogParamsWithTimeout(timeout time.Duration) *GetRetentionTaskLogParams {
	return &GetRetentionTaskLogParams{
		timeout: timeout,
	}
}

// NewGetRetentionTaskLogParamsWithContext creates a new GetRetentionTaskLogParams object
// with the ability to set a context for a request.
func NewGetRetentionTaskLogParamsWithContext(ctx context.Context) *GetRetentionTaskLogParams {
	return &GetRetentionTaskLogParams{
		Context: ctx,
	}
}

// NewGetRetentionTaskLogParamsWithHTTPClient creates a new GetRetentionTaskLogParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRetentionTaskLogParamsWithHTTPClient(client *http.Client) *GetRetentionTaskLogParams {
	return &GetRetentionTaskLogParams{
		HTTPClient: client,
	}
}

/* GetRetentionTaskLogParams contains all the parameters to send to the API endpoint
   for the get retention task log operation.

   Typically these are written to a http.Request.
*/
type GetRetentionTaskLogParams struct {

	/* Eid.

	   Retention execution ID.

	   Format: int64
	*/
	Eid int64 `js:"eid"`

	/* ID.

	   Retention ID.

	   Format: int64
	*/
	ID int64 `js:"id"`

	/* Tid.

	   Retention execution ID.

	   Format: int64
	*/
	Tid int64 `js:"tid"`

	timeout    time.Duration
	Context    context.Context `js:"context"`
	HTTPClient *http.Client    `js:"httpClient"`
}

// WithDefaults hydrates default values in the get retention task log params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRetentionTaskLogParams) WithDefaults() *GetRetentionTaskLogParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get retention task log params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRetentionTaskLogParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get retention task log params
func (o *GetRetentionTaskLogParams) WithTimeout(timeout time.Duration) *GetRetentionTaskLogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get retention task log params
func (o *GetRetentionTaskLogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get retention task log params
func (o *GetRetentionTaskLogParams) WithContext(ctx context.Context) *GetRetentionTaskLogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get retention task log params
func (o *GetRetentionTaskLogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get retention task log params
func (o *GetRetentionTaskLogParams) WithHTTPClient(client *http.Client) *GetRetentionTaskLogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get retention task log params
func (o *GetRetentionTaskLogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEid adds the eid to the get retention task log params
func (o *GetRetentionTaskLogParams) WithEid(eid int64) *GetRetentionTaskLogParams {
	o.SetEid(eid)
	return o
}

// SetEid adds the eid to the get retention task log params
func (o *GetRetentionTaskLogParams) SetEid(eid int64) {
	o.Eid = eid
}

// WithID adds the id to the get retention task log params
func (o *GetRetentionTaskLogParams) WithID(id int64) *GetRetentionTaskLogParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get retention task log params
func (o *GetRetentionTaskLogParams) SetID(id int64) {
	o.ID = id
}

// WithTid adds the tid to the get retention task log params
func (o *GetRetentionTaskLogParams) WithTid(tid int64) *GetRetentionTaskLogParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the get retention task log params
func (o *GetRetentionTaskLogParams) SetTid(tid int64) {
	o.Tid = tid
}

// WriteToRequest writes these params to a swagger request
func (o *GetRetentionTaskLogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param eid
	if err := r.SetPathParam("eid", swag.FormatInt64(o.Eid)); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// path param tid
	if err := r.SetPathParam("tid", swag.FormatInt64(o.Tid)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
