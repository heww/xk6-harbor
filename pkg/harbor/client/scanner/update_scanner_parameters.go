// Code generated by go-swagger; DO NOT EDIT.

package scanner

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

	"github.com/heww/xk6-harbor/pkg/harbor/models"
)

// NewUpdateScannerParams creates a new UpdateScannerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateScannerParams() *UpdateScannerParams {
	return &UpdateScannerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateScannerParamsWithTimeout creates a new UpdateScannerParams object
// with the ability to set a timeout on a request.
func NewUpdateScannerParamsWithTimeout(timeout time.Duration) *UpdateScannerParams {
	return &UpdateScannerParams{
		timeout: timeout,
	}
}

// NewUpdateScannerParamsWithContext creates a new UpdateScannerParams object
// with the ability to set a context for a request.
func NewUpdateScannerParamsWithContext(ctx context.Context) *UpdateScannerParams {
	return &UpdateScannerParams{
		Context: ctx,
	}
}

// NewUpdateScannerParamsWithHTTPClient creates a new UpdateScannerParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateScannerParamsWithHTTPClient(client *http.Client) *UpdateScannerParams {
	return &UpdateScannerParams{
		HTTPClient: client,
	}
}

/* UpdateScannerParams contains all the parameters to send to the API endpoint
   for the update scanner operation.

   Typically these are written to a http.Request.
*/
type UpdateScannerParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string `js:"xRequestID"`

	/* Registration.

	   A scanner registraiton to be updated.
	*/
	Registration *models.ScannerRegistrationReq `js:"registration"`

	/* RegistrationID.

	   The scanner registration identifier.
	*/
	RegistrationID string `js:"registrationID"`

	timeout    time.Duration
	Context    context.Context `js:"context"`
	HTTPClient *http.Client    `js:"httpClient"`
}

// WithDefaults hydrates default values in the update scanner params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateScannerParams) WithDefaults() *UpdateScannerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update scanner params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateScannerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update scanner params
func (o *UpdateScannerParams) WithTimeout(timeout time.Duration) *UpdateScannerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update scanner params
func (o *UpdateScannerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update scanner params
func (o *UpdateScannerParams) WithContext(ctx context.Context) *UpdateScannerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update scanner params
func (o *UpdateScannerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update scanner params
func (o *UpdateScannerParams) WithHTTPClient(client *http.Client) *UpdateScannerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update scanner params
func (o *UpdateScannerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the update scanner params
func (o *UpdateScannerParams) WithXRequestID(xRequestID *string) *UpdateScannerParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the update scanner params
func (o *UpdateScannerParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithRegistration adds the registration to the update scanner params
func (o *UpdateScannerParams) WithRegistration(registration *models.ScannerRegistrationReq) *UpdateScannerParams {
	o.SetRegistration(registration)
	return o
}

// SetRegistration adds the registration to the update scanner params
func (o *UpdateScannerParams) SetRegistration(registration *models.ScannerRegistrationReq) {
	o.Registration = registration
}

// WithRegistrationID adds the registrationID to the update scanner params
func (o *UpdateScannerParams) WithRegistrationID(registrationID string) *UpdateScannerParams {
	o.SetRegistrationID(registrationID)
	return o
}

// SetRegistrationID adds the registrationId to the update scanner params
func (o *UpdateScannerParams) SetRegistrationID(registrationID string) {
	o.RegistrationID = registrationID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateScannerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Registration != nil {
		if err := r.SetBodyParam(o.Registration); err != nil {
			return err
		}
	}

	// path param registration_id
	if err := r.SetPathParam("registration_id", o.RegistrationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
