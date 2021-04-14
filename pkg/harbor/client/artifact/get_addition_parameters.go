// Code generated by go-swagger; DO NOT EDIT.

package artifact

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

// NewGetAdditionParams creates a new GetAdditionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAdditionParams() *GetAdditionParams {
	return &GetAdditionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAdditionParamsWithTimeout creates a new GetAdditionParams object
// with the ability to set a timeout on a request.
func NewGetAdditionParamsWithTimeout(timeout time.Duration) *GetAdditionParams {
	return &GetAdditionParams{
		timeout: timeout,
	}
}

// NewGetAdditionParamsWithContext creates a new GetAdditionParams object
// with the ability to set a context for a request.
func NewGetAdditionParamsWithContext(ctx context.Context) *GetAdditionParams {
	return &GetAdditionParams{
		Context: ctx,
	}
}

// NewGetAdditionParamsWithHTTPClient creates a new GetAdditionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAdditionParamsWithHTTPClient(client *http.Client) *GetAdditionParams {
	return &GetAdditionParams{
		HTTPClient: client,
	}
}

/* GetAdditionParams contains all the parameters to send to the API endpoint
   for the get addition operation.

   Typically these are written to a http.Request.
*/
type GetAdditionParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string `js:"xRequestID"`

	/* Addition.

	   The type of addition.
	*/
	Addition string `js:"addition"`

	/* ProjectName.

	   The name of the project
	*/
	ProjectName string `js:"projectName"`

	/* Reference.

	   The reference of the artifact, can be digest or tag
	*/
	Reference string `js:"reference"`

	/* RepositoryName.

	   The name of the repository. If it contains slash, encode it with URL encoding. e.g. a/b -> a%252Fb
	*/
	RepositoryName string `js:"repositoryName"`

	timeout    time.Duration
	Context    context.Context `js:"context"`
	HTTPClient *http.Client    `js:"httpClient"`
}

// WithDefaults hydrates default values in the get addition params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAdditionParams) WithDefaults() *GetAdditionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get addition params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAdditionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get addition params
func (o *GetAdditionParams) WithTimeout(timeout time.Duration) *GetAdditionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get addition params
func (o *GetAdditionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get addition params
func (o *GetAdditionParams) WithContext(ctx context.Context) *GetAdditionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get addition params
func (o *GetAdditionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get addition params
func (o *GetAdditionParams) WithHTTPClient(client *http.Client) *GetAdditionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get addition params
func (o *GetAdditionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get addition params
func (o *GetAdditionParams) WithXRequestID(xRequestID *string) *GetAdditionParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get addition params
func (o *GetAdditionParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithAddition adds the addition to the get addition params
func (o *GetAdditionParams) WithAddition(addition string) *GetAdditionParams {
	o.SetAddition(addition)
	return o
}

// SetAddition adds the addition to the get addition params
func (o *GetAdditionParams) SetAddition(addition string) {
	o.Addition = addition
}

// WithProjectName adds the projectName to the get addition params
func (o *GetAdditionParams) WithProjectName(projectName string) *GetAdditionParams {
	o.SetProjectName(projectName)
	return o
}

// SetProjectName adds the projectName to the get addition params
func (o *GetAdditionParams) SetProjectName(projectName string) {
	o.ProjectName = projectName
}

// WithReference adds the reference to the get addition params
func (o *GetAdditionParams) WithReference(reference string) *GetAdditionParams {
	o.SetReference(reference)
	return o
}

// SetReference adds the reference to the get addition params
func (o *GetAdditionParams) SetReference(reference string) {
	o.Reference = reference
}

// WithRepositoryName adds the repositoryName to the get addition params
func (o *GetAdditionParams) WithRepositoryName(repositoryName string) *GetAdditionParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the get addition params
func (o *GetAdditionParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *GetAdditionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param addition
	if err := r.SetPathParam("addition", o.Addition); err != nil {
		return err
	}

	// path param project_name
	if err := r.SetPathParam("project_name", o.ProjectName); err != nil {
		return err
	}

	// path param reference
	if err := r.SetPathParam("reference", o.Reference); err != nil {
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
