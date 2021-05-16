// Code generated by go-swagger; DO NOT EDIT.

package gc

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

// NewGetGCParams creates a new GetGCParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetGCParams() *GetGCParams {
	return &GetGCParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetGCParamsWithTimeout creates a new GetGCParams object
// with the ability to set a timeout on a request.
func NewGetGCParamsWithTimeout(timeout time.Duration) *GetGCParams {
	return &GetGCParams{
		timeout: timeout,
	}
}

// NewGetGCParamsWithContext creates a new GetGCParams object
// with the ability to set a context for a request.
func NewGetGCParamsWithContext(ctx context.Context) *GetGCParams {
	return &GetGCParams{
		Context: ctx,
	}
}

// NewGetGCParamsWithHTTPClient creates a new GetGCParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetGCParamsWithHTTPClient(client *http.Client) *GetGCParams {
	return &GetGCParams{
		HTTPClient: client,
	}
}

/* GetGCParams contains all the parameters to send to the API endpoint
   for the get GC operation.

   Typically these are written to a http.Request.
*/
type GetGCParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string `js:"xRequestID"`

	/* GCID.

	   The ID of the gc log

	   Format: int64
	*/
	GCID int64 `js:"gcid"`

	timeout    time.Duration
	Context    context.Context `js:"context"`
	HTTPClient *http.Client    `js:"httpClient"`
}

// WithDefaults hydrates default values in the get GC params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGCParams) WithDefaults() *GetGCParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get GC params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGCParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get GC params
func (o *GetGCParams) WithTimeout(timeout time.Duration) *GetGCParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get GC params
func (o *GetGCParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get GC params
func (o *GetGCParams) WithContext(ctx context.Context) *GetGCParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get GC params
func (o *GetGCParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get GC params
func (o *GetGCParams) WithHTTPClient(client *http.Client) *GetGCParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get GC params
func (o *GetGCParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get GC params
func (o *GetGCParams) WithXRequestID(xRequestID *string) *GetGCParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get GC params
func (o *GetGCParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithGCID adds the gCID to the get GC params
func (o *GetGCParams) WithGCID(gCID int64) *GetGCParams {
	o.SetGCID(gCID)
	return o
}

// SetGCID adds the gcId to the get GC params
func (o *GetGCParams) SetGCID(gCID int64) {
	o.GCID = gCID
}

// WriteToRequest writes these params to a swagger request
func (o *GetGCParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param gc_id
	if err := r.SetPathParam("gc_id", swag.FormatInt64(o.GCID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
