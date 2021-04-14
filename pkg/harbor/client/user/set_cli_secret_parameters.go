// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewSetCliSecretParams creates a new SetCliSecretParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetCliSecretParams() *SetCliSecretParams {
	return &SetCliSecretParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetCliSecretParamsWithTimeout creates a new SetCliSecretParams object
// with the ability to set a timeout on a request.
func NewSetCliSecretParamsWithTimeout(timeout time.Duration) *SetCliSecretParams {
	return &SetCliSecretParams{
		timeout: timeout,
	}
}

// NewSetCliSecretParamsWithContext creates a new SetCliSecretParams object
// with the ability to set a context for a request.
func NewSetCliSecretParamsWithContext(ctx context.Context) *SetCliSecretParams {
	return &SetCliSecretParams{
		Context: ctx,
	}
}

// NewSetCliSecretParamsWithHTTPClient creates a new SetCliSecretParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetCliSecretParamsWithHTTPClient(client *http.Client) *SetCliSecretParams {
	return &SetCliSecretParams{
		HTTPClient: client,
	}
}

/* SetCliSecretParams contains all the parameters to send to the API endpoint
   for the set cli secret operation.

   Typically these are written to a http.Request.
*/
type SetCliSecretParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string `js:"xRequestID"`

	// Secret.
	Secret *models.OIDCCliSecretReq `js:"secret"`

	/* UserID.

	   User ID

	   Format: int
	*/
	UserID int64 `js:"userID"`

	timeout    time.Duration
	Context    context.Context `js:"context"`
	HTTPClient *http.Client    `js:"httpClient"`
}

// WithDefaults hydrates default values in the set cli secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetCliSecretParams) WithDefaults() *SetCliSecretParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set cli secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetCliSecretParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set cli secret params
func (o *SetCliSecretParams) WithTimeout(timeout time.Duration) *SetCliSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set cli secret params
func (o *SetCliSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set cli secret params
func (o *SetCliSecretParams) WithContext(ctx context.Context) *SetCliSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set cli secret params
func (o *SetCliSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set cli secret params
func (o *SetCliSecretParams) WithHTTPClient(client *http.Client) *SetCliSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set cli secret params
func (o *SetCliSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the set cli secret params
func (o *SetCliSecretParams) WithXRequestID(xRequestID *string) *SetCliSecretParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the set cli secret params
func (o *SetCliSecretParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithSecret adds the secret to the set cli secret params
func (o *SetCliSecretParams) WithSecret(secret *models.OIDCCliSecretReq) *SetCliSecretParams {
	o.SetSecret(secret)
	return o
}

// SetSecret adds the secret to the set cli secret params
func (o *SetCliSecretParams) SetSecret(secret *models.OIDCCliSecretReq) {
	o.Secret = secret
}

// WithUserID adds the userID to the set cli secret params
func (o *SetCliSecretParams) WithUserID(userID int64) *SetCliSecretParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the set cli secret params
func (o *SetCliSecretParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *SetCliSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Secret != nil {
		if err := r.SetBodyParam(o.Secret); err != nil {
			return err
		}
	}

	// path param user_id
	if err := r.SetPathParam("user_id", swag.FormatInt64(o.UserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
