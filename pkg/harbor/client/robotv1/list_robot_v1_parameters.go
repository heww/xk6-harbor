// Code generated by go-swagger; DO NOT EDIT.

package robotv1

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

// NewListRobotV1Params creates a new ListRobotV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListRobotV1Params() *ListRobotV1Params {
	return &ListRobotV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewListRobotV1ParamsWithTimeout creates a new ListRobotV1Params object
// with the ability to set a timeout on a request.
func NewListRobotV1ParamsWithTimeout(timeout time.Duration) *ListRobotV1Params {
	return &ListRobotV1Params{
		timeout: timeout,
	}
}

// NewListRobotV1ParamsWithContext creates a new ListRobotV1Params object
// with the ability to set a context for a request.
func NewListRobotV1ParamsWithContext(ctx context.Context) *ListRobotV1Params {
	return &ListRobotV1Params{
		Context: ctx,
	}
}

// NewListRobotV1ParamsWithHTTPClient creates a new ListRobotV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewListRobotV1ParamsWithHTTPClient(client *http.Client) *ListRobotV1Params {
	return &ListRobotV1Params{
		HTTPClient: client,
	}
}

/*
ListRobotV1Params contains all the parameters to send to the API endpoint

	for the list robot v1 operation.

	Typically these are written to a http.Request.
*/
type ListRobotV1Params struct {

	/* XIsResourceName.

	   The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name.
	*/
	XIsResourceName *bool `js:"xIsResourceName"`

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

	/* ProjectNameOrID.

	   The name or id of the project
	*/
	ProjectNameOrID string `js:"projectNameOrID"`

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

// WithDefaults hydrates default values in the list robot v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListRobotV1Params) WithDefaults() *ListRobotV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list robot v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListRobotV1Params) SetDefaults() {
	var (
		xIsResourceNameDefault = bool(false)

		pageDefault = int64(1)

		pageSizeDefault = int64(10)
	)

	val := ListRobotV1Params{
		XIsResourceName: &xIsResourceNameDefault,
		Page:            &pageDefault,
		PageSize:        &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list robot v1 params
func (o *ListRobotV1Params) WithTimeout(timeout time.Duration) *ListRobotV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list robot v1 params
func (o *ListRobotV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list robot v1 params
func (o *ListRobotV1Params) WithContext(ctx context.Context) *ListRobotV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list robot v1 params
func (o *ListRobotV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list robot v1 params
func (o *ListRobotV1Params) WithHTTPClient(client *http.Client) *ListRobotV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list robot v1 params
func (o *ListRobotV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXIsResourceName adds the xIsResourceName to the list robot v1 params
func (o *ListRobotV1Params) WithXIsResourceName(xIsResourceName *bool) *ListRobotV1Params {
	o.SetXIsResourceName(xIsResourceName)
	return o
}

// SetXIsResourceName adds the xIsResourceName to the list robot v1 params
func (o *ListRobotV1Params) SetXIsResourceName(xIsResourceName *bool) {
	o.XIsResourceName = xIsResourceName
}

// WithXRequestID adds the xRequestID to the list robot v1 params
func (o *ListRobotV1Params) WithXRequestID(xRequestID *string) *ListRobotV1Params {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the list robot v1 params
func (o *ListRobotV1Params) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithPage adds the page to the list robot v1 params
func (o *ListRobotV1Params) WithPage(page *int64) *ListRobotV1Params {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list robot v1 params
func (o *ListRobotV1Params) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the list robot v1 params
func (o *ListRobotV1Params) WithPageSize(pageSize *int64) *ListRobotV1Params {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list robot v1 params
func (o *ListRobotV1Params) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithProjectNameOrID adds the projectNameOrID to the list robot v1 params
func (o *ListRobotV1Params) WithProjectNameOrID(projectNameOrID string) *ListRobotV1Params {
	o.SetProjectNameOrID(projectNameOrID)
	return o
}

// SetProjectNameOrID adds the projectNameOrId to the list robot v1 params
func (o *ListRobotV1Params) SetProjectNameOrID(projectNameOrID string) {
	o.ProjectNameOrID = projectNameOrID
}

// WithQ adds the q to the list robot v1 params
func (o *ListRobotV1Params) WithQ(q *string) *ListRobotV1Params {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the list robot v1 params
func (o *ListRobotV1Params) SetQ(q *string) {
	o.Q = q
}

// WithSort adds the sort to the list robot v1 params
func (o *ListRobotV1Params) WithSort(sort *string) *ListRobotV1Params {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the list robot v1 params
func (o *ListRobotV1Params) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *ListRobotV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param project_name_or_id
	if err := r.SetPathParam("project_name_or_id", o.ProjectNameOrID); err != nil {
		return err
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
