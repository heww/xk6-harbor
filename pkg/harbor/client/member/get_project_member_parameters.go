// Code generated by go-swagger; DO NOT EDIT.

package member

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

// NewGetProjectMemberParams creates a new GetProjectMemberParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProjectMemberParams() *GetProjectMemberParams {
	return &GetProjectMemberParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectMemberParamsWithTimeout creates a new GetProjectMemberParams object
// with the ability to set a timeout on a request.
func NewGetProjectMemberParamsWithTimeout(timeout time.Duration) *GetProjectMemberParams {
	return &GetProjectMemberParams{
		timeout: timeout,
	}
}

// NewGetProjectMemberParamsWithContext creates a new GetProjectMemberParams object
// with the ability to set a context for a request.
func NewGetProjectMemberParamsWithContext(ctx context.Context) *GetProjectMemberParams {
	return &GetProjectMemberParams{
		Context: ctx,
	}
}

// NewGetProjectMemberParamsWithHTTPClient creates a new GetProjectMemberParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProjectMemberParamsWithHTTPClient(client *http.Client) *GetProjectMemberParams {
	return &GetProjectMemberParams{
		HTTPClient: client,
	}
}

/* GetProjectMemberParams contains all the parameters to send to the API endpoint
   for the get project member operation.

   Typically these are written to a http.Request.
*/
type GetProjectMemberParams struct {

	/* XIsResourceName.

	   The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name.
	*/
	XIsResourceName *bool `js:"xIsResourceName"`

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string `js:"xRequestID"`

	/* Mid.

	   The member ID

	   Format: int64
	*/
	Mid int64 `js:"mid"`

	/* ProjectNameOrID.

	   The name or id of the project
	*/
	ProjectNameOrID string `js:"projectNameOrID"`

	timeout    time.Duration
	Context    context.Context `js:"context"`
	HTTPClient *http.Client    `js:"httpClient"`
}

// WithDefaults hydrates default values in the get project member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectMemberParams) WithDefaults() *GetProjectMemberParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get project member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectMemberParams) SetDefaults() {
	var (
		xIsResourceNameDefault = bool(false)
	)

	val := GetProjectMemberParams{
		XIsResourceName: &xIsResourceNameDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get project member params
func (o *GetProjectMemberParams) WithTimeout(timeout time.Duration) *GetProjectMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get project member params
func (o *GetProjectMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get project member params
func (o *GetProjectMemberParams) WithContext(ctx context.Context) *GetProjectMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get project member params
func (o *GetProjectMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get project member params
func (o *GetProjectMemberParams) WithHTTPClient(client *http.Client) *GetProjectMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get project member params
func (o *GetProjectMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXIsResourceName adds the xIsResourceName to the get project member params
func (o *GetProjectMemberParams) WithXIsResourceName(xIsResourceName *bool) *GetProjectMemberParams {
	o.SetXIsResourceName(xIsResourceName)
	return o
}

// SetXIsResourceName adds the xIsResourceName to the get project member params
func (o *GetProjectMemberParams) SetXIsResourceName(xIsResourceName *bool) {
	o.XIsResourceName = xIsResourceName
}

// WithXRequestID adds the xRequestID to the get project member params
func (o *GetProjectMemberParams) WithXRequestID(xRequestID *string) *GetProjectMemberParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get project member params
func (o *GetProjectMemberParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithMid adds the mid to the get project member params
func (o *GetProjectMemberParams) WithMid(mid int64) *GetProjectMemberParams {
	o.SetMid(mid)
	return o
}

// SetMid adds the mid to the get project member params
func (o *GetProjectMemberParams) SetMid(mid int64) {
	o.Mid = mid
}

// WithProjectNameOrID adds the projectNameOrID to the get project member params
func (o *GetProjectMemberParams) WithProjectNameOrID(projectNameOrID string) *GetProjectMemberParams {
	o.SetProjectNameOrID(projectNameOrID)
	return o
}

// SetProjectNameOrID adds the projectNameOrId to the get project member params
func (o *GetProjectMemberParams) SetProjectNameOrID(projectNameOrID string) {
	o.ProjectNameOrID = projectNameOrID
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param mid
	if err := r.SetPathParam("mid", swag.FormatInt64(o.Mid)); err != nil {
		return err
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
