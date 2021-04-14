// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewDeleteProjectParams creates a new DeleteProjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteProjectParams() *DeleteProjectParams {
	return &DeleteProjectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteProjectParamsWithTimeout creates a new DeleteProjectParams object
// with the ability to set a timeout on a request.
func NewDeleteProjectParamsWithTimeout(timeout time.Duration) *DeleteProjectParams {
	return &DeleteProjectParams{
		timeout: timeout,
	}
}

// NewDeleteProjectParamsWithContext creates a new DeleteProjectParams object
// with the ability to set a context for a request.
func NewDeleteProjectParamsWithContext(ctx context.Context) *DeleteProjectParams {
	return &DeleteProjectParams{
		Context: ctx,
	}
}

// NewDeleteProjectParamsWithHTTPClient creates a new DeleteProjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteProjectParamsWithHTTPClient(client *http.Client) *DeleteProjectParams {
	return &DeleteProjectParams{
		HTTPClient: client,
	}
}

/* DeleteProjectParams contains all the parameters to send to the API endpoint
   for the delete project operation.

   Typically these are written to a http.Request.
*/
type DeleteProjectParams struct {

	/* XIsResourceName.

	   The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name.
	*/
	XIsResourceName *bool `js:"xIsResourceName"`

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string `js:"xRequestID"`

	/* ProjectNameOrID.

	   The name or id of the project
	*/
	ProjectNameOrID string `js:"projectNameOrID"`

	timeout    time.Duration
	Context    context.Context `js:"context"`
	HTTPClient *http.Client    `js:"httpClient"`
}

// WithDefaults hydrates default values in the delete project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteProjectParams) WithDefaults() *DeleteProjectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteProjectParams) SetDefaults() {
	var (
		xIsResourceNameDefault = bool(false)
	)

	val := DeleteProjectParams{
		XIsResourceName: &xIsResourceNameDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete project params
func (o *DeleteProjectParams) WithTimeout(timeout time.Duration) *DeleteProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete project params
func (o *DeleteProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete project params
func (o *DeleteProjectParams) WithContext(ctx context.Context) *DeleteProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete project params
func (o *DeleteProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete project params
func (o *DeleteProjectParams) WithHTTPClient(client *http.Client) *DeleteProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete project params
func (o *DeleteProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXIsResourceName adds the xIsResourceName to the delete project params
func (o *DeleteProjectParams) WithXIsResourceName(xIsResourceName *bool) *DeleteProjectParams {
	o.SetXIsResourceName(xIsResourceName)
	return o
}

// SetXIsResourceName adds the xIsResourceName to the delete project params
func (o *DeleteProjectParams) SetXIsResourceName(xIsResourceName *bool) {
	o.XIsResourceName = xIsResourceName
}

// WithXRequestID adds the xRequestID to the delete project params
func (o *DeleteProjectParams) WithXRequestID(xRequestID *string) *DeleteProjectParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the delete project params
func (o *DeleteProjectParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithProjectNameOrID adds the projectNameOrID to the delete project params
func (o *DeleteProjectParams) WithProjectNameOrID(projectNameOrID string) *DeleteProjectParams {
	o.SetProjectNameOrID(projectNameOrID)
	return o
}

// SetProjectNameOrID adds the projectNameOrId to the delete project params
func (o *DeleteProjectParams) SetProjectNameOrID(projectNameOrID string) {
	o.ProjectNameOrID = projectNameOrID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XIsResourceName != nil {

		// header param X-Is-Resource-Name
		if err := r.SetHeaderParam("X-Is-Resource-Name", swag.FormatBool(*o.XIsResourceName)); err != nil {
			return err
		}
	}

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}

	// path param project_name_or_id
	if err := r.SetPathParam("project_name_or_id", o.ProjectNameOrID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
