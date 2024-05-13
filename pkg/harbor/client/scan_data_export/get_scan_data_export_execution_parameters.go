// Code generated by go-swagger; DO NOT EDIT.

package scan_data_export

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

// NewGetScanDataExportExecutionParams creates a new GetScanDataExportExecutionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetScanDataExportExecutionParams() *GetScanDataExportExecutionParams {
	return &GetScanDataExportExecutionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetScanDataExportExecutionParamsWithTimeout creates a new GetScanDataExportExecutionParams object
// with the ability to set a timeout on a request.
func NewGetScanDataExportExecutionParamsWithTimeout(timeout time.Duration) *GetScanDataExportExecutionParams {
	return &GetScanDataExportExecutionParams{
		timeout: timeout,
	}
}

// NewGetScanDataExportExecutionParamsWithContext creates a new GetScanDataExportExecutionParams object
// with the ability to set a context for a request.
func NewGetScanDataExportExecutionParamsWithContext(ctx context.Context) *GetScanDataExportExecutionParams {
	return &GetScanDataExportExecutionParams{
		Context: ctx,
	}
}

// NewGetScanDataExportExecutionParamsWithHTTPClient creates a new GetScanDataExportExecutionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetScanDataExportExecutionParamsWithHTTPClient(client *http.Client) *GetScanDataExportExecutionParams {
	return &GetScanDataExportExecutionParams{
		HTTPClient: client,
	}
}

/*
GetScanDataExportExecutionParams contains all the parameters to send to the API endpoint

	for the get scan data export execution operation.

	Typically these are written to a http.Request.
*/
type GetScanDataExportExecutionParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string `js:"xRequestID"`

	/* ExecutionID.

	   Execution ID
	*/
	ExecutionID int64 `js:"executionID"`

	timeout    time.Duration
	Context    context.Context `js:"context"`
	HTTPClient *http.Client    `js:"httpClient"`
}

// WithDefaults hydrates default values in the get scan data export execution params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScanDataExportExecutionParams) WithDefaults() *GetScanDataExportExecutionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get scan data export execution params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScanDataExportExecutionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get scan data export execution params
func (o *GetScanDataExportExecutionParams) WithTimeout(timeout time.Duration) *GetScanDataExportExecutionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get scan data export execution params
func (o *GetScanDataExportExecutionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get scan data export execution params
func (o *GetScanDataExportExecutionParams) WithContext(ctx context.Context) *GetScanDataExportExecutionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get scan data export execution params
func (o *GetScanDataExportExecutionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get scan data export execution params
func (o *GetScanDataExportExecutionParams) WithHTTPClient(client *http.Client) *GetScanDataExportExecutionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get scan data export execution params
func (o *GetScanDataExportExecutionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get scan data export execution params
func (o *GetScanDataExportExecutionParams) WithXRequestID(xRequestID *string) *GetScanDataExportExecutionParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get scan data export execution params
func (o *GetScanDataExportExecutionParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithExecutionID adds the executionID to the get scan data export execution params
func (o *GetScanDataExportExecutionParams) WithExecutionID(executionID int64) *GetScanDataExportExecutionParams {
	o.SetExecutionID(executionID)
	return o
}

// SetExecutionID adds the executionId to the get scan data export execution params
func (o *GetScanDataExportExecutionParams) SetExecutionID(executionID int64) {
	o.ExecutionID = executionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetScanDataExportExecutionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param execution_id
	if err := r.SetPathParam("execution_id", swag.FormatInt64(o.ExecutionID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
