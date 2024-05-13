// Code generated by go-swagger; DO NOT EDIT.

package auditlog

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

// NewListAuditLogsParams creates a new ListAuditLogsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListAuditLogsParams() *ListAuditLogsParams {
	return &ListAuditLogsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListAuditLogsParamsWithTimeout creates a new ListAuditLogsParams object
// with the ability to set a timeout on a request.
func NewListAuditLogsParamsWithTimeout(timeout time.Duration) *ListAuditLogsParams {
	return &ListAuditLogsParams{
		timeout: timeout,
	}
}

// NewListAuditLogsParamsWithContext creates a new ListAuditLogsParams object
// with the ability to set a context for a request.
func NewListAuditLogsParamsWithContext(ctx context.Context) *ListAuditLogsParams {
	return &ListAuditLogsParams{
		Context: ctx,
	}
}

// NewListAuditLogsParamsWithHTTPClient creates a new ListAuditLogsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListAuditLogsParamsWithHTTPClient(client *http.Client) *ListAuditLogsParams {
	return &ListAuditLogsParams{
		HTTPClient: client,
	}
}

/*
ListAuditLogsParams contains all the parameters to send to the API endpoint

	for the list audit logs operation.

	Typically these are written to a http.Request.
*/
type ListAuditLogsParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string `js:"xRequestID"`

	/* Page.

	   The page number

	   Format: int64
	   Default: 1
	*/
	Page *int64 `js:"page"`

	/* PageSize.

	   The size of per page

	   Format: int64
	   Default: 10
	*/
	PageSize *int64 `js:"pageSize"`

	/* Q.

	   Query string to query resources. Supported query patterns are "exact match(k=v)", "fuzzy match(k=~v)", "range(k=[min~max])", "list with union releationship(k={v1 v2 v3})" and "list with intersetion relationship(k=(v1 v2 v3))". The value of range and list can be string(enclosed by " or '), integer or time(in format "2020-04-09 02:36:00"). All of these query patterns should be put in the query string "q=xxx" and splitted by ",". e.g. q=k1=v1,k2=~v2,k3=[min~max]
	*/
	Q *string `js:"q"`

	/* Sort.

	   Sort the resource list in ascending or descending order. e.g. sort by field1 in ascending order and field2 in descending order with "sort=field1,-field2"
	*/
	Sort *string `js:"sort"`

	timeout    time.Duration
	Context    context.Context `js:"context"`
	HTTPClient *http.Client    `js:"httpClient"`
}

// WithDefaults hydrates default values in the list audit logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAuditLogsParams) WithDefaults() *ListAuditLogsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list audit logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAuditLogsParams) SetDefaults() {
	var (
		pageDefault = int64(1)

		pageSizeDefault = int64(10)
	)

	val := ListAuditLogsParams{
		Page:     &pageDefault,
		PageSize: &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list audit logs params
func (o *ListAuditLogsParams) WithTimeout(timeout time.Duration) *ListAuditLogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list audit logs params
func (o *ListAuditLogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list audit logs params
func (o *ListAuditLogsParams) WithContext(ctx context.Context) *ListAuditLogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list audit logs params
func (o *ListAuditLogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list audit logs params
func (o *ListAuditLogsParams) WithHTTPClient(client *http.Client) *ListAuditLogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list audit logs params
func (o *ListAuditLogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the list audit logs params
func (o *ListAuditLogsParams) WithXRequestID(xRequestID *string) *ListAuditLogsParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the list audit logs params
func (o *ListAuditLogsParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithPage adds the page to the list audit logs params
func (o *ListAuditLogsParams) WithPage(page *int64) *ListAuditLogsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list audit logs params
func (o *ListAuditLogsParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the list audit logs params
func (o *ListAuditLogsParams) WithPageSize(pageSize *int64) *ListAuditLogsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list audit logs params
func (o *ListAuditLogsParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithQ adds the q to the list audit logs params
func (o *ListAuditLogsParams) WithQ(q *string) *ListAuditLogsParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the list audit logs params
func (o *ListAuditLogsParams) SetQ(q *string) {
	o.Q = q
}

// WithSort adds the sort to the list audit logs params
func (o *ListAuditLogsParams) WithSort(sort *string) *ListAuditLogsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the list audit logs params
func (o *ListAuditLogsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *ListAuditLogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize int64

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
		}
	}

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
