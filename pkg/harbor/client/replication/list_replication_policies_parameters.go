// Code generated by go-swagger; DO NOT EDIT.

package replication

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

// NewListReplicationPoliciesParams creates a new ListReplicationPoliciesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListReplicationPoliciesParams() *ListReplicationPoliciesParams {
	return &ListReplicationPoliciesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListReplicationPoliciesParamsWithTimeout creates a new ListReplicationPoliciesParams object
// with the ability to set a timeout on a request.
func NewListReplicationPoliciesParamsWithTimeout(timeout time.Duration) *ListReplicationPoliciesParams {
	return &ListReplicationPoliciesParams{
		timeout: timeout,
	}
}

// NewListReplicationPoliciesParamsWithContext creates a new ListReplicationPoliciesParams object
// with the ability to set a context for a request.
func NewListReplicationPoliciesParamsWithContext(ctx context.Context) *ListReplicationPoliciesParams {
	return &ListReplicationPoliciesParams{
		Context: ctx,
	}
}

// NewListReplicationPoliciesParamsWithHTTPClient creates a new ListReplicationPoliciesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListReplicationPoliciesParamsWithHTTPClient(client *http.Client) *ListReplicationPoliciesParams {
	return &ListReplicationPoliciesParams{
		HTTPClient: client,
	}
}

/* ListReplicationPoliciesParams contains all the parameters to send to the API endpoint
   for the list replication policies operation.

   Typically these are written to a http.Request.
*/
type ListReplicationPoliciesParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string `js:"xRequestID"`

	/* Name.

	   Deprecated, use "query" instead. The policy name.
	*/
	Name *string `js:"name"`

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

	   Sort the resource list in ascending or descending order. e.g. sort by field1 in ascending orderr and field2 in descending order with "sort=field1,-field2"
	*/
	Sort *string `js:"sort"`

	timeout    time.Duration
	Context    context.Context `js:"context"`
	HTTPClient *http.Client    `js:"httpClient"`
}

// WithDefaults hydrates default values in the list replication policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListReplicationPoliciesParams) WithDefaults() *ListReplicationPoliciesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list replication policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListReplicationPoliciesParams) SetDefaults() {
	var (
		pageDefault = int64(1)

		pageSizeDefault = int64(10)
	)

	val := ListReplicationPoliciesParams{
		Page:     &pageDefault,
		PageSize: &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list replication policies params
func (o *ListReplicationPoliciesParams) WithTimeout(timeout time.Duration) *ListReplicationPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list replication policies params
func (o *ListReplicationPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list replication policies params
func (o *ListReplicationPoliciesParams) WithContext(ctx context.Context) *ListReplicationPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list replication policies params
func (o *ListReplicationPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list replication policies params
func (o *ListReplicationPoliciesParams) WithHTTPClient(client *http.Client) *ListReplicationPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list replication policies params
func (o *ListReplicationPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the list replication policies params
func (o *ListReplicationPoliciesParams) WithXRequestID(xRequestID *string) *ListReplicationPoliciesParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the list replication policies params
func (o *ListReplicationPoliciesParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithName adds the name to the list replication policies params
func (o *ListReplicationPoliciesParams) WithName(name *string) *ListReplicationPoliciesParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the list replication policies params
func (o *ListReplicationPoliciesParams) SetName(name *string) {
	o.Name = name
}

// WithPage adds the page to the list replication policies params
func (o *ListReplicationPoliciesParams) WithPage(page *int64) *ListReplicationPoliciesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list replication policies params
func (o *ListReplicationPoliciesParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the list replication policies params
func (o *ListReplicationPoliciesParams) WithPageSize(pageSize *int64) *ListReplicationPoliciesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list replication policies params
func (o *ListReplicationPoliciesParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithQ adds the q to the list replication policies params
func (o *ListReplicationPoliciesParams) WithQ(q *string) *ListReplicationPoliciesParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the list replication policies params
func (o *ListReplicationPoliciesParams) SetQ(q *string) {
	o.Q = q
}

// WithSort adds the sort to the list replication policies params
func (o *ListReplicationPoliciesParams) WithSort(sort *string) *ListReplicationPoliciesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the list replication policies params
func (o *ListReplicationPoliciesParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *ListReplicationPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
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
