// Code generated by go-swagger; DO NOT EDIT.

package usergroup

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

// NewUpdateUserGroupParams creates a new UpdateUserGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateUserGroupParams() *UpdateUserGroupParams {
	return &UpdateUserGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUserGroupParamsWithTimeout creates a new UpdateUserGroupParams object
// with the ability to set a timeout on a request.
func NewUpdateUserGroupParamsWithTimeout(timeout time.Duration) *UpdateUserGroupParams {
	return &UpdateUserGroupParams{
		timeout: timeout,
	}
}

// NewUpdateUserGroupParamsWithContext creates a new UpdateUserGroupParams object
// with the ability to set a context for a request.
func NewUpdateUserGroupParamsWithContext(ctx context.Context) *UpdateUserGroupParams {
	return &UpdateUserGroupParams{
		Context: ctx,
	}
}

// NewUpdateUserGroupParamsWithHTTPClient creates a new UpdateUserGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateUserGroupParamsWithHTTPClient(client *http.Client) *UpdateUserGroupParams {
	return &UpdateUserGroupParams{
		HTTPClient: client,
	}
}

/* UpdateUserGroupParams contains all the parameters to send to the API endpoint
   for the update user group operation.

   Typically these are written to a http.Request.
*/
type UpdateUserGroupParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string `js:"xRequestID"`

	/* GroupID.

	   Group ID

	   Format: int64
	*/
	GroupID int64 `js:"groupID"`

	// Usergroup.
	Usergroup *models.UserGroup `js:"usergroup"`

	timeout    time.Duration
	Context    context.Context `js:"context"`
	HTTPClient *http.Client    `js:"httpClient"`
}

// WithDefaults hydrates default values in the update user group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUserGroupParams) WithDefaults() *UpdateUserGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update user group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUserGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update user group params
func (o *UpdateUserGroupParams) WithTimeout(timeout time.Duration) *UpdateUserGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update user group params
func (o *UpdateUserGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update user group params
func (o *UpdateUserGroupParams) WithContext(ctx context.Context) *UpdateUserGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update user group params
func (o *UpdateUserGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update user group params
func (o *UpdateUserGroupParams) WithHTTPClient(client *http.Client) *UpdateUserGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update user group params
func (o *UpdateUserGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the update user group params
func (o *UpdateUserGroupParams) WithXRequestID(xRequestID *string) *UpdateUserGroupParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the update user group params
func (o *UpdateUserGroupParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithGroupID adds the groupID to the update user group params
func (o *UpdateUserGroupParams) WithGroupID(groupID int64) *UpdateUserGroupParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the update user group params
func (o *UpdateUserGroupParams) SetGroupID(groupID int64) {
	o.GroupID = groupID
}

// WithUsergroup adds the usergroup to the update user group params
func (o *UpdateUserGroupParams) WithUsergroup(usergroup *models.UserGroup) *UpdateUserGroupParams {
	o.SetUsergroup(usergroup)
	return o
}

// SetUsergroup adds the usergroup to the update user group params
func (o *UpdateUserGroupParams) SetUsergroup(usergroup *models.UserGroup) {
	o.Usergroup = usergroup
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUserGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param group_id
	if err := r.SetPathParam("group_id", swag.FormatInt64(o.GroupID)); err != nil {
		return err
	}
	if o.Usergroup != nil {
		if err := r.SetBodyParam(o.Usergroup); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
