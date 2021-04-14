// Code generated by go-swagger; DO NOT EDIT.

package systeminfo

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

// NewGetSysteminfoVolumesParams creates a new GetSysteminfoVolumesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSysteminfoVolumesParams() *GetSysteminfoVolumesParams {
	return &GetSysteminfoVolumesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSysteminfoVolumesParamsWithTimeout creates a new GetSysteminfoVolumesParams object
// with the ability to set a timeout on a request.
func NewGetSysteminfoVolumesParamsWithTimeout(timeout time.Duration) *GetSysteminfoVolumesParams {
	return &GetSysteminfoVolumesParams{
		timeout: timeout,
	}
}

// NewGetSysteminfoVolumesParamsWithContext creates a new GetSysteminfoVolumesParams object
// with the ability to set a context for a request.
func NewGetSysteminfoVolumesParamsWithContext(ctx context.Context) *GetSysteminfoVolumesParams {
	return &GetSysteminfoVolumesParams{
		Context: ctx,
	}
}

// NewGetSysteminfoVolumesParamsWithHTTPClient creates a new GetSysteminfoVolumesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSysteminfoVolumesParamsWithHTTPClient(client *http.Client) *GetSysteminfoVolumesParams {
	return &GetSysteminfoVolumesParams{
		HTTPClient: client,
	}
}

/* GetSysteminfoVolumesParams contains all the parameters to send to the API endpoint
   for the get systeminfo volumes operation.

   Typically these are written to a http.Request.
*/
type GetSysteminfoVolumesParams struct {
	timeout    time.Duration
	Context    context.Context `js:"context"`
	HTTPClient *http.Client    `js:"httpClient"`
}

// WithDefaults hydrates default values in the get systeminfo volumes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSysteminfoVolumesParams) WithDefaults() *GetSysteminfoVolumesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get systeminfo volumes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSysteminfoVolumesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get systeminfo volumes params
func (o *GetSysteminfoVolumesParams) WithTimeout(timeout time.Duration) *GetSysteminfoVolumesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get systeminfo volumes params
func (o *GetSysteminfoVolumesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get systeminfo volumes params
func (o *GetSysteminfoVolumesParams) WithContext(ctx context.Context) *GetSysteminfoVolumesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get systeminfo volumes params
func (o *GetSysteminfoVolumesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get systeminfo volumes params
func (o *GetSysteminfoVolumesParams) WithHTTPClient(client *http.Client) *GetSysteminfoVolumesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get systeminfo volumes params
func (o *GetSysteminfoVolumesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetSysteminfoVolumesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
