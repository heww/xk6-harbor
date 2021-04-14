// Code generated by go-swagger; DO NOT EDIT.

package repository

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

// NewDeleteRepositoryParams creates a new DeleteRepositoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteRepositoryParams() *DeleteRepositoryParams {
	return &DeleteRepositoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRepositoryParamsWithTimeout creates a new DeleteRepositoryParams object
// with the ability to set a timeout on a request.
func NewDeleteRepositoryParamsWithTimeout(timeout time.Duration) *DeleteRepositoryParams {
	return &DeleteRepositoryParams{
		timeout: timeout,
	}
}

// NewDeleteRepositoryParamsWithContext creates a new DeleteRepositoryParams object
// with the ability to set a context for a request.
func NewDeleteRepositoryParamsWithContext(ctx context.Context) *DeleteRepositoryParams {
	return &DeleteRepositoryParams{
		Context: ctx,
	}
}

// NewDeleteRepositoryParamsWithHTTPClient creates a new DeleteRepositoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteRepositoryParamsWithHTTPClient(client *http.Client) *DeleteRepositoryParams {
	return &DeleteRepositoryParams{
		HTTPClient: client,
	}
}

/* DeleteRepositoryParams contains all the parameters to send to the API endpoint
   for the delete repository operation.

   Typically these are written to a http.Request.
*/
type DeleteRepositoryParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string `js:"xRequestID"`

	/* ProjectName.

	   The name of the project
	*/
	ProjectName string `js:"projectName"`

	/* RepositoryName.

	   The name of the repository. If it contains slash, encode it with URL encoding. e.g. a/b -> a%252Fb
	*/
	RepositoryName string `js:"repositoryName"`

	timeout    time.Duration
	Context    context.Context `js:"context"`
	HTTPClient *http.Client    `js:"httpClient"`
}

// WithDefaults hydrates default values in the delete repository params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRepositoryParams) WithDefaults() *DeleteRepositoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete repository params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRepositoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete repository params
func (o *DeleteRepositoryParams) WithTimeout(timeout time.Duration) *DeleteRepositoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete repository params
func (o *DeleteRepositoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete repository params
func (o *DeleteRepositoryParams) WithContext(ctx context.Context) *DeleteRepositoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete repository params
func (o *DeleteRepositoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete repository params
func (o *DeleteRepositoryParams) WithHTTPClient(client *http.Client) *DeleteRepositoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete repository params
func (o *DeleteRepositoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the delete repository params
func (o *DeleteRepositoryParams) WithXRequestID(xRequestID *string) *DeleteRepositoryParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the delete repository params
func (o *DeleteRepositoryParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithProjectName adds the projectName to the delete repository params
func (o *DeleteRepositoryParams) WithProjectName(projectName string) *DeleteRepositoryParams {
	o.SetProjectName(projectName)
	return o
}

// SetProjectName adds the projectName to the delete repository params
func (o *DeleteRepositoryParams) SetProjectName(projectName string) {
	o.ProjectName = projectName
}

// WithRepositoryName adds the repositoryName to the delete repository params
func (o *DeleteRepositoryParams) WithRepositoryName(repositoryName string) *DeleteRepositoryParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the delete repository params
func (o *DeleteRepositoryParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRepositoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param project_name
	if err := r.SetPathParam("project_name", o.ProjectName); err != nil {
		return err
	}

	// path param repository_name
	if err := r.SetPathParam("repository_name", o.RepositoryName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
